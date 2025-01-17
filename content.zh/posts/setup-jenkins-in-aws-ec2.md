---
title: 用AWS EC2从零搭建Jenkins并实现GitOps CI/CD
date: 2025-01-16T02:01:58+05:30
tags: [ computer-science, aws, jenkins, cicd ]
categories: study
canonicalUrl: https://wenstudy.com/posts/setup-jenkins-in-aws-ec2/
---

上一篇：[用AWS EC2从零搭建Kubernetes和ArgoCD](/posts/setup-k8s-cluster-in-aws-ec2-without-using-eks/)

前一篇里，我们在 `AWS EC2` 上从无到有，搭建了 `Kubernetes` 集群和 `ArgoCD`。因为终究缺少 `CI` 环境以实现真正的实时持续集成，这一篇将手把手，在
`AWS EC2` 上搭建 `Jenkins` 实现一个 `GitOps CI/CD` 工作流。

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

## 预备代码仓库

### 应用仓库

假设我们已经有一个可以被 `Docker` 构建的极简 `hello-world HTTP` 服务应用，代码仓库地址为以下。

```
https://github.com/username/hello-world.git
```

> 这个仓库只管理应用代码，不包含部署配置。

### 部署配置仓库

我们还需要一个仓库用来管理用于 `Kubernetes/ArgoCD` 消费的部署配置文件，这个仓库地址为

```
https://github.com/username/hello-world-deployment.git
```

仓库的文件结构为

```
./staging/deployment.yml
./production/deployment.yml
```

> 这个仓库只管理部署配置，不包含应用代码。
>
> 两个 `manifest` 文件分别用于 `staging` 和 `production` 环境。

### 部署配置文件

以 `staging` 环境的 `deployment.yml` 文件为例。一般来说，这个文件包含以下几个部分，即以下几个 `Kind`：

| Section        | Description                                       |
|----------------|---------------------------------------------------|
| **Deployment** | 管理应用生命周期的资源，它定义了应用的副本数、容器镜像、更新策略等。                |
| **Service**    | 定义如何在集群内或外部访问 Pods。它充当了 Pods 的访问入口，能够提供负载均衡和端口映射。 |
| **Namespace**  | 用于资源逻辑隔离                                          |

以下是一个基础的 `deployment.yml` 文件。

```yaml
# Create a namespace
apiVersion: v1
kind: Namespace
metadata:
  name: YOUR_NAMESPACE
---
## define the deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: YOUR_DEPLOYMENT_NAME
  namespace: YOUR_NAMESPACE
  labels:
    app: YOUR_LABEL
spec:
  replicas: 2
  selector:
    matchLabels:
      app: YOUR_LABEL
  template:
    metadata:
      labels:
        app: YOUR_LABEL
    spec:
      imagePullSecrets:
        - name: ecr-pull-secret
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  - key: kubernetes.io/arch
                    operator: In
                    values:
                      - arm64
                      - amd64
      containers:
        - name: YOUR_CONTAINER_NAME
          image: YOUR_AWS_ACCOUNT_ID.dkr.ecr.ap-northeast-1.amazonaws.com/hello-world:latest
          ports:
            - name: http
              containerPort: 8080
          imagePullPolicy: Always
      nodeSelector:
        kubernetes.io/os: linux
---
## define the service
apiVersion: v1
kind: Service
metadata:
  name: YOUR_SERVICE_NAME
  namespace: YOUR_NAMESPACE
  labels:
    app: YOUR_LABEL
spec:
  selector:
    app: YOUR_LABEL
  ports:
    - port: 80
      targetPort: 8080
      nodePort: 30000
  type: NodePort

```

> **重点字段说明：**
>
> 请根据实际情况替换 `YOUR_` 开头的字段。
> 
> replica: 副本数，根据实际情况调整。
> 
> imagePullPolicy: 镜像拉取策略，`Always` 表示每次都拉取最新镜像，`IfNotPresent` 表示只有本地没有时才拉取。
> 
> nodePort: 用于外部访问的端口，即可以用 `http://<EC2-Public-IP>:30000` 访问服务。

### 关于 `imagePullSecrets`

配置中看到到 `imagePullSecrets` 字段，是用于从 ECR 拉取镜像时的校验，需要提前创建，这非常重要。可使用以下命令：

```bash
kubectl create secret docker-registry ecr-pull-secret \
    --docker-server=YOUR_AWS_ACCOUNT_ID.dkr.ecr.ap-northeast-1.amazonaws.com \
    --docker-username=AWS \
    --docker-password=$(aws ecr get-login-password --region ap-northeast-1) \
    --docker-email=none
```

> 其中 `ecr-pull-secret` 是一个自定义的名称，可以根据实际情况替换，但必须保证和 `deployment.yml` 中的一致。

## 配置 Pipeline

Pipeline的关键是 `Jenkinsfile`，它定义了整个 `CI/CD` 流程。通常对于 `staging`（预览环境）和 `production`（生产环境）有不同的流程。

| 环境             | 触发条件              | 镜像版本                   | 部署方式                   |
|----------------|-------------------|------------------------|------------------------|
| **staging**    | `push` 到目标分支      | 分支的 `HEAD commit hash` | 由CI自动触发 `ArgoCD` 的同步命令 |
| **production** | `tag` 并 `release` | 同 `tag`                | 人工手动在 `ArgoCD` 上触发同步   |

以下是针对 `staging` 环境的 `Jenkinsfile`。

```groovy
pipeline {
    agent any
    // 定义环境变量以便在整个流程中引用
    // 这里我们用 `AWS` 的 `ECR`（`Elastic Container Registry`）作为 `Docker` 镜像仓库。
    environment {
        ECR_REGISTRY = '733089366385.dkr.ecr.ap-northeast-1.amazonaws.com'
        ECR_REPOSITORY = 'hello-world'
        REGION = 'ap-northeast-1'
    }
    stages {
        stage('Checkout') {
            steps {
                // 从 GitHub 仓库拉取代码
                // 目标分支是 `main`
                git branch: 'main', url: 'https://github.com/username/hello-world.git',
                        // 用 github 的 `Personal Access Token` 作为凭证，最佳实践，比直接用密码或者 SSH 更安全。
                        credentialsId: 'fr-github'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // 用 `git` 的 `HEAD commit hash` 作为镜像版本
                    IMAGE_TAG = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
                    // 贴到环境变量中以便在后续步骤中引用
                    env.IMAGE_TAG = IMAGE_TAG
                }
                // docker 构建镜像
                sh '''
                    docker build -t ${ECR_REGISTRY}/${ECR_REPOSITORY}:${IMAGE_TAG} --build-arg ENVIRONMENT=staging .
                '''
            }
        }

        stage('Push to AWS ECR') {
            steps {
                // 用 AWS CLI 获取 ECR 登录密码
                // 并推送到 ECR 仓库
                sh '''
                    aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ECR_REGISTRY}
                    docker push ${ECR_REGISTRY}/${ECR_REPOSITORY}:${IMAGE_TAG}
                '''
            }
        }

        stage('Update Deployment Config') {
            steps {
                // 更新另一个 GitHub 仓库中的 `deployment.yml` 文件，这个文件是 `ArgoCD` 读取的用于部署的配置文件。
                // 将其中的 `image` 字段更新为最新的镜像版本。
                // 然后提交并推送到 GitHub 仓库。
                git branch: 'main', url: 'https://github.com/username/hello-world-deployment.git', credentialsId: 'fr-github'
                withCredentials([usernamePassword(credentialsId: 'fr-github', usernameVariable: 'GIT_USERNAME', passwordVariable: 'GIT_PASSWORD')]) {
                    sh '''
                        sed -i "s|image: .*|image: ${ECR_REGISTRY}/${ECR_REPOSITORY}:${IMAGE_TAG}|g" staging/deployment.yml
                        git config user.name "Jenkins CI"
                        git config user.email "jenkins@example.com"
                        git add staging/deployment.yml
                        if git diff-index --quiet HEAD; then
                            echo "No changes to commit."
                            exit 0
                        fi
                        git commit -m "Update staging deployment to ${IMAGE_TAG}"
                        git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/username/hello-world-deployment.git main
                    '''
                }
            }
        }

        // 最后触发 `ArgoCD` 同步命令。
        stage('Trigger ArgoCD Sync') {
            steps {
                sh '''
                    argocd login 13.231.229.5:32135 --username admin --password admin1234 --insecure
                    argocd app sync poc-staging
                '''
            }
        }
    }
}

```

配置好保存后，一个 `staging` 环墶的 `CI/CD` 流程就配置完成了。可以尝试手动 build 一次。预期结果是应用即刻用目标分支的最新代码构建并部署。

## 删除 Jenkins

有时因操作错误需重新安装 Jenkins，以下命令用于删除 Jenkins 以备不时之需。

```bash
sudo systemctl stop jenkins
sudo dnf remove jenkins -y
sudo rm -rf /var/lib/jenkins
sudo rm -rf /usr/share/jenkins
sudo rm -rf /etc/sysconfig/jenkins
```
