---
title: AWS Well-Architected Framework
date: 2024-11-20T02:01:58+05:30
tags: [ computer-science, aws ]
categories: study
math: true
canonicalUrl: https://wenstudy.com/posts/aws-well-architect-framework/
---

要用AWS搭建一个优秀的系统，根据官方资料和一些第三方讲义，我把要义总结为四个需求八个字：

*安全可靠，物美价廉*。

可以用买房子来类比，四个需求的层级是不一样的。除此之外，还有两朵锦上添花，最后再议。
<!--more-->
 ![image of system-requirement pyramid](/images/aws-well-architected-framework/aws-well-architected-framework.png "system-requirement-pyramid")

## Shared Responsibility Model

开始前，先要阐明一个AWS需要与用户达成的重要共识，叫做共同责任模型（Shared Responsibility Model）。

_一句话概括：AWS负责云服务设施本身的安全，用户负责云上财产的安全。_

这个模型的类比是，云是AWS盖好的精装房，用户是租户。房子的大门、用户自身的房门、窗、煤气管道等，安全质量由AWS保障。但用户可以忘记锁门、不关窗、不关煤气，导致自身财产损失。当然负责的用户不仅会关好门窗，还会加装额外的安全措施。

![image of shared responsibility model](/images/aws-well-architected-framework/aws-shared-responsibility-model.png "Shared Responsibility Model")

## 安全

_系统在正常运行时，不因缺陷而允许恶意行为损害相关人的利益。_

## 可靠

_系统不容易故障，故障了也容易快速修复。_

### 可用性
系统的可靠程度最终用可用性（Availability）量化，是指系统在指定时段内（比如一年），正常运行时间的百分比。公式是

$$
Availability = \frac{Uptime}{Total Time = (Uptime + Downtime)}
$$

| 年可用性    | 可用时间              |
|---------|-------------------|
| 99%     | 3 天, 15 小时, 39 分钟 |
| 99.9%   | 8 小时, 45 分钟       |
| 99.95%  | 4 小时, 22 分钟       |
| 99.99%  | 52 分钟             |
| 99.999% | 5 分钟              |

### 难出故障
#### 冗余
一个EC2的可用性是90%的话，即失败率是10%。那对于可分布式的服务来说，部属三个EC2，可用性将提高到99.9%。因为三个全部失败才能导致服务不可用。
$$
1 - 0.1^3 = 99.9%
$$

但如果你还有数据库、ALB，他们有各自的可用性，最终可用性将是各部分的乘积，会更低。例如其他两个的分别是 \(99.8%\)，那么整体可用性将是 \(99.5%\)，等于一年有将近两天不可用。

$$
99.9\% * 99.8\% * 99.8\% = 99.5\%
$$

> 热知识：Route 53的SLA是100%。

#### 弹性
EC2 auto-scaling 根据负载自动增减实例，是保障高可用的重要手段。

### 坏了好修

## 物美

_尽量高效迅速满足用户需求。_

## 价廉

_在满足用户需求的情况下，尽量少花钱。_

## 锦上添花

### 卓越运营

_让人可以持续高效使用并维护系统。_

### 可持续发展

_减少系统对环境的负面影响。_
