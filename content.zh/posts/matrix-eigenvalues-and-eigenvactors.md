---
title: 理解矩阵特征值和特征向量
date: 2024-11-25T02:01:58+05:30
tags: [ matrix, math ]
categories: study
canonicalUrl: https://wenstudy.com/posts/matrix-eigenvalues-and-eigenvactors.md/
math: true
---

矩阵的根本属性就是变换空间，所以 _特征值描述其对某个方向上的缩放力度。特征向量描述那个方向。_

本篇前提：[理解矩阵行列式](https://wenstudy.com/posts/matrix-determinant/)

## 意义
特征向量是一个非零向量\(v\)，其被矩阵 \(A\) 变换后，方向不变，仅伸长或缩短，缩放倍数 \(\lambda\) 就是特征值。

$$
A \cdot v = \lambda \cdot v
$$

![image of matrix-eigen](/images/matrix-eigenvalues-and-eigenvectors/matrix-eigen.png "matrix-eigen")

## 方阵！
矩阵必须是方形矩阵，即 \(n \times n\)，并且特征向量也是一个 \(n\) 维向量。**因为只有在输入和输出维度相同的情况下，才能讨论一个向量在变换后方向不变的问题。**

## 特征值（Eigen Values）
根据定义，有如下推导：
$$
A \cdot v - \lambda \cdot v = (A - \lambda \cdot I) \cdot v = 0
$$

得到新的同质矩阵 \(M = A - \lambda \cdot I\) 乘以向量 \(v\) 得到零向量。

$$
M \cdot v = 0
$$

因为 \(v\) 是非零向量，所以矩阵 \(M\) 必须是奇异矩阵，**因为只有奇异矩阵才有可能将一个非零向量变换为零向量**。那么 \(M\) 的行列式为零，于是得到特征方程：

$$
\det(A - \lambda \cdot I) = 0
$$

> 相反的，在 \(n\) 维上的非奇异矩阵一定可以将 \(n\) 维空间中的任意非零向量变换为另一个非零向量。

\(A - \lambda \cdot I\) 表示从矩阵 \(A\) 的对角线元素中减去特征值 \(\lambda\)。例如对 \(3 \times 3\) 矩阵来说：

$$
\lambda \cdot I = \begin{bmatrix}
\lambda & 0 & 0 \\
0 & \lambda & 0 \\
0 & 0 & \lambda
\end{bmatrix}
$$

例如对于二维矩阵 \(\begin{bmatrix} 4 & 1 \\ 2 & 3 \end{bmatrix}\) 来说，若行列式为零，则：

$$
\begin{align}
\det(\begin{bmatrix} 4 - \lambda & 1 \\ 2 & 3 - \lambda \end{bmatrix}) = (4 - \lambda)(3 - \lambda) - 2 \\
= \lambda^2 - 7\lambda + 10 = 0
\end{align}
$$

于是得到两个特征值：\(\lambda_1 = 2, \lambda_2 = 5\)。

## 特征向量（Eigen Vectors）
在得到特征值后，将其代入特征方程，将得到特征向量。

$$
(A - \lambda \cdot I) \cdot v = 0
$$

首先，对于 \(\lambda_1 = 2\)，代入得到：

$$
\begin{align}
A - 2 \cdot I = \begin{bmatrix} 4 - 2 & 1 \\ 2 & 3 - 2 \end{bmatrix} = \begin{bmatrix} 2 & 1 \\ 2 & 1 \end{bmatrix}
\end{align}
$$

将其转换为增广矩阵，得到：

$$
\begin{bmatrix} 2 & 1 \\ 2 & 1 \end{bmatrix} \begin{bmatrix} x \\ y \end{bmatrix} = \begin{bmatrix} 0 \\ 0 \end{bmatrix}
$$

即：

$$
\begin{align}
2x + y = 0 \\
2x + y = 0
\end{align}
$$

在这个方向上，存在无数可选的向量，但一般我们选择单位向量 （即各，即 \(x = 1\)，得到 \(y = -2\)，所以特征向量 \(v_1\) 为：

$$
\begin{bmatrix} 1 \\ -2 \end{bmatrix}
$$

同理，再代入 \(\lambda_2 = 5\)，得到：

$$
\begin{align}
A - 5 \cdot I = \begin{bmatrix} 4 - 5 & 1 \\ 2 & 3 - 5 \end{bmatrix} = \begin{bmatrix} -1 & 1 \\ 2 & -2 \end{bmatrix}
\end{align}
$$

最终得到特征向量 \(v_2\) 为：\(\begin{bmatrix} 1 \\ 1 \end{bmatrix}\)

这个结果说明，矩阵 \(A\) 的"变换能力"是：在 \(v_1 = \begin{bmatrix} 1 \\ -2 \end{bmatrix}\) 方向上伸展 \(2\) 倍，在 \(v_2 = \begin{bmatrix} 1 \\ 1 \end{bmatrix}\) 方向上伸展 \(5\) 倍。

## 为何有多个值？

因为例子的矩阵 \(A\) 是二维矩阵，必定有两个方向的伸缩。对于 \(n\) 维矩阵来说，就有 \(n\) 个特征值和特征向量。
