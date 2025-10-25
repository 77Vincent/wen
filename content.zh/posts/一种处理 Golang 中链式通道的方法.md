---
title: 一种处理 Golang 中链式通道的方法
date: 2023-11-07T02:01:58+05:30
tags: [golang]
categories: tech
canonicalUrl: https://wenstudy.com/posts/一种处理 Golang 中链式通道的方法/
---

当使用通道来串行化任务时，我们因此创建了一个异步流。在异步编程中很容易搞砸，尤其是有额外的要求时，比如超时和取消。 真实世界的例子，是从文件读取器逐行读取数据，然后传递给另一个异步工作的文本解析器。

我找到了一个通用模式，可以用易于理解的代码覆盖大多数这样的场景。

> 异步生产者后跟同步消费者。

<!--more-->
形如：

```go
go provider()
consumer()
```

例如：

```go
var (
  provider         = make(chan int)
  consumer         = make(chan int)
  ctx, cancel      = context.WithTimeout(context.Background, time.Second)
)
defer cancel()

// 生产者
go func() {
    defer close(consumer)

    for {
        select {
        case <-ctx.Done():
            return
        case v, ok := <-provider:
            if !ok {
                return
            }
            consumer <- v
        }
    }
}()

consume:
for {
    select {
    case <-ctx.Done():
        return
    case v, ok := <-consumer:
        if !ok {
            break consume
        }
    }
}

// 后续逻辑
```

这里，我们有一个`生产者`和一个`消费者`。`生产者`是一个 goroutine，它从`生产者`通道中读取数据并将其传递给`消费者`通道。`消费者`是一个阻塞的 for 循环，它从`消费者`通道中读取数据。

```go
var err error

consume:
for {
    select {
    case <-ctx.Done():
        return
    case v, ok := <-consumer:
        if !ok {
            break consume
        }

        if err = process(v); err != nil {
            break consume
        }
    }
}

// 后续逻辑
if err != nil {
    log.Error(err)
}
```

祝我们在异步编程中好运！
