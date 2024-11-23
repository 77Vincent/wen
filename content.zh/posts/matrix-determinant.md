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

<!--more-->
> **矩阵的秩:** 指矩阵中的线性无关行或列的最大数目。表示矩阵在空间中所能描述的独立方向的数目，即矩阵实际能表达的维度上限。

本篇前提：[理解矩阵乘法](https://wenstudy.com/posts/matrix-multiplication/)

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

## 计算
### 二维矩阵
先从这个 \(2*2\) 矩阵开始。按行或按列书写向量不重要，结果一样，比如用列向量：

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

为什么这样计算？先是直观的解释：因为结果其实是下图平行四边形的面积。所以不论用行还是列向量书写，面积是不会变的。
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

但如果交换两列，行列式的值会变号。

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

### 右手定则
你可能还是会问，不对，我看上面的图两个向量交换了位置，面积不还是一样的吗？这里必须要提的是，_坐标空间是有一个绝对的方向的，可以通过右手定则来确定_。它没有为什么，不是右手就得是左手，总要有一个定义。右手定则的特征就是，从\(z\) 往下看的话，\(x\) 轴 在 \(y\) 轴的顺时针方向。

![image of right-hand-rule](/images/matrix-determinant/right-hand-rule.jpg "right-hand-rule")

所以上面所说的交换两列，实际上是把 \(x\) 轴和 \(y\) 轴交换了位置，在绝对的坐标空间里，相当于翻转了空间，所以有向面（体）积就会变号。

于是回过头看上述二维空间行列式的计算公式，\(a\) 是第一维上的底边，\(d\) 就是与其垂直的第二维度上的高，相乘就得到了面积；因为维度之间是平等的，所以同样的事，\(b\) 和 \(c\) 也要做一次。因此相加两个面积就可以了：
$$
ad + bc = \text{面积1} + \text{面积2}
$$

但是第二部分里的 \(b\) 其实代表的是第二维度上的底边（它是第一个向量的 \(y\)），等于说 \(bc\) 是在翻转过（左手法则）的空间里计算的，结果天然带符号。为了纠正这个翻转，计算时要主动乘上一个负号。

$$
ad + (-bc) = ad - bc
$$

### 三维矩阵
三维矩阵行列式的计算逻辑和上述针对二维的异曲同工，只是需要一个递归过程降到二维。

$$
\begin{vmatrix}
x & y & z \\
x_1 & y_1 & z_1 \\
x_2 & y_2 & z_2
\end{vmatrix} = x\begin{vmatrix}
y_1 & z_1 \\
y_2 & z_2
\end{vmatrix} - y\begin{vmatrix}
x_1 & z_1 \\
x_2 & z_2
\end{vmatrix} + z\begin{vmatrix}
x_1 & y_1 \\
x_2 & y_2
\end{vmatrix}
$$

用一个特殊的例子可以很好解释为什么上面的公式是对的。假设三个向量都在三条坐标轴上，那么 \(x\begin{vmatrix}y_1 & z_1 \\ y_2 & z_2\end{vmatrix}\) 刚好就是以 \(x\) 为高，乘以在 \(y-z\) 平面上的底面的体积，而那个底面积刚好就是对应子二维矩阵的行列式。同理，这个计算步骤在 \(y\) 和 \(z\) 都要来一次，因为大家都是平等的。

![image of 3d-determinant](/images/matrix-determinant/3d-determinant.png "3d-determinant")

这里又出现了正负号，还是交替，为何？每个轴为高往下看的时候，底面矩阵必须符合右手定则，比如 \(x\) 为高的时候，底面 \(y-z\) 平面的 \(y\) 正确地位于 \(z\) 的顺时针方向，所以符合右手定则，无需乘负号。\(z\) 也符合。但 \(y\) 为高的时候，底面 \(x-z\) 平面的 \(x\) 位于 \(z\) 的逆时针方向，不符合右手定则，等于是在一个翻转的空间里，所以要乘负号纠正。

### 拉普拉斯展开
这个递归的过程可以一直延伸到 \(n\) 维，这就是拉普拉斯展开（Laplace Expansion）。这种展开过程会涉及到矩阵中的每一个元素与它的代数余子式（algebraic minor）相乘，其中代数余子式前面会有一个符号因子，符号因子就是所谓的正负号交替。

#### 余子式（Minor）
于矩阵中的某个元素 \(a_{ij}\) 而言，它的余子式 \(M_{ij}\) 是去掉第 \(i\) 行和第 \(j\) 列的子矩阵的行列式。

#### 代数余子式（Algebraic Minor）

代数余子式是余子式再乘以对应的符号因子。对于矩阵元素 \(a_{ij}\) 的代数余子式 \(C_{ij}\) 的计算公式如下：

$$
C_{ij} = (-1)^{i+j}M_{ij}
$$

其中 \(M_{ij}\) 就是 \(a_{ij}\) 的余子式，代表了去掉第 \(i\) 行和第 \(j\) 列的子矩阵的行列式。

这样，对于 \(3 * 3\) 通用的拉普拉斯展开公式（也叫代数余子式展开）就是：
$$
\det(A) = a_{11}C_{11} + a_{12}C_{12} + a_{13}C_{13}
$$

## 特性
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

### 行列式为零

行列式为零是一个重要的特性，代表矩阵无法张成 \(n\) 维空间，最多 \(n-1\) 维，所以在 \(n\) 维上体积为零，比如一个没有高的面或体，不管底边（面）有多大都白搭。其根本原因是存在线性相关的向量。即有一些向量是多余的，没有引入关于那个维度的独立的信息。

存在相同行（列）的行列式为零。因为这意味着存在至少一组线性相关的向量。
$$
\begin{vmatrix}
a & a \\
b & b
\end{vmatrix} = a*b - a*b = 0
$$

或一个向量是另一个倍数。
$$
\begin{vmatrix}
a & ka \\
b & kb
\end{vmatrix} = k\begin{vmatrix}
a & a \\
b & b
\end{vmatrix} = k*0 = 0
$$

或存在零向量。这其实是上面 \(k=0\) 的特例。（少一个维度，当然体积为零）

$$
\begin{vmatrix}
a & 0 \\
b & 0
\end{vmatrix} = a*0 - 0*b = 0
$$

或者一个向量是另几个向量的线性组合，导致的结果依旧是存在非独立向量（多余的）。这个等式用到了上述的加法和数乘性质。

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

### 行列式不变
把其中一行（列）的倍数加到另一行（列）上，行列式不变。这个结论由上面的结论推导而来。

$$
\begin{vmatrix}
a & c + ka \\
b & d + kb
\end{vmatrix} = \begin{vmatrix}
a & c \\
b & d
\end{vmatrix} + \begin{vmatrix}
a & ka \\
b & kb
\end{vmatrix} = \begin{vmatrix}
a & c \\
b & d
\end{vmatrix} + 0 = \begin{vmatrix}
a & c \\
b & d
\end{vmatrix}
$$

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
