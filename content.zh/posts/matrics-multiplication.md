---
title: 理解矩阵乘法
date: 2024-03-18T02:01:58+05:30
tags: [ matrix, math ]
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

一切要从点积开始。点积是两个向量之间的运算，结果是一个标量。理解为两个向量“合并”后可以走多远。单纯的模长乘积是不对的，因为向量带方向。为了共线，就要乘以夹角的余弦值。
``
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

_意义是矩阵（左）变换向量（右）。_

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


> \(n \times p\) * \(p \times 1\) = \(n \times 1\)

![image of matrix-vector-multiplication](/images/matrix-multiplication/m-v.png "matrix vector multiplication")

向量里 \(n\) 行代表 \(n\) 维。矩阵里，列代表输入向量的维度，行代表输出向量的维度。所以矩阵列数必须等于向量的行数；而输出维度
\(m\) 可以随意。

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

### 降维

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

### 升维

当然也可以升维，虽然向量不提供第三维度的信息，但第三维度的新信息来自前两个维度加权后的合。

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

_行向量（左）压缩矩阵（右）_

> \(1 \times p\) * \(p \times m\) = \(1 \times m\)

![image of vector-matrix-multiplication](/images/matrix-multiplication/v-m.png "vector matrix multiplication")

行向量就等于只有一行的矩阵。理解右乘后，左乘就等于把右边的矩阵压缩到一维的输出。

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

_意义是矩阵（左）变换矩阵（右）_

> \(n \times p\) * \(p \times m\) = \(n \times m\)

![image of vector-matrix-multiplication](/images/matrix-multiplication/m-m.png "matrix matrix multiplication")

就等于是线性组合，从右向左读，最右边的矩阵是初始状态，被左边的矩阵变换，可以无限串联下去。
计算方法可以把右侧的矩阵看做是诸多列向量，分别被左边的矩阵变换，再合并。

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

可以发现，最左边矩阵的输入维度是永远不会改变的。因为矩阵就是一个变换函数，函数的输入当然是固定的，而输出可以变。所以矩阵的连续相乘等同于复合函数。

> \(x\) 是任意合法（维度相符）的输入

$$
A = \begin{bmatrix} a_1 & a_2 \\ a_3 & a_4 \end{bmatrix} = f_A(x)
$$

$$
B = \begin{bmatrix} b_1 & b_2 \\ b_3 & b_4 \end{bmatrix} = f_B(x)
$$

$$
AB = f_A(f_B(x))
$$
