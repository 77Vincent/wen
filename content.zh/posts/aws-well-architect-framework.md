---
title: AWS Well-Architected Framework
date: 2024-11-20T02:01:58+05:30
tags: [ computer-science, math ]
categories: study
canonicalUrl: https://wenstudy.com/posts/aws-well-architect-framework/
---

AWS Well-Architected Framework (AWS优良架构框架) 四个目标，可以用一个八字决来概括：*安全可靠，物美价廉*。

| Goal                     | 目标     |
|--------------------------|--------|
| Secure                   | **安全** |
| Resilient/Reliable       | **可靠** |
| High-Performing          | **物美** |
| Cost-Optimized/Efficient | **价廉** |

<!--more-->

## 安全
> 系统在正常运行时，不因缺陷而允许恶意行为损害相关人的利益。

### Shared Responsibility Model

_一句话概括，AWS负责云服务设施本身的安全，用户负责云上财产的安全。_

提到安全可靠，必然先涉及责任划分。这个模型的类比是，云是AWS盖好的精装房，用户是租户。房子的大门、用户自身的房门、窗、煤气管道等，安全质量由AWS保障。但用户可以忘记锁门、不关窗、不关煤气，导致自身财产损失。当然负责的用户不仅会关好门窗，还会加装额外的安全措施。

![image of shared responsibility model](/images/aws-well-architected-framework/aws-shared-responsibility-model.png "Shared Responsibility Model")

## 可靠
> 因各种原因导致故障，无法正常运行。

## 物美
> 在尽可能多的情况下，高效满足用户需求。

## 价廉
> 在满足用户需求的情况下，尽可能少的花费资源。
