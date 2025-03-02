---
title: 用AWS EC2从零搭建Kubernetes和ArgoCD
date: 2025-01-10T02:01:58+05:30
tags: [ computer-science, aws, kubernetes, argocd, cicd ]
categories: study
canonicalUrl: https://wenstudy.com/posts/setup-k8s-cluster-in-aws-ec2-without-using-eks/
---

在AWS上实现 Kubernetes 集群最简单的方法是走 EKS（Elastic Kubernetes Service）托管服务（managed service）。但后来发现 EKS
成本高昂，会有一笔跟 EC2 计算费用无关的起步价，仅仅来源于 EKS。为了学习（省钱），我们来用裸机 EC2 实例搭建 Kubernetes 集群，以及集成
ArgoCD 实现 CD (Continuous Deployment)。

<!--more-->

## 准备EC2实例

通过AWS控制台或者AWS CLI创建EC2实例。一般Production环境配置至少两个节点，一个作为Master节点，另一个作为Worker节点。实例最低需求：

| 节点类型   | vCPU | 内存  | 磁盘容量 | 用途                          |
|--------|------|-----|------|-----------------------------|
| Master | 2    | 4GB | 20GB | 用于 Kubernetes control-plane |
| Worker | 1    | 2GB | 20GB | 用于运行应用程序                    |

> **也可以单节点部署，即Master和Worker放在一个EC2上。本教程会使用单节点部署，节约成本，用于学习。**

使用AWS AMI创建EC2实例，选择 Amazon Linux xxx AMI。

![选择Amazon Linux AMI](/images/setup-k8s-cluster-in-aws-ec2-without-using-eks/aws-ec2-launch-ami.png "选择Amazon Linux AMI")

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

### 安装

```bash
sudo dnf install -y containerd
```

### 启动

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now containerd
```

### 创建默认配置文件

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

`CNI` (`Container Network Interface`) 插件为 `Kubernetes` 提供网络功能。它主要负责：

1. 为 `Pod` 分配 `IP` 地址。
2. 设置容器之间的网络通信。
3. 确保 `Pod` 可以与其他 `Pod`、服务（`Service`）以及外部世界通信。

如果没有正确安装和配置 `CNI` 插件，`Kubernetes` 的 `Pod` 网络将无法正常工作，导致 `Pod` 无法互通或无法分配 `IP`
地址。先查看是否已经安装了
`CNI` 插件

```bash
ls /opt/cni/bin
```

如果以上命令没有输出如下信息，说明没有安装 `CNI` 插件

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
bandwidth  bridge  dhcp  dummy  firewall  flannel  host-device  host-local
ipvlan  loopback  macvlan  portmap  ptp  sbr  static  tap  tuning  vlan  vrf
```

## 其他必要配置

### 启动两个内核模块

```bash
sudo modprobe overlay  
sudo modprobe br_netfilter
```

> `overlay` 和 `br_netfilter` 是 `Kubernetes` 集群所必需的两个内核模块，需要手动加载。

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

### 修改 `sysctl` 配置

```bash
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF
```

> 这个操作使得 `iptables` 能够正确工作。

刷新 `sysctl` 配置以使其生效

```bash
sudo sysctl --system
```

### 禁用 `swap`

```bash
sudo swapon -s # 查看 swap 分区
sudo swapoff -a
```

> 原因是 Kubernetes 不支持 swap 分区，因为 swap 分区会导致 Pod 的内存限制无效。

## 安装 Kubernetes

### 安装 `curl`

```bash
sudo dnf install -y curl
```

> 已有了 `curl` 的话则跳过这一步。

### 添加 `Kubernetes` 仓库地址

以下命令添加了 Kubernetes 仓库的地址到 `/etc/yum.repos.d/kubernetes.repo` 文件中

```bash
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://pkgs.k8s.io/core:/stable:/v1.28/rpm/
enabled=1
gpgcheck=1
gpgkey=https://pkgs.k8s.io/core:/stable:/v1.28/rpm/repodata/repomd.xml.key
exclude=kubelet kubeadm kubectl cri-tools kubernetes-cni
EOF
```

更新包管理器的缓存让其识别新的仓库

```bash
sudo dnf makecache
```

### 安装

安装 `kubeadm`, `kubelet` 和 `kubectl` 并启动 `kubelet` 服务

```bash
sudo dnf install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
sudo systemctl enable --now kubelet
```

> kubelet 服务是 Kubernetes 的主要组件之一，它负责管理 Pod 的生命周期，包括创建、销毁、监控 Pod 等。
>
> kubeadm 是 Kubernetes 的初始化工具，它可以帮助我们快速初始化一个 Kubernetes 集群。
>
> kubectl 是 Kubernetes 的命令行工具，用于与 Kubernetes 集群交互。

安装 `tc (Traffic Control)` 包，以便 `Kubernetes` 集群能够正常工作，首先查看包含 `tc` 命令的软件包

```bash
dnf provides tc
```

一般都存在于 `iproute-tc` 软件包中，因此安装 iproute-tc 软件包

```bash
sudo dnf install -y iproute-tc
```

### 启动

通过 `kubeadm` 初始化 `Kubernetes` 集群

```bash
sudo kubeadm init --pod-network-cidr=10.244.0.0/16
```

> 一定要用 `sudo` 执行 `kubeadm init` 命令，因为这个命令会修改系统的配置文件，否则会报错。
>
> 这里 CIDR 使用 10.244.0.0/16，因为稍后配置的网络插件 Flannel 默认使用这个 CIDR。


初始化完成后，会输出类似如下的信息，其中有两个命令，一个是 `kubeadm join` 命令，另一个是 `kubectl apply` 命令，分别用于加入节点和安装网络插件。

```
Your Kubernetes control-plane has initialized successfully!
```

### 检查 `kubelet` 服务状态

```bash
systemctl status kubelet
```

应得到如下输出，kubelet 服务应该是 `Active: active (running)` 状态

```
● kubelet.service - Kubernetes Kubelet
   Loaded: loaded (/usr/lib/systemd/system/kubelet.service; enabled; vendor preset: disabled)
   Active: active (running) since Fri 2024-11-22 09:00:00 UTC; 1min 30s ago
     Docs: https://kubernetes.io/docs/
 Main PID: 12345 (kubelet)
    Tasks: 123
   Memory: 123.4M
   CGroup: /system.slice/kubelet.service
           └─12345 /usr/bin/kubelet --config=/var/lib/kubelet/config.yaml --kubeconfig=/var/lib/kubelet/kubeconfig --network-plugin=cni --pod-infra-container-image=k8s.gcr.io/pause:3.5.1 --resolv-conf=/etc/resolv.conf
```

将 `kubectl` 配置文件拷贝到当前用户的家目录下

```bash
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

### 移除 `taint`（可选）
如果是单节点集群，那么 control-plane 节点也需要运行普通 Pod，即 argocd 相关 Pod。然而 control-plane 节点一般带有
`node-role.kubernetes.io/control-plane` taint，这会导致普通 Pod 无法调度到 control-plane 节点上。
这里提供用于测试的快速解决方案，即移除 control-plane 节点上的 taint，使其可以运行普通 Pod。

```bash
kubectl taint nodes --all node-role.kubernetes.io/control-plane-
```

## 安装 CNI 插件

安装CNI网络插件，这里使用 `Flannel`。Flannel 是一种简单的 `Kubernetes` 网络解决方案。它会为每个 `Pod` 分配一个唯一的 `IP`
地址，并确保不同节点之间的 `Pod` 能通过虚拟网络通信。

### 部署 `Flannel`

```bash
kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml
```

### 检查服务状态

```bash
kubectl get pods -n kube-flannel
```

应得到如下输出，flannel的Pod应该是 `Running` 状态

```
NAME                    READY   STATUS    RESTARTS       AGE
kube-flannel-ds-n2nzv   1/1     Running   20 (52s ago)   67m
```

以及用以下命令查看 `Flannel DaemonSet`

```bash
kubectl get daemonset kube-flannel-ds -n kube-flannel
```

应得到如下输出，`AVAILABLE` 列应该是 `1`

```
NAME             DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
kube-flannel-ds  1         1         1       1            1           <none>                   67m
```

查看系统相关的 `Pod`

```bash
kubectl get pods -n kube-system
```

应得到如下输出，所有的 Pod 都应该是 `Running` 状态，如果有 `ContainerCreating` 状态，等一下再查看。

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
ip-123-12-12-12.ap-northeast-1.compute.internal   Ready    control-plane   72m   v1.28.0
```

## 安装 `Metrics Server`（可选）

`Metrics Server` 是 `Kubernetes` 的一个聚合器，用于收集集群中的资源使用情况。有了 `Metrics Server`，就可以使用
`kubectl top`
命令查看集群实时的资源使用情况。以下先部署：

```bash
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
``` 

编辑 `Metrics Server` 的 `Deployment`，添加 `--kubelet-insecure-tls` 参数

```bash
kubectl edit deployment metrics-server -n kube-system
```

找到 `containers` 下的 `args` 字段，添加 `--kubelet-insecure-tls` 参数，如下，保存退出后，`Deployment` 会自动更新。

```yaml
containers:
  - args:
      - --kubelet-insecure-tls
```

> 这个操作的原因是，Metrics Server 默认使用自签名证书，而 Kubernetes API Server 可能拒绝与其通信。

检查 `Metrics Server` 是否正常运行

```bash
kubectl get pods -n kube-system | grep metrics-server
```

应得到如下输出，`metrics-server` 的 `Pod` 应该是 `Running` 状态

```
metrics-server-6d4c8c5f99-7z5zv   1/1     Running   0          3m
```

于是可以用 `kubectl top` 命令查看集群实时的资源使用情况

```bash
kubectl top nodes
kubectl top pods -n <namespace>
```

## 集成 `ArgoCD`

`ArgoCD` 是一个用于 `GitOps` 部署的工具，它可以帮助我们将应用程序的配置文件存储在 `Git` 仓库中，并通过 `ArgoCD` 自动同步到
`Kubernetes` 集群中。

### 创建命名空间

```bash
kubectl create namespace argocd
```

### 安装 `ArgoCD`

```bash
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

### 查看相关 `Pod` 状态

```bash
kubectl get pods -n argocd
```

所有的 Pod 都应该是 `Running` 状态，同样，有 `Pod` 处于 `ContainerCreating` 或 `Init` 状态，等一下再查看。

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

### 暴露 `ArgoCD Server`

```bash
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'
```

查看 `ArgoCD Server` 的 `NodePort` 端口

```bash
kubectl get svc argocd-server -n argocd
```

应得到如下输出

```
NAME            TYPE       CLUSTER-IP    EXTERNAL-IP   PORT(S)                      AGE
argocd-server   NodePort   10.97.82.16   <none>        80:31027/TCP,443:32130/TCP   14m
```

> 其中 `80:31027/TCP` 是 HTTP 端口，`443:32130/TCP` 是 HTTPS 端口。可以通过 `http://<control-plane-ip>:31027` 访问 ArgoCD
> UI。

获取 `control-plane` 节点的公网 `IP` 地址，然后访问 `http://<control-plane-ip>:31027`。对于 `AWS EC2` 实例，可以在 `AWS`
控制台中查看。

![ArgoCD 登陆界面](/images/setup-k8s-cluster-in-aws-ec2-without-using-eks/argocd-login-page.png "ArgoCD 登陆界面")

### 获取 `ArgoCD Server` 初始密码

```bash
kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d
```

> 用户名为 `admin`，密码为上面的命令输出。建议在登陆后立即修改管理员密码。

## 安装 `ArgoCD CLI`

后续为了与 `CI` 集成实现自动化部署，`ArgoCD CLI` （命令行工具）是必不可少的。

### 下载

```bash
curl -sSL -o argocd https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64
```

### 赋予执行权限

```bash
chmod +x argocd
```

### 移动可执行文件

将 `argocd` 移动到 `/usr/local/bin` 目录下

```bash
sudo mv ./argocd /usr/local/bin
```

### 验证安装

```bash
argocd version
```

`Kubernetes` 以及 `ArgoCD` 集成完成。接下去就是通过 `ArgoCD` 部署应用程序了。

### 配置

```bash
argocd login <ARGOCD_SERVER> --username admin --password <YOUR_PASSWORD> --insecure
```

> `ARGOCD_SERVER` 是 `ArgoCD` 服务器的地址，`YOUR_PASSWORD` 是 `ArgoCD` 管理员的密码。
> 
> `ArgoCD` 的服务器地址可以通过 `kubectl get svc argocd-server -n argocd` 这个命令查看，选择其中的 `EXTERNAL-IP` 地址。
>
> `--insecure` 参数是因为 `ArgoCD` 默认使用自签名证书，而 `ArgoCD CLI` 可能拒绝与其通信。

### 验证连接

```bash
argocd app list
```

到这里，`Kubernetes` 与 `ArgoCD` 的部署和集成完成，可以开始部署应用程序了。

由于 `ArgoCD` 其实只负责 `CD`部分（显而易见），应用的打包构建还是没有被自动化。尤其对于预览和测试环境来说，我们通常希望拥有实时的构建和部署。因此我们还是需要搭建
`CI` 环境。

请参见下一篇：[用AWS EC2从零搭建Jenkins](/posts/setup-jenkins-in-aws-ec2)
