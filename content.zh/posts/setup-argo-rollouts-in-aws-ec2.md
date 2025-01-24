---
title: 在AWS EC2上集成Argo Rollouts实现Canary部署
date: 2025-01-23T02:01:58+05:30
tags: [ computer-science, aws, cicd, argocd, argo-rollouts, kubernetes, gitops ]
categories: study
canonicalUrl: https://wenstudy.com/posts/setup-argo-rollouts-in-aws-ec2/
---

上一篇：[用AWS EC2从零搭建Jenkins并实现GitOps CI/CD](/posts/setup-jenkins-in-aws-ec2/)

在上一篇里，我们在 `AWS EC2` 上从无到有，搭建了 `Jenkins` 并结合 `ArgoCD` 实现了一个 `GitOps CI/CD` 工作流。
然而 `ArgoCD` 本身只具备基础的滚动部署（`rolling update`），对于 `Canary` 部署支持并不友好。要实现全自动、智能化、省心的
`Canary` 部署策略，我们需要借助 `Argo Rollouts`。

<!--more-->

## 部署策略

一般来说部署策略有以下四种，金丝雀部署（`Canary`）是其中一种，也可谓是最高级，需要最多技术支撑的一种。

| 策略      | 描述                                          | 特点         |
|---------|---------------------------------------------|------------|
| **金丝雀** | 流量以线性或接近线性的方式，从老版本（blue）流到新版本（green）        | 较慢，低风险，高可用 |
| **蓝绿**  | 新（green）老（blue）版本共存但隔离，在老版本中完成测试，直接全面导流到新版本 | 较慢，低风险，高可用 |
| **滚动**  | 新版本逐渐替代老版本                                  | 较快，中风险，高可用 |
| **原地**  | 杀掉老版本，在原基建上启动新版本                            | 最快，高风险，低可用 |

> 金丝雀和蓝绿部署的差异在于，金丝雀部署是逐步导流，而蓝绿部署是一次性导流。

## Argo Rollouts

`Argo Rollouts` 是 `ArgoCD` 的一个子项目，它提供了更高级的部署策略支持，包括 `Canary` 部署、蓝绿部署等。

### 安装

```bash
kubectl create namespace argo-rollouts
kubectl apply -n argo-rollouts -f https://raw.githubusercontent.com/argoproj/argo-rollouts/stable/manifests/install.yaml
```

### 部署配置

在应用部署的配置文件里，原本是 `Deployment`，现在改为 `Rollout`。并添加 `strategy` 字段，指定部署策略。这个例子里，使用了
`Canary` 部署策略。导流将会分为三个阶段，每个阶段都会执行一个 `analysis` 模板。只有在 `analysis` 模板通过后，才会继续下一个阶段。
及流量先从 `0%` 到 `25%`，然后到 `50%`，最后到 `100%`。

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Rollout
metadata:
  name: <rollout-name>
  namespace: <namespace>
  labels:
    app: <app-name>
spec:
  replicas: 10
  strategy:
    canary:
      steps:
        - setWeight: 25
        - analysis:
            templates:
              - templateName: <template-name>
        - setWeight: 50
        - analysis:
            templates:
              - templateName: <template-name>
```

> 这里 25%、50% 等流量百分比，是通过控制容器数量来控制的。例如，目标（已有）是10个容器，要实现 25%
> 的流量在新版本里，就要部署3个新版本容器并杀掉1个旧版本，此时共 12 个容器里，3 个新容器将分的 25% 的流量。
> Argo Rollouts 会自动计算数量并执行。

### analysis 模板

`analysis` 模板是一个 `Job`，用于执行一些测试，例如 `Prometheus` 指标、`HTTP` 请求等。只有当 `analysis` 模板通过后。

一个极简的演示例子是，每五秒，连续三次成功请求某网址，才算通过。

```yaml
apiVersion: argoproj.io/v1alpha1
kind: AnalysisTemplate
metadata:
  name: <template-name>
spec:
  metrics:
    - name: success-rate
      interval: 5s
      count: 3
      successCondition: result == 200
      provider:
        web:
          url: https://www.google.com
```

### 指标分析

我的项目里的实践是根据 `Datadog` 的指标来判断是否通过。集成配置如下：

```yaml
## ...
spec:
  metrics:
    - name: <metric-name>
      successCondition: default(result, 0) < 1
      failureLimit: 1
      interval: 10s
      count: 2
      provider:
        datadog:
          apiVersion: v2
          interval: 5m
          queries:
            a: sum:trace.http.request.errors{service:argo-dummy-service}.as_count()
            b: sum:trace.http.request.hits{service:argo-dummy-service}.as_count()
          formula: "(moving_rollup(a, 60, 'sum') / moving_rollup(b, 60, 'sum')) * 100"
```

> 这个配置代表，每 10 秒，连续 2 次，`Datadog` 的 `argo-dummy-service` 服务的错误率小于 1% 才算通过。

### 部署

准备好后只需要部署以上配置（通过UI同步或命令行工具部署），下次发布新版本时，`Argo Rollouts` 就会全权接管发布的过程。

## Hub-And-Spoke
