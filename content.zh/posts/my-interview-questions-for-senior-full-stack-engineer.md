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

#### 以下 `CIDR` 地址块有多少个可用的IP地址？

```
10.0.0.0/24
```

> 2^(32-24)-2=2^8-2=254

#### 其对应的 IP 地址范围是什么？

> 10.0.0.1 ~ 10.0.0.254
>
> 10.0.0.0 不可分配，是子网的起始地址，用于标识网络
>
> 10.0.0.255 不可分配，是子网的广播地址，用于向子网内所有设备发送消息

### 子网掩码

#### 有以下子网掩码和两个IP，它们是否在同一个子网？

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

### 默认网关？

默认网关（`Default Gateway`）是一个网络设备，用于将数据包从一个网络发送到另一个网络。当设备要发送数据包到另一个网络时，它检查目标IP地址是否在同一子网内。

* 在，不需走网关，设备可直接通信。
* 不在，将数据包发送到默认网关，由默认网关转发到目标网络。

> 默认网关通常是路由器，它连接了两个不同的网络。

### NAT

NAT（`Network Address Translation`）是一种将私有IP地址转换为公共IP地址的技术。在私有子网中，每个设备只有一个私有IP地址，无法直接与Internet通信。NAT通过将私有IP地址映射到公共IP地址。

### UDP VS TCP

简单说，TCP是面向连接的，可靠的，有序的，基于字节流的传输协议。UDP是无连接的，不可靠的，无序的，基于数据包的传输协议。

| Feature                     | TCP                                                          | UDP                                                                             |
|-----------------------------|--------------------------------------------------------------|---------------------------------------------------------------------------------|
| **Connection**              | Connection-oriented                                          | Connectionless                                                                  |
| **Reliability**             | Reliable due to error checking, retransmission, reassembling | Unreliable due to no error checking and retransmission, reassembling            |
| **Speed**                   | Slower due to overhead                                       | Faster due to no overhead                                                       |
| **Flow/Congestion control** | Yes                                                          | No                                                                              |
| **Header size**             | Larger (20 - 60 bytes)                                       | Smaller (8 bytes)                                                               |
| **Use case**                | Web browsing, file transfer, email                           | Real-time application like video streaming, online gaming, VoIP (voice over IP) |

### DNS

#### A 记录（A record）

> A 记录是将域名映射到IPv4地址的记录。其中 `@` 表示根域名。

| example.com | record type | value     | TTL   |
|-------------|-------------|-----------|-------|
| @           | A           | 192.0.2.1 | 14400 |

#### CNAME 记录

> CNAME 记录是将域名映射到另一个域名的记录。等于是域名的别名。例如：

| blog.example.com | record type | value                      | TTL   |
|------------------|-------------|----------------------------|-------|
| @                | CNAME       | is an alias of example.com | 32600 |

### HTTP

#### 为什么要三次握手？

1. 第一次证明 **客户端** 可以 **发送数据**（SYN）
2. 第二次证明 **服务端** 可以 **接收数据和发送数据**（SYN + ACK）
3. 第三次证明 **客户端** 可以 **接收数据**（ACK）

> 类似打电话，A说“可以听到吗”，B说“可以听到，你呢？”，A说“我也可以听到”。

#### 为什么要四次挥手？

1. 第一次 **客户端** 告诉 **服务端** 客户端不再发送数据（FIN）
2. 第二次 **服务端** 告诉 **客户端** 服务端还有数据要发送（ACK）
3. 第三次 **服务端** 告诉 **客户端** 服务端不再发送数据（FIN）
4. 第四次 **客户端** 告诉 **服务端** 可以关闭连接（ACK）

> TCP 是全双工的，所以关闭连接需要两次。类似于结束通话，A说“我不说了”，B说“我还有话要说”，B说“我也不说了”，A说“好的”。

#### 请求 URL 长度有限制吗？

> HTTP协议没有规定URL的最大长度，但是浏览器和服务器都有限制。例如，Chrome浏览器的URL长度限制是 2KB。

#### 如何禁止浏览器缓存？

> 可以通过设置HTTP响应头（response header）来禁止浏览器缓存，其中：

```
Cache-Control: no-cache, no-store, must-revalidate
Pragma: no-cache
Expires: 0
```

1. Cache-Control 中的 no-cache 表示不缓存，no-store 表示不存储，must-revalidate 表示必须重新验证。
2. Pragma 中的 no-cache 表示不缓存。（HTTP/1.0）
3. Expires 设置为 0 表示立即过期。（HTTP/1.0）

## 数据库

### ACID

`ACID` 是数据库事务的四个特性，保证数据库事务在并发和故障情况下的可靠性和一致性。

| Term            | Explanation             | In one word |
|-----------------|-------------------------|-------------|
| **Atomicity**   | 事务中的所有部分只能全部完成或全部失败     | 有或没有        |
| **Consistency** | 事务开始和结束时，数据库的完整性约束没有被破坏 | 一致正确        |
| **Isolation**   | 事务之间是相互隔离的，互不干扰         | 并发安全        |
| **Durability**  | 事务一旦提交，对数据库的改变是永久的      | 持久性         |

### 隔离级别

隔离级别是数据库事务并发的一个概念，定义了事务之间的隔离程度。SQL标准定义了四个隔离级别：

1. 读未提交（`Read Uncommitted`）
2. 读已提交（`Read Committed`）
3. 可重复读（`Repeatable Read`）
4. 序列化（`Serializable`）

### 并发异常

这三个现象是数据库事务并发的问题，是由于事务隔离级别不同导致的。

| 现象        | 描述                                 | 所需隔离等级 |
|-----------|------------------------------------|--------|
| **脏独**    | 读取到未提交的数据                          | 读已提交   |
| **不可重复度** | 一个事务中，两次读取之间，另一个事务修改了数据，导致结果不一致    | 可重复读   |
| **幻读**    | 一个事务中，两次读取之间，另一个事务插入/删除了数据，导致行数不一致 | 序列化    |

### 数据库安全

1. 将数据库放在私有子网中，只允许特定IP访问。
2. 分离数据库和应用服务器的网络。
3. 使用强密码，定期更换，最小曝光。
4. 在传输和存储时加密数据。

### 索引

#### 联合索引

已知以下表结构和索引，表有千万行数据。

```sql
CREATE TABLE users (
    a INT,
    b INT,
    c INT
);

CREATE INDEX idx ON users (a, b, c);
```

这个查询会用到索引吗？

```sql
SELECT * FROM users WHERE a > 1 ORDER BY c;
```

> 不会
> 
> 根据最左匹配原则，索引的第一个字段是 `a`，但是查询条件是单边范围选取 `a > 1`，不是精确匹配且范围过大，所以不会用到索引。

## 开发

### 堆（Heap）

堆是一个完全二叉树，每个节点的值大于等于其子节点的值，为最大堆，反之为最小堆。

1. 插入和删除的时间复杂度是 \(O(log n)\)
2. 查询最值的时间复杂度是 \(O(1)\)

> 凡是想要快速找到最值的场景，都可以考虑使用堆，例如优先队列、Dijkstra算法、堆排序等。

### 栈（Stack）

栈是一种先进后出（FILO）的数据结构，只能在栈顶进行插入和删除操作。

1. 插入和删除的时间复杂度是 \(O(1)\)
2. 查询的时间复杂度是 \(O(n)\)

> 凡是需要先进后出的场景，都可以考虑使用栈，例如函数调用、表达式求值、括号匹配等。

### 竞态条件

竞态条件（`Race Condition`）是多个线程或进程并发执行时，由于执行顺序不确定，导致结果不一致的情况。

1. 数据不一致，每次执行可能得到不同的结果。
2. 死锁，或数据处于非法状态。
3. 难以复现和调试。

错误示例

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	for i := 0; i < 10; i++ {
		go func() {
			result = i
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(result)
}
```

> 重点是执行顺序不确定，而非总是错误，比如可能有时候按 1-2-3 顺序正确执行，有时候混乱。
 
正确写法

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	result := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
        go func(i int) {
            result = i
            wg.Done()
        }(i)
		wg.Wait()
	}
    fmt.Println(result)
}
```

> 重点在于同步化，保证顺序

### 设计模式

设计模式是解决软件设计中常见问题的经验总结，常用的有：

1. 工厂
2. 单例
3. 策略
4. 适配器
5. 模版

下面是两种我认为工作中常用到的设计模式：`单例模式` 和 `策略模式`。

#### 单例模式

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

#### 策略模式（Strategy）

策略模式是一种行为设计模式，定义一系列算法，将每个算法封装到独立的类中，并使它们可以互相替换。以下是Go语言的实现。

```go
package main

import "fmt"

type Strategy interface {
	Do()
}

type StrategyA struct{}

func (s *StrategyA) Do() {
	fmt.Println("Strategy A")
}

type StrategyB struct{}

func (s *StrategyB) Do() {
	fmt.Println("Strategy B")
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) Execute() {
	c.strategy.Do()
}

func main() {
	context := &Context{}

	if true {
        context.SetStrategy(&StrategyA{})
        context.Execute()
    } else {
        context.SetStrategy(&StrategyB{})
        context.Execute()
    }
}
```

> 调用许发生在运行时，可以动态切换策略，而不需要修改代码。

### 算法

#### 快速幂（Exponentiation by squaring）

快速幂是一种计算 \(a^b\) 的算法，时间复杂度是 \(O(\log b)\)。非常考察对位运算和二进制的理解。以 \(a^{13}\) 为例，二进制为
`1101`。

$$
1101 = 2^3 + 2^2 + 0\cdot2^1 + 2^0
$$

所以

$$
a^{13} = a^{2^3 + 2^2 + 0\cdot2^1 + 2^0} = a^{2^3} \cdot a^{2^2} \cdot a^{0\cdot2^1} \cdot a^{2^0}
$$

从右往左看，如果二进制位是1，则结果需乘以对应的 \(a^{2^i}\)，否则跳过。

1. 第 \(i = 0\) 位是 \(1\)，结果乘以 \(a^{2^0}\)
2. 第 \(i = 1\) 位是 \(0\)，跳过
3. 第 \(i = 2\) 位是 \(1\)，结果乘以 \(a^{2^2}\)
4. 第 \(i = 3\) 位是 \(1\)，结果乘以 \(a^{2^3}\)

以 Golang 实现为例，优先选择迭代法：

```go
package main

func Pow(a float64, b int) float64 {
	if b < 0 {
		a, b = 1/a, -b
	}

	// iterative method
	var res float64 = 1
	for b > 0 {
		if b&1 == 1 {
			res *= a
		}
		a *= a
		b >>= 1 // b = b / 2
	}
	return res
}

```

考察情形：

| a   | b   | res          |
|-----|-----|--------------|
| any | 0   | 1            |
| 0   | any | 0            |
| 2   | 10  | 1024         |
| 2   | -10 | 0.0009765625 |

> 实现方式有迭代法和递归法，迭代法优越。如果未实现 \(O(\log b)\) 的算法，则不合格。

## 运维

### VSZ 和 RSS

- VSZ (Virtual Set Size) 是进程虚拟内存的大小，包括代码段、数据段、堆和栈。
- RSS (Resident Set Size) 是进程实际使用的物理内存大小，包括代码段、数据段和堆。

| Metrics        | VSZ                                                                                   | RSS                                                              |
|----------------|---------------------------------------------------------------------------------------|------------------------------------------------------------------|
| **Definition** | Total virtual memory allocated to the process                                         | Actual physical memory (RAM) currently being used by the process |
| **Scope**      | All virtual memory, including swapped memory, memory-mapped files, and shared memory	 | Only the pages of the process currently in physical RAM          |
| **Size**       | Larger                                                                                | Smaller                                                          |
| **Use Case**   | Show memory usage of all time                                                         | Realtime RAM usage                                               |

#### 查看进程的 VSZ 和 RSS

```bash
ps -eo pid,%mem,vsz,rss,comm | grep <process_name>
```

#### 查询实时内存使用量前十的进程

```bash
ps -eo pid,%mem,vsz,rss,comm | sort -nk2 -r | head -n 10
```

## 架构

### 云计算

#### 可用区（AZ）是什么？

可用区（Availability Zone）是一个独立的数据中心，通常位于同一地理区域内，但是物理上相互隔离。可用区之间有独立的电力、网络和冷却系统，确保一个可用区的故障不会影响其他可用区。

#### 什么是堡垒机？

堡垒机（Bastion Host）是一种安全工具，用于管理和监控服务器的访问。用户通过堡垒机访问服务器，堡垒机记录用户的操作，提供审计和安全保障。

> 堡垒机也叫跳板机。

### 负载均衡

#### 轮转调度（Round Robin）

轮转调度是一种负载均衡算法，将请求依次分配给服务器列表中的每个服务器。例如，有三个服务器，请求依次分配给
A、B、C、A、B、C、A、B、C...

> 与此类似的算法还有加权轮询，根据服务器的负载情况，分配不同的权重。

### 评价一个系统

四字诀：安全可靠，无廉价美。

![image of system-requirement pyramid](/images/aws-well-architected-framework/aws-well-architected-framework.png "system-requirement-pyramid")

### CDN

CDN（Content Delivery Network）是一种分布式网络，用于缓存和分发静态资源，提高网站性能和用户体验。

1. 客户端与边缘节点之间的距离更近，减少延迟。
2. CDN相当于缓存，降低源服务器的网络带宽，负载。
3. 因此提高源服务的可用性。
4. 提高安全性，CDN可以防止DDoS攻击。

> CDN 本质就是宏观层面的缓存。

### 网络安全

#### 什么是 DDoS 攻击？

> DDoS（Distributed Denial of Service）是一种网络攻击，通过分布式客户端，产生大量请求，使得目标服务过载或由此触发错误，导致服务不可用。

#### 如何防范？

> 设置防火墙，限制访问频率，使用CDN，使用DDoS防护服务，使用IP黑名单，使用CAPTCHA验证。

### 高可用

1. 多地区和多可用区部署。
2. 应用服务器前部署负载均衡。
3. workload自动缩扩容
4. 防灾冗余，包括使用 passive-active，active-active 部署。
5. 充分的监控和报警

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

### 队列理论

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

为什么？排序本质是比较，可以用一个决策树表示，决策树的深度即为最坏情况下的比较次数。对于 \(n\) 个元素的排序，有 \(n!\)
种排列，所以决策树的叶子节点数为 \(n!\)，高度为 \(h\)，满足：

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

## 参考
1. [use-the-index-luke](https://use-the-index-luke.com/)
