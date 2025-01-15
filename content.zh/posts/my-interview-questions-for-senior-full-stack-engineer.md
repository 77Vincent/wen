---
title: 软件工程师面试问题
date: 2025-01-01T02:01:58+05:30
tags: [ computer-science, interview, engineering ]
categories: study
canonicalUrl: https://wenstudy.com/posts/my-interview-questions-for-senior-full-stack-engineer/
---

由于工作中经常需要为团队招聘进行面试，以下总结了一份问题集，针对中高级全栈工程师，从基础通用知识，前端、后端到架构。以经典为主，主打永不过时。

## 计算机网络

### CIDR

以下 `CIDR` 地址块有多少个可用的IP地址？

```
10.0.0.0/24
```

2^(32-24)=2^8=256

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

### 什么是竞态条件？

是多个线程或进程并发执行，修改同一个变量时，由于执行顺序不确定，可能导致：

1. 数据不一致，每次执行可能得到不同的结果。
2. 死锁，或数据处于非法状态。
3. 难以复现和调试。

> 重点是执行顺序不确定，而非总是错误，比如可能有时候按123顺序正确执行，有时候混乱。
