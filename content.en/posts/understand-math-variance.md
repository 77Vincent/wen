---
title: Understand Variance in Math
date: 2023-11-30T02:01:58+05:30
tags: [ probability, statistics, math ]
categories: study
canonicalUrl: https://wenstudy.com/posts/en/understand-math-variance/
math: true
---

The variance measures the dispersion of a random variable, which is the expectation of the squared difference between each data point and the expectation.
<!--more-->

$$
Var(X) = E[(X - E[X])^2]
$$

The variance \(Var(X)\) is also written as \(\sigma^2\), where \(\sigma\) is the standard deviation.

## Equivalent Formula
_Variance is the expectation of the square of each data point minus the square of the expectation._

$$
Var(X) = E[X^2] - E[X]^2
$$

Why? Because of the standard definition:

$$
E[(X - E[X])^2] = E[X^2 - 2XE[X] + E[X]^2]
$$

According to the linearity property of expectation (i.e., each part can be calculated separately):

$$
E[X^2 - 2XE[X] + E[X]^2] = E[X^2] - 2E[X]E[E[X]] + E[X]^2
$$

Among them, \(E[X]\) is a constant, so \(E[E[X]] = E[X]\). Finally, we get:

$$
E[X^2] - 2E[X]E[X] + E[X]^2 = E[X^2] - E[X]^2
$$

> This equivalent formula is more convenient when calculating the variance of some distributions.

## Multiplication

For any constant \(a\) and random variable \(X\), the variance changes with the square of the constant, _because the variance itself is squared_.

$$
Var(aX) = a^2 Var(X)
$$

How to prove it? First, because the expectation calculation is linear:

$$
E[aX] = aE[X]
$$

So:

$$
\begin{align}
Var(aX) &= E[(aX - E[aX])^2] \\
&= E[(aX - aE(X))^2] \\
&= E[a^2(X - E(X)^2] \\
&= a^2E[(X - E(X))^2] \\
&= a^2Var(X)
\end{align}
$$

## Addition

For two random variables \(X\) and \(Y\), we have:

$$
Var(X + Y) = Var(X) + Var(Y) + 2Cov(X, Y)
$$

Where \(Cov(X, Y)\) is the covariance of \(X\) and \(Y\), indicating the linear relationship between \(X\) and \(Y\).

Why? First, expand according to the definition of variance:

$$
Var(X + Y) = E[(X + Y - E[X + Y])^2]
$$

According to the linearity property of expectation:

$$
E[X + Y] = E[X] + E[Y]
$$

So:

$$
\begin{align}
Var(X + Y) &= E[(X + Y - E[X] - E[Y])^2] \\
&= E[((X - E[X]) + (Y - E[Y]))^2]
\end{align}
$$

Expand the square term:

$$
Var(X + Y) = E[(X - E[X])^2 + 2(X - E[X])(Y - E[Y]) + (Y - E[Y])^2]
$$

Still, the expectation is linear, so it can be calculated separately:

$$
Var(X + Y) = E[(X - E[X])^2] + 2E[(X - E[X])(Y - E[Y])] + E[(Y - E[Y])^2]
$$

Among them:
1. \(E[(X - E[X])^2] = Var(X)\)
2. \(E[(Y - E[Y])^2] = Var(Y)\)
3. \(E[(X - E[X])(Y - E[Y])] = Cov(X, Y)\) this is the covariance of \(X\) and \(Y\).

So we get:

$$
Var(X + Y) = Var(X) + Var(Y) + 2Cov(X, Y)
$$

If \(X\) and \(Y\) are independent, then their covariance, i.e., \(Cov(X, Y)\) is 0, then:

$$
Var(X + Y) = Var(X) + Var(Y)
$$

If it is two identical random variables \(X\)? Equivalent to the multiplication case, we have:

$$
Var(X + X) = Var(2X) = 4Var(X)
$$

> That is, the same random variable is superimposed, and the discrete degree will increase.
