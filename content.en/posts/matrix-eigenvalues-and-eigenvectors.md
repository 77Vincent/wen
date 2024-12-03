---
title: Understanding Matrix Eigenvalues and Eigenvectors
date: 2024-11-25T02:01:58+05:30
tags: [ matrix, math ]
categories: study
canonicalUrl: https://wenstudy.com/en/posts/matrix-eigenvalues-and-eigenvactors.md/
math: true
---

The fundamental property of a matrix is the transformation space, so _eigenvalues describe the scaling strength in a certain direction. Eigenvectors describe that direction._

<!-- more -->
![image of matrix-eigen](/images/matrix-eigenvalues-and-eigenvectors/matrix-eigen.png "matrix-eigen")

Pre-requisite: [Understanding Matrix Determinant](/en/posts/matrix-determinant/)

## Meaning
An eigenvector is a non-zero vector \(v\) whose direction remains unchanged after being transformed by matrix \(A\), only elongated or shortened, and the scaling factor \(\lambda\) is the eigenvalue.

$$
A \cdot v = \lambda \cdot v
$$

## Square Matrix!
The matrix must be a square matrix, i.e. \(n \times n\), and the eigenvector is also an \(n\)-dimensional vector. **Because only when the input and output dimensions are the same can we discuss the problem of a vector not changing direction after transformation.**

## Eigenvalues
According to the definition, the following derivation is obtained:
$$
A \cdot v - \lambda \cdot v = (A - \lambda \cdot I) \cdot v = 0
$$

Obtain a new homogeneous matrix \(M = A - \lambda \cdot I\) multiplied by the vector \(v\) to obtain the zero vector.

$$
M \cdot v = 0
$$

Because \(v\) is a non-zero vector, the matrix \(M\) must be a singular matrix, **because only singular matrices can transform a non-zero vector into a zero vector**. Then the determinant of \(M\) is zero, so the characteristic equation is obtained:

$$
\det(A - \lambda \cdot I) = 0
$$

> Conversely, a non-singular matrix in \(n\) dimensions can always transform any non-zero vector in an \(n\)-dimensional space into another non-zero vector.

\(A - \lambda \cdot I\) represents subtracting the eigenvalue \(\lambda\) from the diagonal elements of matrix \(A\). For example, for a \(3 \times 3\) matrix:

$$
\lambda \cdot I = \begin{bmatrix}
\lambda & 0 & 0 \\
0 & \lambda & 0 \\
0 & 0 & \lambda
\end{bmatrix}
$$

For example, for a two-dimensional matrix \(\begin{bmatrix} 4 & 1 \\ 2 & 3 \end{bmatrix}\), if the determinant is zero, then:

$$
\begin{align}
\det(\begin{bmatrix} 4 - \lambda & 1 \\ 2 & 3 - \lambda \end{bmatrix}) = (4 - \lambda)(3 - \lambda) - 2 \\
= \lambda^2 - 7\lambda + 10 = 0
\end{align}
$$

Thus, two eigenvalues are obtained: \(\lambda_1 = 2, \lambda_2 = 5\).

## Eigenvectors
After obtaining the eigenvalues, substitute them into the characteristic equation to obtain the eigenvectors.

$$
(A - \lambda \cdot I) \cdot v = 0
$$

First, for \(\lambda_1 = 2\), substitute to get:

$$
\begin{align}
A - 2 \cdot I = \begin{bmatrix} 4 - 2 & 1 \\ 2 & 3 - 2 \end{bmatrix} = \begin{bmatrix} 2 & 1 \\ 2 & 1 \end{bmatrix}
\end{align}
$$

Then convert it to an augmented matrix:

$$
\begin{bmatrix} 2 & 1 \\ 2 & 1 \end{bmatrix} \begin{bmatrix} x \\ y \end{bmatrix} = \begin{bmatrix} 0 \\ 0 \end{bmatrix}
$$

Which means:

$$
\begin{align}
2x + y = 0 \\
2x + y = 0
\end{align}
$$

In this direction, there are countless vectors to choose from, but generally, we choose the unit vector (i.e., \(x = 1\)), which gives \(y = -2\). So the eigenvector \(v_1\) is:

$$
\begin{bmatrix} 1 \\ -2 \end{bmatrix}
$$

Similarly, substitute \(\lambda_2 = 5\) to get:

$$
\begin{align}
A - 5 \cdot I = \begin{bmatrix} 4 - 5 & 1 \\ 2 & 3 - 5 \end{bmatrix} = \begin{bmatrix} -1 & 1 \\ 2 & -2 \end{bmatrix}
\end{align}
$$

Finally, the eigenvector \(v_2\) is obtained as: \(\begin{bmatrix} 1 \\ 1 \end{bmatrix}\)

The result shows that the "transformation ability" of matrix \(A\) is: stretching \(2\) times in the direction of \(v_1 = \begin{bmatrix} 1 \\ -2 \end{bmatrix}\) and stretching \(5\) times in the direction of \(v_2 = \begin{bmatrix} 1 \\ 1 \end{bmatrix}\).

## Why are there multiple values?

The reason why there are multiple eigenvalues and eigenvectors is that the matrix \(A\) is a two-dimensional matrix, which must have two directions of stretching. For an \(n\)-dimensional matrix, there are \(n\) eigenvalues and eigenvectors.
