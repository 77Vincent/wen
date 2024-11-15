---
title: 随机过程中的碰撞次数
description: 通过学习并运用泊松过程、柏松分布，指数分布，来计算随机事件情况下的碰撞次数的期望，以及推导过程，解决真实世界中排他性事件处理的问题。
date: 2024-11-12T02:01:58+05:30
tags: [math, randomness]
categories: study 
math: true
canonicalUrl: https://wenstudy.com/posts/poisson-process-collision-rate/
---

这可能是你也会遇到的问题，我找到了一个简单的计算方法。从结果来看，应该是正确的。
<!-- more -->
一个网站每小时平均有1000个请求，一次只能处理一个，用时1秒。当一个请求到达时，若前一个还未完成，就会报错。那么在一个小时内，平均会报多少次错？

通过把网站流量看作泊松过程，可以计算出碰撞概率。

## 泊松过程
泊松过程是一个随机过程，其中事件以恒定的速率发生：
1. 单位时间内事件发生次数服从泊松分布。
![image of poisson distribution](/images/poisson-process-collision-rate/poisson-distribution.png "poisson distribution")

2. 时间间隔服从指数分布。
![image of exponential distribution](/images/poisson-process-collision-rate/exponential-distribution.png "exponential distribution")

## 推导
知名的二项分布，\(P(X=k)\)表示 \(n\) 次试验中成功 \(k\) 次的概率。
$$
P(X=k) = C_n^k p^k (1-p)^{n-k}
$$

其极限表达式等于泊松分布的概率质量函数(PMF)，因为源于二项分布，所以也是离散的。其意义是发生速率恒定的情况下， 单位时间内事件发生 \(k\) 次的概率。
$$
\lim_{n \to \infty} C_n^k p^k (1-p)^{n-k} = \frac{\lambda^k e^{-\lambda}}{k!}
$$

显然，单位时间内，一次事件也没发生的概率是
$$
P(X=0) = \frac{\lambda^0 e^{-\lambda}}{0!} = e^{-\lambda}
$$

在\(t\)个单位时间内，都没发生的概率是（即 \(P(X=0)\)的 \(t\) 次连乘）
$$
P(X=0) = e^{-\lambda t}
$$

显然，在\(t\)个单位时间内，至少发生一次的概率是
$$
P(X \geq 1) = 1 - e^{-\lambda t}
$$

## 碰撞次数
真实世界里的碰撞，就是上一个事件发生后，在其还未结束时，下一个事件就发生了。可以如下计算碰撞的期望
$$
\begin{align}
E(X) &= (请求总数-1) \times P(X \geq 1) \\
&= (T\lambda-1) \times (1 - e^{-\lambda t})
\end{align}
$$

其中，\(T\) 是时间单位，\(\lambda\) 是发生速率。

可以这样计算的原因是，每次事件都有同等概率和前一个事件碰撞（源于指数分布的无记忆性），而这样产生碰撞的机会共有 \(T\lambda-1\) 次。

> 减1是因为第一个事件没有前一个事件与其碰撞。

## 结论
回到开头的问题，一个请求的持续时间是1秒，一小时，3600秒内平均有1000个请求，所以发生碰撞的期望是

$$
E(X) = (1000 - 1) \times (1 - e^{-1000/3600}) \approx 0.632
$$

同理易得，一天内发生碰撞的期望是
$$
E(X) = (1000 \times 24 - 1) \times (1 - e^{-1000/3600 \times 24}) \approx 15.17
$$
