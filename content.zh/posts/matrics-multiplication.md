---
title: 理解矩阵乘法
date: 2024-03-18T02:01:58+05:30
tags: [matrix, math]
categories: study 
canonicalUrl: https://wenstudy.com/posts/matrics-multiplication/
math: true
---

矩阵乘法分为四种情况，本质是一样的，都是向量之间的运算。
1. 向量乘向量（点积） 
2. 矩阵乘向量（右乘）
3. 向量乘矩阵（左乘）
4. 矩阵乘矩阵
<!--more-->

## 向量乘向量（点积） 
点击是两个向量之间的运算，结果是一个标量。理解为两个向量“合并”后可以走多远。单纯的模长乘积是不对的，因为向量带方向。为了共线，就要乘以夹角的余弦值。

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

![image of dot product](/images/dot-product.png "dot product")

更广泛的：
$$
a \cdot b = \Sigma_{i=1}^{n} a_ib_i = a_1b_1 + a_2b_2 + \cdots + a_nb_n
$$

## 矩阵乘向量（右乘）
_意义是矩阵变换向量。_

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


> \(m \times n\) * \(n \times 1\) = \(m \times 1\)

![image of matrix-vector-multiplication](/images/matrix-vector-multiplication.png "dot product")

向量里 \(n\) 行代表 \(n\) 维。矩阵里，列代表输入向量的维度，行代表输出向量的维度。所以矩阵列数必须等于向量的行数；而输出维度 \(m\) 可以随意。


$$
x_{\text{new}} = \\
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
y_{\text{new}} = \\
\begin{bmatrix}
x_2 & y_2
\end{bmatrix}
\begin{bmatrix}
x \\
y
\end{bmatrix}
= x_2x + y_2y
$$

这里表示把输入的二维向量变换成二维空间里的另一个向量。 \([x_1, y_1]\) 告诉新分量 \(x_{\text{new}}\) 如何因两个维度而变化。\([x_2, y_2]\) 告诉新分量 \(y_{\text{new}}\) 如何因两个维度而变化。更高维度同理。计算方法是点积。

$$
Ax =
\begin{bmatrix}
x_{\text{new}} \\
y_{\text{new}}
\end{bmatrix}
$$

若矩阵只有一行，则输出空间是一维，无第二个分量 \(y\) 的变化指示。输入的二维向量变成了标量，空间由平面变成了线（只剩了长度）。

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

另一种降维是，输出空间依旧是二维，但第二个维度分量为 \(0\)，则输入的二维向量被投影到二维空间中的一条线 - \(x\) 轴上。

$$
\begin{bmatrix}
x_1 & y_1 \\
0 & 0
\end{bmatrix}
\begin{bmatrix}
x \\
y
\end{bmatrix}
= \begin{bmatrix}
x_1x + y_1y \\
0
\end{bmatrix}
$$

当然也可以升维，虽然向量不提供第三维度的信息，但第三维度的数值可以来自于前两个维度加权后的合。

$$
\begin{bmatrix}
x_1 & y_1 \\
x_2 & y_2 \\
x_3 & y_3
\end{bmatrix}
\begin{bmatrix}
x \\
y \\
\end{bmatrix}
= \begin{bmatrix}
x_1x + y_1y \\
x_2x + y_2y \\
x_3x + y_3y
\end{bmatrix}
$$

## 向量乘矩阵（左乘）
_行向量（一行矩阵）压缩右边的矩阵_

> \(1 \times n\) * \(n \times m\) = \(1 \times 1\)

![image of vector-matrix-multiplication](/images/vector-matrix-multiplication.png "dot product")

行向量理解为一行的矩阵，等于把右边的矩阵压缩到一维的输出维度。

$$
\begin{align}
\begin{bmatrix}
x & y
\end{bmatrix}
\begin{bmatrix}
x_1 & x_2 \\
y_1 & y_2
\end{bmatrix}
&= \begin{bmatrix}
\begin{bmatrix}
x & y
\end{bmatrix}
\begin{bmatrix}
x_1 \\
y_1
\end{bmatrix}
\begin{bmatrix}
x & y
\end{bmatrix}
\begin{bmatrix}
x_2 \\
y_2
\end{bmatrix}
\end{bmatrix} \\
&=
\begin{bmatrix}
x_1x + x_2y & y_1x + y_2y
\end{bmatrix}
\end{align}
$$

## 矩阵乘矩阵
矩阵相乘就是线性组合，从右向左读。

具体计算方法可以把右侧的矩阵看做是诸多列向量，分别被左边的矩阵变换，再合并。

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
