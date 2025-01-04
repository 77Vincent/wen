---
title: 随机过程中的碰撞次数
description: 通过学习并运用泊松过程、柏松分布，指数分布，来计算随机事件情况下的碰撞次数的期望，以及推导过程，解决真实世界中排他性事件处理的问题。
date: 2024-11-12T02:01:58+05:30
tags: [math, randomness]
categories: study 
math: true
canonicalUrl: https://wenstudy.com/posts/poisson-process-collision-rate/
---

<!-- more -->
一个网站每小时平均有1000个请求，一次只能处理一个，用时50毫秒。当一个请求到达时，若前一个还未完成，就会报错。那么在一个小时内，平均会报多少次错？

## 排队论（Queueing Theory）
这是一个经典的 M/M/1 队列问题，通过计算系统利用率，可以得到拒绝率。系统利用率是：

$$
\rho = \frac{\lambda}{\mu}
$$

其中，\(\lambda\) 是到达率，等于单位时间内到达的请求数，即每秒 \(1000/3600\) 个：

$$
\lambda = \frac{1000}{3600} \approx 0.2778
$$

服务时间是50毫秒，即每秒可以处理 \(1/0.05 = 20\) 个请求， 所以服务率是：

$$
\mu = 20
$$

在 M/M/1 队列中，无等待位的情况下，系统利用率等于拒绝率（因为利用中的时候就会拒绝新请求），所以拒绝率是：

$$
\rho = \frac{\lambda}{\mu} = \frac{0.2778}{20} = 0.0139
$$

平均每小时有1000个请求，所以拒绝的请求数是：

$$
1000 \times 0.0139 = 13.9
$$

这个结果用柏松过程一样可以得到。

## 泊松过程（Poisson Process）
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
E(X) &= 请求总数 \times P(X \geq 1) \\
&= T\lambda \times (1 - e^{-\lambda t})
\end{align}
$$

其中，\(T\) 是时间单位，\(\lambda\) 是发生速率。

可以这样计算的原因是，每次事件都有同等概率和前一个事件碰撞（源于指数分布的无记忆性），而这样产生碰撞的机会共有 \(T\lambda-1\) 次。

## 结论
回到开头的问题，一个请求的持续时间是50毫秒，每小时1000个请求，一个小时3600秒，对于每个请求和前一个请求碰撞的概率是：

$$
\lambda = 1 - e^{-\frac{1000}{3600} \times 0.05} \approx 0.0138
$$

一小时的请求总数是1000，所以碰撞次数的期望是：

$$
E(X) = 1000 \times (1 - e^{-\frac{1000}{3600} \times 0.05}) \approx 13.8
$$
