---
title: 理解汉明码
date: 2024-03-18T02:01:58+05:30
description: 学习汉明距离和汉明码，通过直观的，创新性的图示的方法，理解错误检测和纠正公式的正确性和原因。
tags: [computer-science, math]
categories: study 
canonicalUrl: https://wenstudy.com/posts/hamming-distance/
---

汉明距离是两个等长字符串之间对应位置上不同字符的个数。在编码理论中，用来衡量两个码字之间的差异。 `001` 和 `101` 的汉明距离是1。 对于一个码字间最短汉明距离为 m 的编码，它可以检测出 r 比特个错误只要满足以下关系:

> 1 + r ≤ m

或者可以纠正r比特个错误只要满足

> 1 + 2r ≤ m

为什么呢？

## 错误检测

假设已定义两个码字A和B，然后我们接收到C，它本应是A和B中的一个。因为C距离A或B的最大距离只能是 r（即最大为r比特的错误），而根据上述不等式，A和B的最小距离都要比 r 多1，所以错误为r比特的C一定不会被误认为任何一个合法码字。然而我们无法得知C原本应是A还是B，因为它可能来自于从A出发的1比特错误，或者从B出发的r比特错误。

![image of error detecting](/images/hamming-distance/error-detecting.png "error detecting")

This is actually a Hamming(3, 1) error correcting code [^1]

[^1]: [Hamming code (3,1)](https://en.wikipedia.org/wiki/Hamming_code)

## 错误纠正

为了将C还原为A或B，即纠正错误，我们必须知道C距离哪个合法码字更近。即然A和B可以同时出现 r 比特的错误，只需在此距离上加1，就必然会出现赢家，即只会距离A或B中的一个更近，而没有平局。若不添加这1比特的距离，C既可能来自于A的r比特错误，也可能来自于B，因而无法定论。

![image of error correction](/images/hamming-distance/error-correction.png "error correction")

## 汉明码

真实世界里，一个比特的错误很少，两个的更是罕见，因而著名的汉明码针对的是一个比特错误的修复。我们从最基本的情况来学习。

假设只传输一个比特，错误便只有一种情况，即那唯一比特的翻转。而错误检测和纠正也成为一回事。此时我们需要几个校验位？

先添加一个校验位，因校验位也可能出错，所以我们无法判断是校验位还是原信息位出错。那么我们要再添加一个校验位，两个校验位工作方式相同，只与信息位耦合，形成奇数校验（或偶数校验）。当错误发生时，有以下三种情况：

1. 信息位出错
2. 校验位1号出错
3. 校验位2号出错

对这三种情况，我们有明确的答案：

1. 信息位出错，校验位值会相同。
2. 校验位1号出错，校验位2号会和信息位的奇偶性匹配。
3. 同理，校验位2号出错，1号会和信息位匹配。

![image of 1 bit error detection](/images/hamming-distance/1-bit-error.png "1 bit error detection")

## 为了修复一比特错误而需要的校验位数量

更一般的说，信息论角度说，为了**检测** m 比特信息中，一比特错误而所需的最少校验位数量 r。

> 1 + r + m ≤ 2^r

即然一个码字的总宽度是 r + m 而每一位都可能出错（当然错误总数只能是1），于是总共有 r + m 种错误状态。另外，还有一种"没有错误"的状态。所以有 1 + r + m 种状态要区分。显而易见，r 位校验位可以表达 2^r 种状态，因而得到了上述公式。
