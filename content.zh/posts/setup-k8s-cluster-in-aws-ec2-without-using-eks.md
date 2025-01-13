---
title: 使用AWS EC2搭建Kubernetes集群
date: 2024-11-20T02:01:58+05:30
tags: [ computer-science, aws, kubernetes ]
categories: study
canonicalUrl: https://wenstudy.com/posts/setup-k8s-cluster-in-aws-ec2-without-using-eks/
---

## 准备EC2实例
通过AWS控制台或者AWS CLI创建两个EC2实例，一个作为Master节点，另一个作为Worker节点。

使用AWS AMI创建EC2实例，选择Amazon Linux 2 AMI。

创建key pair，用于远程SSH登录EC2实例。创建后会下载一个私钥文件，一般存入`~/.ssh`目录下，并修改权限为只允许拥有者读取。
```bash
chmod 400 your-private-key.pem
```

## 预备

创建好后，需要在裸机上安装各种依赖。首先更新 dnf 包管理器。因为用的是 Amazon Linux 2 AMI，所以用的是 dnf 包管理器，而不是 yum。
```bash
sudo dnf update -y
```

### 准备容器运行时 containerd
```bash
sudo dnf install -y containerd
```

启动 containerd 服务
```bash
sudo systemctl daemon-reload
sudo systemctl enable --now containerd
```

为 containerd 创建默认配置文件
```bash
sudo containerd config default | sudo tee /etc/containerd/config.toml
```

编辑 `/etc/containerd/config.toml` 文件，将 `systemd_cgroup = false` 修改为 `systemd_cgroup = true`，以使 containerd 使用 systemd cgroup 驱动。
```bash
sudo sed -i 's/SystemdCgroup = false/SystemdCgroup = true/' /etc/containerd/config.toml
```

重启 containerd 服务并查看状态
```bash
sudo systemctl restart containerd  
sudo systemctl status containerd
```

### 准备网络插件 CNI
CNI (Container Network Interface) 插件为 Kubernetes 提供网络功能。它主要负责：
1. 为 Pod 分配 IP 地址。
2. 设置容器之间的网络通信。
3. 确保 Pod 可以与其他 Pod、服务（Service）以及外部世界通信。
4. 
如果没有正确安装和配置 CNI 插件，Kubernetes 的 Pod 网络将无法正常工作，导致 Pod 无法互通或无法分配 IP 地址。先查看是否已经安装了 CNI 插件
```bash
ls /opt/cni/bin
```

如果以上命令没有输出如下信息，说明没有安装 CNI 插件
```
ls: cannot access '/opt/cni/bin': No such file or directory
```

通过以下命令安装
```bash
CNI_VERSION=1.4.0
sudo wget -P /usr/local/src https://github.com/containernetworking/plugins/releases/download/v${CNI_VERSION}/cni-plugins-linux-amd64-v${CNI_VERSION}.tgz
sudo mkdir -p /opt/cni/bin
sudo tar -C /opt/cni/bin -xf /usr/local/src/cni-plugins-linux-amd64-v${CNI_VERSION}.tgz
```

再使用 `ls /opt/cni/bin` 命令查看，应当看到如下输出
```
bridge
host-local
loopback
portmap
firewall
...
```

### 其他必要配置
禁用 swap
```bash
sudo swapoff -a
```

编辑 `/etc/fstab` 文件，注释掉 swap 行，以永久禁用 swap
```bash
sudo sed -i '/ swap / s/^/#/' /etc/fstab
```

启动 overlay 和 br_netfilter 内核模块，它们是 Kubernetes 集群所必需的两个内核模块，需要手动加载。

```bash
sudo modprobe overlay  
sudo modprobe br_netfilter
```

通过编辑 `/etc/modules-load.d/k8s.conf` 文件，使得这两个模块在系统启动时自动加载
```bash
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF
```

检查是否加载成功
```bash
lsmod | grep -e overlay -e br_netfilter
```

修改 sysctl 配置
这里修改 sysctl 配置，使得 iptables 能够正确工作。
```bash
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF
```

### 安装 Kubernetes

首先安装 curl（已有则跳过）
```bash
sudo dnf install -y curl
```

添加 Kubernetes 仓库，这条命令保证了 yum 能够从 Google 的仓库中下载 Kubernetes 的软件包
```bash
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo tee /etc/yum.repos.d/kubernetes.repo > /dev/null
```

以下命令添加了 Kubernetes 仓库的地址到 `/etc/yum.repos.d/kubernetes.repo` 文件中
```bash
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://pkgs.k8s.io/core:/stable:/v1.32/rpm/
enabled=1
gpgcheck=1
gpgkey=https://pkgs.k8s.io/core:/stable:/v1.32/rpm/repodata/repomd.xml.key
exclude=kubelet kubeadm kubectl cri-tools kubernetes-cni
EOF
```

更新包管理器的缓存让其识别新的仓库
```bash
sudo dnf makecache
```

安装 kubeadm, kubelet 和 kubectl

```bash
sudo dnf install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
```

安装 `tc (Traffic Control) 包`，以便 Kubernetes 集群能够正常工作，首先查看包含 tc 命令的软件包
```bash
dnf provides tc
```

一般都存在于 `iproute-tc` 软件包中，安装 iproute-tc 软件包
```bash
sudo dnf install -y iproute-tc
```

启动 kubelet 服务
```bash
sudo systemctl enable kubelet
```

```bash
sudo systemctl start kubelet
```

检查 kubelet 服务状态
```bash
sudo systemctl status kubelet
```


初始化 Kubernetes 集群
```bash
sudo kubeadm init --pod-network-cidr=192.168.0.0/16
```

初始化完成后，会输出类似如下的信息，其中有两个命令，一个是 `kubeadm join` 命令，另一个是 `kubectl apply` 命令，分别用于加入节点和安装网络插件。

```bash
Your Kubernetes control-plane has initialized successfully!
```

将 `kubectl` 配置文件拷贝到当前用户的家目录下
```bash
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

安装CNI网络插件，这里使用 Calico 网络插件
```bash
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.29.1/manifests/tigera-operator.yaml

```

```bash
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.29.1/manifests/custom-resources.yaml
```

查看节点状态
```bash
kubectl get pods -n calico-system
```

### 删除 kubernetes 及其依赖

```bash
sudo systemctl stop kubelet
sudo dnf remove -y kubelet kubeadm kubectl
sudo rm -rf /etc/kubernetes /var/lib/kubelet /var/lib/etcd ~/.kube
sudo kubeadm reset -f
sudo rm -rf /etc/cni/net.d
sudo systemctl restart containerd
```

> 这是为了在重新安装或升级 kubernetes 时，不会因为之前的配置文件残留导致安装失败。
