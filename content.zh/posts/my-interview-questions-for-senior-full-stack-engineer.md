---
title: 软件工程师的通用面试问题
date: 2025-01-01T02:01:58+05:30
tags: [ computer-science, interview, software-engineering ]
categories: study
math: true
canonicalUrl: https://wenstudy.com/posts/my-interview-questions-for-senior-full-stack-engineer/
---

由于工作中经常需要为团队招聘进行面试，以下总结了一份问题集，针对中高级全栈工程师，从基础通用知识，前端、后端到架构。以经典为主，主打永不过时。

<!--more-->

## 计算机网络

### CIDR

以下 `CIDR` 地址块有多少个可用的IP地址？

```
10.0.0.0/24
```

> 2^(32-24)-2=2^8-2=254

其对应的 IP 地址范围是什么？

> 10.0.0.1 ~ 10.0.0.254
> 
> 10.0.0.0 不可分配，是子网的起始地址，用于标识网络
> 
> 10.0.0.255 不可分配，是子网的广播地址，用于向子网内所有设备发送消息

### 子网掩码

有以下子网掩码和两个IP，它们是否在同一个子网？
```
255.255.255.0
192.168.0.10
192.168.1.10
```

> 不在，因为两个IP地址与子网掩码的AND结果不同。分别是
> 
> 192.168.0.10 & 255.255.255.0 = 192.168.0.0
> 
> 192.168.1.10 & 255.255.255.0 = 192.168.1.0

### 什么是默认网关？

默认网关是一个网络设备，用于将数据包从一个网络发送到另一个网络。当设备要发送数据包到另一个网络时，它首先检查目标IP地址是否在同一子网内。如果不在，它会将数据包发送到默认网关，由默认网关转发到目标网络。

同一个字网内的设备不需要走网关，因为它们可以直接通信。

> 默认网关通常是路由器，它连接了两个不同的网络。

### 什么是NAT？

NAT(Network Address Translation)是一种将私有IP地址转换为公共IP地址的技术。在私有子网中，每个设备只有一个私有IP地址，无法直接与Internet通信。NAT通过将私有IP地址映射到公共IP地址。

### UDP VS TCP

简单说，TCP是面向连接的，可靠的，有序的，基于字节流的传输协议。UDP是无连接的，不可靠的，无序的，基于数据包的传输协议。

| Feature                     | TCP                                                          | UDP                                                                             |
|-----------------------------|--------------------------------------------------------------|---------------------------------------------------------------------------------|
| Connection                  | Connection-oriented                                          | Connectionless                                                                  |
| Reliability                 | Reliable due to error checking, retransmission, reassembling | Unreliable due to no error checking and retransmission, reassembling            |
| Speed                       | Slower due to overhead                                       | Faster due to no overhead                                                       |
| Flow and Congestion control | Yes                                                          | No                                                                              |
| Header size                 | Larger (20 - 60 bytes)                                       | Smaller (8 bytes)                                                               |
| Use case                    | Web browsing, file transfer, email                           | Real-time application like video streaming, online gaming, VoIP (voice over IP) |

## 数据库

### 解释ACID

ACID是数据库事务的四个特性，保证数据库事务在并发和故障情况下的可靠性和一致性。

| Term        | Explanation                    | In one word |
|-------------|--------------------------------|-------------|
| Atomicity   | 事务中的所有部分只能全部完成或全部失败            | 有或没有        |
| Consistency | 事务开始和结束时，数据库的完整性约束没有被破坏 ｜ 一致正确 |
| Isolation   | 事务之间是相互隔离的，互不干扰                | 并发安全 ｜      |
| Durability  | 事务一旦提交，对数据库的改变是永久的             | 持久性         |

### 什么是隔离级别？

隔离级别是数据库事务并发的一个概念，定义了事务之间的隔离程度。SQL标准定义了四个隔离级别：

1. 读未提交（Read Uncommitted）
2. 读已提交（Read Committed）
3. 可重复读（Repeatable Read）
4. 序列化（Serializable）

### 并发异常

这三个现象是数据库事务并发的问题，是由于事务隔离级别不同导致的。

| 现象                       | 描述                                 | 所需隔离等级 |
|--------------------------|------------------------------------|--------|
| 脏独（Dirty Read）           | 读取到未提交的数据                          | 读已提交   |
| 不可重复度（Unrepeatable Read） | 一个事务中，两次读取之间，另一个事务修改了数据，导致结果不一致    | 可重复读   |
| 幻读（Phantom Read）         | 一个事务中，两次读取之间，另一个事务插入/删除了数据，导致行数不一致 | 序列化    |

## 开发

### Heap 数据结构


### 什么是竞态条件？

是多个线程或进程并发执行，修改同一个变量时，由于执行顺序不确定，可能导致：

1. 数据不一致，每次执行可能得到不同的结果。
2. 死锁，或数据处于非法状态。
3. 难以复现和调试。

> 重点是执行顺序不确定，而非总是错误，比如可能有时候按 1-2-3 顺序正确执行，有时候混乱。

### 设计模式

设计模式是解决软件设计中常见问题的经验总结，常用的有：
1. 工厂
2. 单例
3. 策略
4. 适配器
5. 模版

### 单例模式

单例模式是一种创建型设计模式，确保一个类只有一个实例，并提供一个全局访问点。以下是Go语言的实现。 首先是经典实现：

```go
package main

import "sync"

type Singleton struct{}

var (
	s  *Singleton
	mu sync.Mutex
)

func Get() *Singleton {
	if s == nil {
		mu.Lock()
		defer mu.Unlock()

		if s == nil {
			s = &Singleton{}
			return s
		}
		return s
	}
	return s
}
```

> 重点：首次检查为空加锁后，需再次检查是否为空，而不能直接创建，否则有可能在加锁完成前，其他线程已经创建了实例而导致重复创建。

比较现代的实现方式是使用 `sync.Once`，确保线程安全。省去了更多累赘的 `if` 判断。

```go
package main

import "sync"

type Singleton struct{}

var (
	s    *Singleton
	once sync.Once
)

func Get() *Singleton {
	once.Do(func() {
		s = &Singleton{}
	})
	return s
}
```

## 云计算

### 可用区（AZ）是什么？

可用区（Availability Zone）是一个独立的数据中心，通常位于同一地理区域内，但是物理上相互隔离。可用区之间有独立的电力、网络和冷却系统，确保一个可用区的故障不会影响其他可用区。

## 架构

### 如何评价一个系统？

四字诀：安全可靠，无廉价美。

![image of system-requirement pyramid](/images/aws-well-architected-framework/aws-well-architected-framework.png "system-requirement-pyramid")

### 为什么用 CDN？

1. 

### DDoS 攻击

什么是 DDoS 攻击？

> DDoS（Distributed Denial of Service）是一种网络攻击，通过分布式客户端，产生大量请求，使得目标服务过载或由此触发错误，导致服务不可用。

如何防范？

> 设置防火墙，限制访问频率，使用CDN，使用DDoS防护服务，使用IP黑名单，使用CAPTCHA验证。

### 如何实现高可用

1. 多地区和多可用区部署。
2. 应用服务器前部署负载均衡。
3. workload自动缩扩容
4. 防灾冗余，包括使用 passive-active，active-active 部署。
5. 充分的监控和报警

### 保障数据库安全

1. 将数据库放在私有子网中，只允许特定IP访问。
2. 分离数据库和应用服务器的网络。
3. 使用强密码，定期更换，最小曝光。
4. 在传输和存储时加密数据。

## 数学

主要考察基础数学知识与临场解决问题的能力，优秀的工程师数学能力必不可少！

### 排列组合

一个房间里99个人，两两握手一次，不重复，共有多少次握手？

$$
C_n^m = \frac{n!}{m!(n-m)!}
$$

$$
C_{99}^2 = \frac{99!}{2!(99-2)!} = \frac{99*98}{2} = 99*49 = 4851
$$

> 99个人中，第一个人有98个人可以握手，第二个人有97个人可以握手，以此类推。本质上是从99个人中选取2个人的排列组合。

### 队列

一小时内平均有100个顾客随机到达，符合泊松分布，平均每个顾客服务时间是100毫秒，一次只能服务一个顾客，无队列，求一小时平均有多少顾客会被拒绝。

$$
\lambda （到达率） = \frac{1}{\text{到达时间间隔}} = \frac{100}{3600} = \frac{1}{36}
$$

$$
\mu （服务率） = \frac{1}{\text{服务时间}} = \frac{1000}{100 (毫秒)} = 10
$$

$$
\rho （利用率） = \frac{\lambda}{\mu} = \frac{1}{36*10} = \frac{1}{360}
$$

$$
\text{一小时内平均拒绝顾客数} = \lambda * \rho = \frac{1}{360} * 100 = \frac{100}{360} = \frac{5}{18} = 0.2777
$$

> M/M/1队列是一个基本的排队模型，M表示到达率和服务率都是指数分布，1表示只有一个服务台。
> 
> M/M/1 队列模型中，拒绝率等于利用率，因为利用时即为拒绝时。

### 离散数学

通用排序算法，即对数据没有任何假设的情况下，时间复杂度的下限是多少？

$$
\log_2(n!) = \Theta(n \log n)
$$

为什么？排序本质是比较，可以用一个决策树表示，决策树的深度即为最坏情况下的比较次数。对于 \(n\) 个元素的排序，有 \(n!\) 种排列，所以决策树的叶子节点数为 \(n!\)，高度为 \(h\)，满足：

$$
2^h \geq n!
$$

$$
h \geq \log_2(n!)
$$

基于斯特林公式，有 

$$
\begin{aligned}
n! &\approx \sqrt{2\pi n}(\frac{n}{e})^n \\
\log_2(n!) &\approx n \log_2 n - n \log_2 e \\
&\approx n \log_2 n
\end{aligned}
$$

所以：

$$
h \geq n \log_2 n
$$

> 简单理解：首先必须要遍历所有元素，因此时间复杂度至少是 \(O(n)\)；每看到一个元素，都要用二分法去寻找它在长度为 \(n\) 的数组中的位置，这里的复杂度是 \(O(\log n)\)。所以总的时间复杂度的下限是 \(O(n \log n)\)。

### 统计

P50 和 P99 是什么？

> P50 是中位数，即数据中位于中间位置的值。P99 是百分位数，即数据中位于前99%的值。

## 开放问题

以下问题无标准答案，考察面试者对其工作领域的了解程度和思考深度。

1. 你最常用的编程语言的优缺点是什么？
2. 上一个负责的项目有多少行代码？多少个API？多少个批任务？多少用户？最大RPS？
3. 如何说服他人，兜售自己的意见？
4. 十年内的目标？
5. 未来工程师的发展方向？
6. 犯过的最大的错是什么
