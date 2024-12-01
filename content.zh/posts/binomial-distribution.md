---
title: 理解二项分布
date: 2022-11-21T02:01:58+05:30
tags: [ probability, statistics, math ]
categories: study
canonicalUrl: https://wenstudy.com/posts/binomial-distribution/
math: true
---

二项分布几乎是一切的基础，它描述了在 n 次独立重复试验（伯努利实验）中，成功次数 k 的概率分布。
<!-- more -->

![binomial-distribution](/images/binomial-distribution/graph.png "binomial-distribution")

## 定义
对于 \(X ~ B(n, p)\)，即参数服从为 \(n\) 次试验，成功概率为 \(p\) 的二项分布。其概率质量函数为：

$$
P(X=k) = C(n, k) \cdot p^k \cdot (1-p)^{n-k}
$$

## 重复独立试验（伯努利实验）

是一种只有两种结果的随机试验，比如抛硬币，要么正要么反，要么输要么赢。且每次试验成功的概率是相同的，即前一次不影响下一次。

![coin-head-tail](/images/binomial-distribution/coin-head-tail.png "coin-head-tail")

## 概率质量函数（PMF）

_概率质量函数（Probability Mass Function, PMF）是离散随机变量在各个取值上的概率_。全部取值的概率之和为 1。

$$
\sum_{x=0}^{n} P(X) = 1\
$$

## 推导
成功概率为 \(p\)，失败概率为 \(1-p\)，现在做 \(n\) 次试验。

全成功的概率是：

$$
P(X=n) = p^n
$$

全失败（即成功次数为 \(0\)）的概率是：

$$
P(X=0) = (1-p)^n
$$

成功 \(1\) 次的概率是从 \(n\) 次里面选 1 次来成功。

$$
P(X=1) = C(n, 1) \cdot p \cdot (1-p)^{n-1}
$$

所以成功 \(k\) 次的概率就是从 \(n\) 次里面选 \(k\) 次来成功。

$$
P(X=k) = C(n, k) \cdot p^k \cdot (1-p)^{n-k}
$$

> 上面所用的组合数，是从 \(n\) 个元素中取 \(k\) 个元素的选法的总数，记作 \(C(n, k)\)。
$$
C(n, k) = \frac{n!}{k!(n-k)!}
$$
例如 \(n\) 个里选 \(1\) 个，显然总选法是 \(n\) 种，即 \(C(n, 1) = n\)。

## 期望（平均成功次数）

直觉上，实验次数乘以成功概率就是期望值。 为什么？对于单次伯努利实验，即每个 \(X_i\) 的期望值是 \(p\)。因为要么成功要么失败：

$$
E(X_i) = 1 \cdot p + 0 \cdot (1-p) = p
$$

利用期望的线性性质， \(n\) 次实验的期望值就是 \(np\)。

$$
E(X) = E(\sum_{i=1}^{n} X_i) = \sum_{i=1}^{n} E(X_i) = np
$$

## 方差（Variance）

方差描述成功次数围绕期望值 \(np\) 的波动程度，_定义是单个随机变量的平方期望值减去期望值的平方_。

$$
Var(X) = E(X^2) - E(X)^2
$$

二项分布的方差是：

$$
Var(X ~ B(n, p)) = np(1-p)
$$

为什么？对于单个伯努利变量（单次伯努利实验的结果）因为 \(X_i\) 只能是 \(0\) 或 \(1\)，所以：

$$
X_i^2 = X_i
$$

> 可以通过代入 \(X_i = 0\) 和 \(X_i = 1\) 来验证。

所以

$$
E(X_i^2) = E(X_i) = p
$$

由于单个变量的期望值已知，即 \(E(X_i) = p\)，所以单个变量的方差是：

$$
Var(X_i) = E(X_i^2) - E(X_i)^2 = p - p^2 = p(1-p)
$$

利用方差的线性性质， \(n\) 次实验的方差就是：

$$
Var(X) = Var(\sum_{i=1}^{n} X_i) = \sum_{i=1}^{n} Var(X_i) = np(1-p)
$$

> 一个重要的直觉是，当 \(p = 0.5\) 时，方差最大。显然，50% 成功率代表的不确定性最高。
> 
> 若成功率是 0% 或 100%，则不会有波动，即必然事件，没有随机性，所以方差为 0。

![square-area](/images/binomial-distribution/square-area.png "square-area")

## 标准差（Standard Deviation）

标准差是方差的平方根。所以显然：
$$
\sigma = \sqrt{np(1-p)}
$$
