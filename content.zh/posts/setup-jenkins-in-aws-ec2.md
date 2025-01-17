---
title: 用AWS EC2从零搭建Jenkins
date: 2025-01-16T02:01:58+05:30
tags: [ computer-science, aws, jenkins, cicd ]
categories: study
canonicalUrl: https://wenstudy.com/posts/setup-jenkins-in-aws-ec2/
---
 
上一篇：[用AWS EC2从零搭建Kubernetes和ArgoCD](/posts/setup-k8s-cluster-in-aws-ec2-without-using-eks/)

前一篇里，我们在 `AWS EC2` 上从无到有，搭建了 `Kubernetes` 集群和 `ArgoCD`。因为终究缺少 `CI` 环境以实现真正的实时持续集成，这一篇将手把手，在 `AWS EC2` 上搭建 `Jenkins` 实现一个 `GitOps CI/CD` 工作流。

<!--more-->

## 预备

### 准备 `EC2` 实例

使用 Amazon Linux 2 AMI。确保

1. 安全组（`Security Group`）开放 `8080` 端口。
2. EC2 的 IAM 角色有对 ECR 的读和写权限。一般是使用 `AmazonEC2ContainerRegistryFullAccess` 策略。

### `SSH` 登录 `EC2`

```bash
ssh -i <your-key.pem> ec2-user@<EC2-Public-IP>
```

### 更新包管理器 `dnf`

```bash
sudo dnf update -y
```

### 安装 `wget`

```bash
sudo dnf install wget -y
```

### 安装 `git`

```bash
sudo dnf install git -y
```

### 安装 `docker`

```bash
sudo dnf install docker -y
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker $USER
```

> 为了避免每次使用 `docker` 命令都需要 `sudo`，将当前用户添加到 `docker` 用户组。
> 
> 参考：[在AWS Linux EC2上准备Docker](/posts/managing-docker-on-aws-ec2/)

### 安装 `Java`

```bash
sudo dnf install java-17-amazon-corretto -y
sudo dnf install fontconfig -y
```

> 因为是在 AWS EC2 上安装，所以使用的是 Amazon Corretto JDK。
>
> fontconfig 是 Java 运行时所需的字体配置。

## 安装 Jenkins

### 添加 Jenkins 仓库

```bash
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo
sudo rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io-2023.key
```

### 安装

```bash
sudo dnf install jenkins -y 
```

### 启动服务

```bash
sudo systemctl start jenkins
sudo systemctl enable jenkins
```

### 查看服务状态

```bash
sudo systemctl status jenkins
```

应当看到类似如下输出，服务应当处于 `active (running)` 状态。

```
● jenkins.service - Jenkins Continuous Integration Server
     Loaded: loaded (/usr/lib/systemd/system/jenkins.service; disabled; preset: disabled)
     Active: active (running) since Thu 2025-01-16 09:13:46 UTC; 7s ago
   Main PID: 976120 (java)
      Tasks: 53 (limit: 4565)
     Memory: 591.4M
        CPU: 24.742s
     CGroup: /system.slice/jenkins.service
             └─976120 /usr/bin/java -Djava.awt.headless=true -jar /usr/share/java/jenkins.war --webroot=/var/cache/jenkins/war --httpPort=8080
```

### 添加 Jenkins 用户到 Docker 用户组

```bash
sudo usermod -aG docker jenkins
sudo systemctl restart jenkins
sudo systemctl restart docker
```

> 对于使用 Docker 进行镜像构建的 Jenkins 任务，这一步很重要。

## 配置 Jenkins

### 获取初始密码

```bash
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
```

### 访问并初始化

1. 在浏览器中输入 `http://<EC2-Public-IP>:8080`，然后输入初始密码，即可进入 Jenkins UI。
2. 选择安装推荐插件。
3. 创建管理员账户。
4. 配置 Jenkins URL。（默认即可）

完成后，进入 Jenkins 主界面。

![Jenkins Setup](/images/setup-jenkins-in-aws-ec2/jenkins-ui-init-plugin-install.png)

## 删除 Jenkins

有时因操作错误需重新安装 Jenkins，以下命令用于删除 Jenkins 以备不时之需。

```bash
sudo systemctl stop jenkins
sudo dnf remove jenkins -y
sudo rm -rf /var/lib/jenkins
sudo rm -rf /usr/share/jenkins
sudo rm -rf /etc/sysconfig/jenkins
```
