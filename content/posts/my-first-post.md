---
title: "轻松理解汉明码"
date: 2018-03-18T02:01:58+05:30
description: "hhh"
tags: [Primer, todo]
featured_image: "/images/notebook.jpg"
categories: Todo
comment : false
---

对于一个码字间最短汉明距离为 m 的编码，它可以检测出 r 比特个错误只要满足

> 1 + r ≤ m

或者可以纠正r比特个错误只要满足

> 1 + 2r ≤ m

为什么呢？

# 错误检测

假设有两个码字A和B，然后我们接收到C，它本应是A和B中的一个。因为C距离A或B的最大距离只能是 r（即最大为r比特的错误），而根据上述不等式，A和B的最小距离也比 r 多1，所以错误为r比特的C一定不会被误认为任何一个合法码字。然而我们无法得知C原本应是A还是B因为它可能来自于从A出发的1比特错误，或者从B出发的r比特错误。

# 错误纠正

为了将C还原为A或B，即纠正错误，我们必须要知道C距离哪个合法码字更近。即然A和B可以同时出现 r 比特的错误，只需在此距离上加1，就必然会出现赢家，即任意一个错误码字C，只会距离A或B中的一个更近，没有其他情况。若不添加这1比特的距离，C既可能来自于A的r比特错误，也可能同样来自于B，因而无法定论。

# 汉明码

真实世界里，一个比特的错误其实很少，而两个的更是罕见，因而著名的汉明码针对的是一个比特错误的修复。我们也用同样的目标来学习，并且用一个极简的例子。
假设我们只传输一个比特的信息，错误便只有一种情况，即那个比特的翻转。而错误检测和纠正也成为一回事。此时我们需要几个校验位方可实现错误纠正？
假设添加一个校验位，因校验位本身也可能出错，所以在只针对一个错误的情况下，我们无法判断是校验位还是原本信息位出错。那么我们便需要再添加一个校验位，两个校验位工作方式相同，各自只与信息位耦合，形成奇数校验或偶数校验。于是当错误发生时，有以下三种情况：

1. 信息位出错
2. 校验位一号出错
3. 校验位二号出错

对这三种情况，我们有明确的答案：

1. 信息位出错，其他两个校验位值会相同。
2. 校验位一号出错，校验位二号会和信息位的奇偶性匹配。
3. 同理，校验位二号出错，一号会和信息位匹配。

# 为了修复一比特错误而需要的校验位数量

更一般的说，我们可以用一个公式来决定，为了检测 m 比特信息中，一比特错误而所需的最少校验位数量 r。为什么呢？

> 1 + r + m ≤ 2^r

即然这个码字的总宽度是 r + m 而每一位都可能出错（当然一次只能有一个），于是总共便有 r + m 种错误情况。另外，还有一种"没有错误"的情况。所以总共有 1 + r + m 种情况要区分，这便是 r 位校验位所需要标记的全部状态。显而易见，r 位校验位可以表达 2^r 种状态，因而得到了上述公式。