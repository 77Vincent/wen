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

> 系统在正常运行时，不因缺陷而允许恶意行为损害相关人的利益。

## 可靠

> 系统不容易故障，故障了也容易快速修复。

系统的可靠程度最终通过可用性（Availability）量化，是指系统在指定时间段内（比如一年），正常运行时间的百分比。公式是

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

### 难出问题

### 坏了好修

## 物美

> 尽量高效迅速满足用户需求。

## 价廉

> 在满足用户需求的情况下，尽量少花钱。

## 卓越运营

> 让人可以持续高效使用并维护系统。

## 可持续发展

> 减少系统对环境的负面影响。
