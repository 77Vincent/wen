---
title: 在AWS Linux EC2上准备Docker
date: 2024-11-20T02:01:58+05:30
tags: [ computer-science, aws, kubernetes ]
categories: study
canonicalUrl: https://wenstudy.com/posts/managing-docker-on-aws-ec2/
---

对于简单的容器化应用的部署，Docker 和 EC2 的搭配是最常见的选择。先创建EC2实例，这里使用 Amazon Linux 2 AMI，因此包管理器是 dnf。首先更新 dnf 包管理器。

```bash
sudo dnf update -y
```

## 安装
安装 Docker
```bash
sudo dnf install -y docker
```

在系统启动时启动 Docker 服务
```bash
sudo systemctl enable docker
```

启动 Docker 服务
```bash
sudo systemctl start docker
```

检查 Docker 服务状态
```bash
sudo systemctl status docker
```

或者检查 Docker 版本
```bash
docker --version
```

把当前用户（一般是 `ec2-user`）加入 docker 用户组，以便不用 sudo 就可以运行 docker 命令
```bash
sudo usermod -aG docker ec2-user
```

测试 Docker 是否安装成功，通过运行 hello-world 镜像
```bash
docker run hello-world
```

## 删除

先停止 Docker 服务（如果正在运行）
```bash
sudo systemctl stop docker
sudo systemctl stop docker.socket
```

禁止 Docker 服务开机启动
```bsah
sudo systemctl disable docker
```

查看已安装的 Docker 或者 Docker 相关的包
```bash
sudo dnf list installed | grep docker
```

根据上面的输出，删除 Docker 或者 Docker 相关的包
```bash
sudo dnf remove -y docker
```

删除 Docker 数据（默认在 `/var/lib/docker`）
```bash
sudo rm -rf /var/lib/docker
```

删除 Docker 配置文件
```bash
sudo rm -rf /etc/docker
```

删除 Docker 创建的网络
```bash
sudo ip link delete docker0
```

> Docker 安装时会创建默认的桥接网络接口 docker0，需要手动删除。
