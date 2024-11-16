---
title: 理解线性运算
date: 2024-11-16T02:01:58+05:30
tags: [math]
categories: study 
canonicalUrl: https://wenstudy.com/posts/understand-linear-operation/
math: true
---

<!-- more -->
线性运算是一种运算性质，其本质图像一定是一条线、平面等。必须满足两个条件：
1. 加法性
2. 齐次性

## 加法性 (Additivity)
即分开算和一起算的结果一样。

$$
f(x+y) = f(x) + f(y)
$$

## 齐次性 (Homogeneity)
即输入的倍数和输出的倍数一样。

$$
f(ax) = af(x)
$$

## 几种非线性运算
### 模数运算 (Modulus)
因为不满足加法性。
$$
\begin{align}
7 \mod 5 &= 2 \\
3 \mod 5 + 4 \mod 5 &= 7
\end{align}
$$

### 幂运算 (Power)
$$
f(x) = x^2
$$

因为不满足加法性:
$$
\begin{align}
f(2+3) &= 5^2 = 25 \\
f(2) + f(3) &= 4 + 9 = 13
\end{align}
$$
