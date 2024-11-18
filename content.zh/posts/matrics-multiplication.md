---
title: 理解矩阵乘法
date: 2024-03-18T02:01:58+05:30
tags: [matrix, math]
categories: study 
canonicalUrl: https://wenstudy.com/posts/matrics-multiplication/
math: true
---

矩阵乘法分为三种情况：点积、矩阵乘向量、矩阵乘矩阵。这三种情况的本质是一样的，都是向量之间的运算。只是维度不同，因此运算的次数不同。
<!--more-->

## 点积
点击是两个向量之间的运算，结果是一个标量。 点积可以理解为两个向量“合并”后可以走多远，单纯的模长乘积是不对的，因为向量带方向。因此为了共线，就需要乘以夹角的余弦值。

$$
\begin{bmatrix}
x_1 & y_1
\end{bmatrix}
\begin{bmatrix}
x_2 \\
y_2
\end{bmatrix}
= x_1x_2 + y_1y_2
= |A||B|cos\theta
$$

更广泛的：
$$
a \cdot b = \Sigma_{i=1}^{n} a_ib_i
$$

## 矩阵乘向量

横着看，因为是二维空间，矩阵里每一行代表一个维度，\(x_1, y_1\) 告诉目标分量 \(x\) 如何因两个维度而变化。第二个分量 \(y\) 同理。更高维度亦同理。

$$
\begin{bmatrix}
x_1 & y_1 \\
x_2 & y_2
\end{bmatrix}
\begin{bmatrix}
x \\
y
\end{bmatrix}
= x
\begin{bmatrix}
x_1 \\
x_2
\end{bmatrix}
+
y
\begin{bmatrix}
y_1 \\
y_2
\end{bmatrix}
= \begin{bmatrix}
x_1x + y_1y \\
x_2x + y_2y
\end{bmatrix}
$$


可以看成是两个维度上的向量，矩阵里每一行的向量“指导”目标向量在对应维度上的变化。分别由一次点积得到。
$$
\begin{bmatrix}
x_1 & y_1
\end{bmatrix}
\begin{bmatrix}
x \\
y
\end{bmatrix}
= x_1x + y_1y
$$

$$
\begin{bmatrix}
x_2 & y_2
\end{bmatrix}
\begin{bmatrix}
x \\
y
\end{bmatrix}
= x_2x + y_2y
$$

若矩阵只有一行，则无第二个分量\(y\)如何变化的指示。因此就降维了，二维的向量变成了标量。
这代表目标向量在“指导”向量上的投影。

$$
\begin{bmatrix}
x_1 & y_1
\end{bmatrix}
\begin{bmatrix}
x \\
y
\end{bmatrix}
= x_1x + y_1y
$$

## 矩阵乘矩阵
那么两个矩阵相乘呢？目标向量变成了矩阵，代表着结果上起点的维度也升高了-至少是二维。计算上由两次点积变成了四次。因为后面两次负责的是第二个维度。

$$
\begin{bmatrix}
x_1 & y_1 \\
x_2 & y_2
\end{bmatrix}
\begin{bmatrix}
a & c \\
b & d 
\end{bmatrix}
= \begin{bmatrix}
\begin{bmatrix}
x_1 & y_1 \\
x_2 & y_2
\end{bmatrix}
\begin{bmatrix}
a \\
b
\end{bmatrix}
\begin{bmatrix}
x_1 & y_1 \\
x_2 & y_2
\end{bmatrix}
\begin{bmatrix}
c \\
d
\end{bmatrix}
\end{bmatrix}
= \begin{bmatrix}
x_1a + y_1b & x_1c + y_1d \\
x_2a + y2_b & x_2c + y_2d
\end{bmatrix}
$$
