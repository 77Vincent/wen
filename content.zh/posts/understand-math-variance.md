---
title: 理解方差
date: 2023-11-30T02:01:58+05:30
tags: [ probability, statistics, math ]
categories: study
canonicalUrl: https://wenstudy.com/posts/understand-math-variance/
math: true
---

方差（Variance）是衡量随机变量离散程度的指标。它是各个数据点与均值之差的平方和的均值。
<!--more-->

$$
Var(X) = E[(X - E[X])^2]
$$

## 等价公式
_方差也等于随机变量的平方的期望减去随机变量的期望的平方。_

$$
Var(X) = E[X^2] - E[X]^2
$$

因为：

$$
E[(X - E[X])^2] = E[X^2 - 2XE[X] + E[X]^2]
$$

根据期望的线性性质（即各部分可以分开算）：

$$
E[X^2 - 2XE[X] + E[X]^2] = E[X^2] - 2E[X]E[E[X]] + E[X]^2
$$

其中，\(E[X]\) 是常数，所以 \(E[E[X]] = E[X]\)。最终得到：

$$
E[X^2] - 2E[X]E[X] + E[X]^2 = E[X^2] - E[X]^2
$$

> 这个等价公式在计算一些分布的方差时会更方便。

## 数乘

对于任意常数 \(a\) 和随机变量 \(X\)，其对方差的变化会有平方的效果，本质原因是方差本身就带有平方。

$$
Var(aX) = a^2 Var(X)
$$

为什么？首先，由于期望的计算是线性的：

$$
E[aX] = aE[X]
$$

所以：

$$
\begin{align}
Var(aX) &= E[(aX - E[aX])^2] \\
&= E[(aX - aE(X))^2] \\
&= E[a^2(X - E(X)^2] \\
&= a^2E[(X - E(X))^2] \\
&= a^2Var(X)
\end{align}
$$

## 加法

对于两个随机变量 \(X\) 和 \(Y\)，有：

$$
Var(X + Y) = Var(X) + Var(Y) + 2Cov(X, Y)
$$

> 其中，\(Cov(X, Y)\) 是 \(X\) 和 \(Y\) 的协方差，表示 \(X\) 和 \(Y\) 的线性相关程度。

为什么？首先根据方差定义式展开：

$$
Var(X + Y) = E[(X + Y - E[X + Y])^2]
$$

根据期望的线性性质：

$$
E[X + Y] = E[X] + E[Y]
$$

所以：

$$
\begin{align}
Var(X + Y) &= E[(X + Y - E[X] - E[Y])^2] \\
&= E[((X - E[X]) + (Y - E[Y]))^2]
\end{align}
$$

展开平方项：

$$
Var(X + Y) = E[(X - E[X])^2 + 2(X - E[X])(Y - E[Y]) + (Y - E[Y])^2]
$$

依旧是期望的线性性质，可以分开计算：

$$
Var(X + Y) = E[(X - E[X])^2] + 2E[(X - E[X])(Y - E[Y])] + E[(Y - E[Y])^2]
$$

其中：
1. \(E[(X - E[X])^2] = Var(X)\)
2. \(E[(Y - E[Y])^2] = Var(Y)\)
3. \(E[(X - E[X])(Y - E[Y])] = Cov(X, Y)\) 这正是 \(X\) 和 \(Y\) 的协方差。

所以得到了：

$$
Var(X + Y) = Var(X) + Var(Y) + 2Cov(X, Y)
$$

如果 \(X\) 和 \(Y\) 是独立的，那么它们的协方差，即 \(Cov(X, Y)\) 为 0，则：

$$
Var(X + Y) = Var(X) + Var(Y)
$$
