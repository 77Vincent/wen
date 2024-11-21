---
title: 理解矩阵行列式
date: 2024-11-21T02:01:58+05:30
tags: [ matrix, math ]
categories: study
canonicalUrl: https://wenstudy.com/posts/matrix-determinant/
math: true
---

_行列式跟“式子“毫无关系，它是一个标量。_

行列式只适用于方形矩阵，描述了 n 个 n 维向量的空间体积。对二维而言，就是面积。行列式为零表示矩阵无法张成 n 维空间，最多 n-1 维，即不满秩。所以在 n 维上，体积为零。

> **矩阵的秩:** 指矩阵中的线性无关行或列的最大数目。表示矩阵在空间中所能描述的独立方向的数目，即矩阵实际能表达的维度上限。

本篇前提：[理解矩阵乘法](https://wenstudy.com/posts/matrix-multiplication/)

<!--more-->
## 写法
以二维空间为例，行列式的写法：

$$
\det(A) = |A| =
\begin{vmatrix}
a & c \\
b & d
\end{vmatrix}
$$


## 意义

这个二维矩阵（两个二维列向量）围成的面积为 12：

$$
\det(
\begin{bmatrix}
4 & 2 \\
2 & 4
\end{bmatrix}
) = 4*4 - 2*2 = 12
$$

![image of matrix-area](/images/matrix-determinant/2d-matrix-area.png "matrix-area")

而这个矩阵，显而易见，因为向量共线，行列式（面积）为零。

$$
\det(
\begin{bmatrix}
2 & -2 \\
4 & -4
\end{bmatrix}
) = 2*(-4) - 4*(-2) = 0
$$

![image of vector-collinear](/images/matrix-determinant/vector-collinear.png "vector collinear")

## 计算与特性
二维矩阵的行列式计算最简单，三维复杂度陡增，概念上理解二维即可。这个 \(2*2\) 矩阵里，按行或按列书写向量不重要，结果一样，比如用列向量：

$$
\begin{vmatrix}
a & c \\
b & d
\end{vmatrix} = ad - bc
$$

写成行向量，结果依旧：
$$
\begin{vmatrix}
a & b \\
c & d
\end{vmatrix} = ad - bc
$$

为什么这样计算？因为结果其实是下图平行四边形的面积。所以不论用行还是列向量书写，面积是不会变的。
$$
\begin{align}
\begin{vmatrix}
a & c \\
b & d
\end{vmatrix}
&= (a+c)(b+d) - ab - cd - 2bc \\
&= ab + ad + cb + dc - ab - cd - 2bc \\
&= ad - bc
\end{align}
$$

![image of 2d-matrix-determinant-calculation](/images/matrix-determinant/2d-matrix-determinant-calculation.png "2d-matrix-determinant-calculation")

如果交换两列，行列式的值会变号。_因为行列式的定义是向量围成的有向面（体）积，交换相当于把空间翻转了，导致变号。_

$$
\begin{vmatrix}
c & a \\
d & b
\end{vmatrix} = cb - ad
= -(ad - bc)
= -\begin{vmatrix}
a & c \\
b & d
\end{vmatrix}
$$

### 数乘
倍增一行（列），即倍增一个向量，行列式也会相应倍增。相当于推动了平行四边形的一个边（三维里是一个面）。

$$
\begin{vmatrix}
ka & c \\
kb & d
\end{vmatrix} = k\begin{vmatrix}
a & c \\
b & d
\end{vmatrix}
= k(ad - bc)
$$

![image of k-times-determinant](/images/matrix-determinant/k-times-determinant.png "k-times-determinant")

_但如果是对整个矩阵乘系数 \(k\)，行列式会倍增相应的 \(k^n\) 倍。因为记住，矩阵是有维度的。_

$$
\det(kA) = k^n \det(A)
$$

### 加法
给其中一个向量叠加另一个向量，行列式也会叠加。相当于以同一条边为基准，叠加了两个平行四边形的面积。

$$
\begin{vmatrix}
a + a' & c \\
b + b' & d
\end{vmatrix} = \begin{vmatrix}
a & c \\
b & d
\end{vmatrix} + \begin{vmatrix}
a' & c \\
b' & d
\end{vmatrix}
= ad - bc + a'd - b'c
$$

### 单位矩阵
单位矩阵的行列式为 1。最为优美，意味着它不会变换空间。

$$
\begin{vmatrix}
1 & 0 \\
0 & 1
\end{vmatrix} = 1 * 1 - 0 * 0 = 1
$$

## 行列式为零

行列式为零是一个重要的特性，代表矩阵无法张成 \(n\) 维空间，最多 \(n-1\) 维，所以在 \(n\) 维上体积为零，比如一个没有高的面或体，不管底边（面）有多大都白搭。其根本原因是存在线性相关的向量。即有一些向量是多余的，没有引入关于那个维度的独立的信息。

比如一个向量是另一个向量的倍数。
$$
\begin{vmatrix}
a & ka \\
b & kb
\end{vmatrix} = kab - kab = 0
$$

或存在零向量。这其实是上面 \(k=0\) 的特例。

$$
\begin{vmatrix}
a & 0 \\
b & 0
\end{vmatrix} = a*0 - 0*b = 0
$$

或者一个向量是另几个向量的线性组合，导致的结果依旧是第三个向量是不独立，多余的。这个等式用到了上述的加法和数乘性质。

$$
\begin{align}
\begin{vmatrix}
a & b & k_1a + k_2b \\
a_1 & b_1 &  k_1a_1 + k_2b_1 \\
a_2 & b_2 &  k_1a_2 + k_2b_2
\end{vmatrix} &= \begin{vmatrix}
a & b & k_1a \\
a_1 & b_1 &  k_1a_1 \\
a_2 & b_2 &  k_1a_2
\end{vmatrix} + \begin{vmatrix}
a & b & k_2b \\
a_1 & b_1 &  k_2b_1 \\
a_2 & b_2 &  k_2b_2
\end{vmatrix} = 0 + 0 = 0
\end{align}
$$

> \(k_1, k_2\) 是任意常数。

## 不可逆矩阵（奇异矩阵）
行列式为零的矩阵是不可逆矩阵，也叫奇异矩阵（Singular Matrix）。因为它不存在一个标准逆矩阵 \(A^{-1}\)，使得所有由它转变的输入都能逆向还原。

$$
A^{-1}(AB) = (AA^{-1})B = IB = B
$$

例如以下奇异矩阵 \(A\) 变换两个不同的输入向量，结果被映射到同一个向量。_多对一的情况下，逆向映射是不可能的。没有额外的信息告诉我们，该回到\(v\) 和 \(w\) 中的哪一个。_
$$
Av =
\begin{bmatrix}
1 & 1 \\
2 & 2
\end{bmatrix}\begin{bmatrix}
1 \\
0
\end{bmatrix} = v' = \begin{bmatrix}
1 \\
2
\end{bmatrix}
$$

$$
Aw =
\begin{bmatrix}
1 & 1 \\
2 & 2
\end{bmatrix}\begin{bmatrix}
0 \\
1
\end{bmatrix} = w' = \begin{bmatrix}
1 \\
2
\end{bmatrix}
$$
![image of many-to-one](/images/matrix-determinant/many-to-one.png "many-to-one-relation")

但凡修改一下 \(A\) 使其行列式不为零，就能逆向映射。因为此时，输入输出是一一对应的。
$$
A'v = \begin{bmatrix}
1 & 1 \\
2 & 3
\end{bmatrix}\begin{bmatrix}
1 \\
0
\end{bmatrix} = \begin{bmatrix}
1 \\
2
\end{bmatrix}
$$

$$
A'w = \begin{bmatrix}
1 & 1 \\
2 & 3
\end{bmatrix}\begin{bmatrix}
0 \\
1
\end{bmatrix} = \begin{bmatrix}
1 \\
3
\end{bmatrix}
$$

代数上，奇异矩阵没有逆矩阵 \(A^{-1}\) 的原因是，计算逆矩阵的公式里，行列式为零时除数为零，无法计算。细节下一篇再说。

$$
A^{-1} = \frac{1}{\det(A)}adj(A)
$$

## 相关阅读
- [线性代数整理(三)行列式特征值和特征向量](https://cloud.tencent.com/developer/article/1797038)
