---
title: 使用AWS EC2搭建Kubernetes集群
date: 2024-11-20T02:01:58+05:30
tags: [ computer-science, aws, kubernetes ]
categories: study
canonicalUrl: https://wenstudy.com/posts/setup-k8s-cluster-in-aws-ec2-without-using-eks/
---

## 准备EC2实例

通过AWS控制台或者AWS CLI创建两个EC2实例，一个作为Master节点，另一个作为Worker节点。实例需求：

- Master节点：2 vCPU，4GB内存，20GB硬盘
- Worker节点：1 vCPU，2GB内存，20GB硬盘

使用AWS AMI创建EC2实例，选择Amazon Linux 2 AMI。

### 安全组配置

在创建EC2实例时，需要配置安全组，如果不完全开放所有端口，至少需要开放以下端口：

创建key pair，用于远程SSH登录EC2实例。创建后会下载一个私钥文件，一般存入`~/.ssh`目录下，并修改权限为只允许拥有者读取。

```bash
chmod 400 your-private-key.pem
```

创建好后，需要在裸机上安装各种依赖。首先更新 dnf 包管理器。因为用的是 Amazon Linux 2 AMI，所以用的是 dnf 包管理器，而不是
yum。

```bash
sudo dnf update -y
```

## 准备 containerd

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

编辑 `/etc/containerd/config.toml` 文件，将 `systemd_cgroup = false` 修改为 `systemd_cgroup = true`，以使 containerd 使用
systemd cgroup 驱动。

```bash
sudo sed -i 's/SystemdCgroup = false/SystemdCgroup = true/' /etc/containerd/config.toml
```

重启 containerd 服务并查看状态

```bash
sudo systemctl restart containerd  
sudo systemctl status containerd
```

## 准备 CNI

CNI (Container Network Interface) 插件为 Kubernetes 提供网络功能。它主要负责：

1. 为 Pod 分配 IP 地址。
2. 设置容器之间的网络通信。
3. 确保 Pod 可以与其他 Pod、服务（Service）以及外部世界通信。
4.

如果没有正确安装和配置 CNI 插件，Kubernetes 的 Pod 网络将无法正常工作，导致 Pod 无法互通或无法分配 IP 地址。先查看是否已经安装了
CNI 插件

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

## 其他必要配置

通过编辑 `/etc/modules-load.d/k8s.conf` 文件，使得这两个模块在系统启动时自动加载

```bash
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF
```

启动 overlay 和 br_netfilter 内核模块，它们是 Kubernetes 集群所必需的两个内核模块，需要手动加载。

```bash
sudo modprobe overlay  
sudo modprobe br_netfilter
```

检查是否加载成功

```bash
lsmod | grep -e overlay -e br_netfilter
```

修改 sysctl 配置，使得 iptables 能够正确工作。

```bash
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF
```

使配置生效

```bash
sudo sysctl --system
```

禁用 swap

```bash
sudo swapon -s # 查看 swap 分区
sudo swapoff -a
```

## 安装 Kubernetes

首先安装 curl（已有则跳过）

```bash
sudo dnf install -y curl
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

安装 kubeadm, kubelet 和 kubectl 并启动 kubelet 服务

```bash
sudo dnf install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
sudo systemctl enable --now kubelet
```

安装 `tc (Traffic Control) 包`，以便 Kubernetes 集群能够正常工作，首先查看包含 tc 命令的软件包

```bash
dnf provides tc
```

一般都存在于 `iproute-tc` 软件包中，安装 iproute-tc 软件包

```bash
sudo dnf install -y iproute-tc
```

初始化 Kubernetes 集群

```bash
sudo kubeadm init --pod-network-cidr=10.244.0.0/16
```

> 1. 一定要用 `sudo` 执行 `kubeadm init` 命令，因为这个命令会修改系统的配置文件，否则会报错。
> 2. 这里 CIDR 使用 10.244.0.0/16，因为稍后配置的网络插件 Flannel 默认使用这个 CIDR。

检查 kubelet 服务状态

```bash
systemctl status kubelet
```

初始化完成后，会输出类似如下的信息，其中有两个命令，一个是 `kubeadm join` 命令，另一个是 `kubectl apply` 命令，分别用于加入节点和安装网络插件。

```
Your Kubernetes control-plane has initialized successfully!
```

将 `kubectl` 配置文件拷贝到当前用户的家目录下

```bash
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

安装CNI网络插件，这里使用 Flannel。
Flannel 是一种简单的 Kubernetes 网络解决方案。它会为每个 Pod 分配一个唯一的 IP 地址，并确保不同节点之间的 Pod 能通过虚拟网络通信。

部署 Flannel

```bash
kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml
```

检查 Flannel Pod 是否正常运行

```bash
kubectl get pods -n kube-flannel
```

应得到如下输出，flannel的Pod应该是 `Running` 状态

```
NAME                    READY   STATUS    RESTARTS       AGE
kube-flannel-ds-n2nzv   1/1     Running   20 (52s ago)   67m
```

以及用以下命令查看 Flannel DaemonSet

```bash
kubectl get daemonset kube-flannel-ds -n kube-flannel
```

应得到如下输出，`AVAILABLE` 列应该是 `1`

```
NAME             DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
kube-flannel-ds  1         1         1       1            1           <none>                   67m
```

查看系统相关的 Pod

```bash
kubectl get pods -n kube-system
```

应得到如下输出，所有的 Pod 都应该是 `Running` 状态

```
NAME                                                                      READY   STATUS    RESTARTS   AGE
coredns-668d6bf9bc-xls74                                                  1/1     Running   0          70m
coredns-668d6bf9bc-zghmb                                                  1/1     Running   0          70m
etcd-ip-123-12-12-12.ap-northeast-1.compute.internal                      1/1     Running   0          70m
kube-apiserver-ip-123-12-12-12.ap-northeast-1.compute.internal            1/1     Running   0          70m
kube-controller-manager-ip-123-12-12-12.ap-northeast-1.compute.internal   1/1     Running   0          70m
kube-proxy-f2nnh                                                          1/1     Running   0          70m
kube-scheduler-ip-123-12-12-12.ap-northeast-1.compute.internal            1/1     Running   0          70m
```

查看节点状态

```bash
kubectl get nodes
```

应得到如下输出

```
NAME                                              STATUS   ROLES           AGE   VERSION
ip-123-12-12-12.ap-northeast-1.compute.internal   Ready    control-plane   72m   v1.32.0
```

## 集成 ArgoCD

ArgoCD 是一个用于 GitOps 部署的工具，它可以帮助我们将应用程序的配置文件存储在 Git 仓库中，并通过 ArgoCD 自动同步到
Kubernetes 集群中。

创建 ArgoCD 命名空间（Namespace）

```bash
kubectl create namespace argocd
```

安装 ArgoCD

```bash
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

如果是单节点集群，那么 control-plane 节点也需要运行普通 Pod，即 argocd 相关 Pod。然而 control-plane 节点一般带有
`node-role.kubernetes.io/control-plane` taint，这会导致普通 Pod 无法调度到 control-plane 节点上。
这里提供用于测试的快速解决方案，即移除 control-plane 节点上的 taint，使其可以运行普通 Pod。

```bash
kubectl taint nodes --all node-role.kubernetes.io/control-plane-
```

查看 ArgoCD 相关的 Pod 状态

```bash
kubectl get pods -n argocd
```

所有的 Pod 都应该是 `Running` 状态

```
NAME                                                READY   STATUS    RESTARTS   AGE
argocd-application-controller-0                     1/1     Running   0          6m14s
argocd-applicationset-controller-84c66c76b4-zf6bd   1/1     Running   0          6m15s
argocd-dex-server-97fdcf666-4pbwc                   1/1     Running   0          6m15s
argocd-notifications-controller-7c44dd7757-2twjj    1/1     Running   0          6m15s
argocd-redis-596f8956cc-wf6md                       1/1     Running   0          6m15s
argocd-repo-server-577664cf68-ngbcz                 1/1     Running   0          6m15s
argocd-server-6b9b64c5fb-qtk6b                      1/1     Running   0          6m15s
```

暴露 ArgoCD Server
    
```bash
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'
```

查看 ArgoCD Server 的 NodePort 端口

```bash
kubectl get svc argocd-server -n argocd
```

应得到如下输出

```
NAME            TYPE       CLUSTER-IP    EXTERNAL-IP   PORT(S)                      AGE
argocd-server   NodePort   10.97.82.16   <none>        80:31027/TCP,443:32130/TCP   14m
```

> 其中 `80:31027/TCP` 是 HTTP 端口，`443:32130/TCP` 是 HTTPS 端口。可以通过 `http://<control-plane-ip>:31027` 访问 ArgoCD UI。

获取 control-plane 节点的公网 IP 地址，然后访问 `http://<control-plane-ip>:31027`。对于 AWS EC2 实例，可以在 AWS 控制台中查看。

获取 ArgoCD Server 的初始密码

```bash
kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d
```

> 用户名为 `admin`，密码为上面的命令输出。建议在登陆后立即修改管理员密码。

## 加入节点

在 Worker 节点上执行 `kubeadm join` 命令，将 Worker 节点加入到 Kubernetes 集群中。

```bash
sudo kubeadm join
```

## 问题排查

以下命令创建一个临时的 Pod，用于排查网络问题

```bash
kubectl run test-pod --rm -it --image=busybox --restart=Never -- sh
```
