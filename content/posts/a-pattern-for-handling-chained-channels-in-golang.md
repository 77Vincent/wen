---
title: A pattern for handling chained channels in Golang
date: 2023-11-07T02:01:58+05:30
description: When we use channels to serialize tasks, we create an asynchronous flow. It is easy to mess things up in asynchronous programming, especially if we have additional requirements like timeout and cancelation.
tags: [computer-science, english]
featured_image: https://wen-images.s3.ap-northeast-1.amazonaws.com/blog/a-pattern-for-handling-chained-channels-in-golang/a-pattern-for-handling-chained-channels-in-golang.webp
categories: study 
canonicalUrl: https://wenstudy.com/posts/a-pattern-for-handling-chained-channels-in-golang/
---

When we use channels to serialize tasks, we create an asynchronous flow. It is easy to mess things up in asynchronous programming, especially if we have additional requirements like timeout and cancelation.

A real-world example of chained channels is reading data line by line from a file reader and then passing it to a text parser, which also works asynchronously.

I found a common pattern that may cover most such scenarios with minimum and understandable codes.

> Asynchronous provider followed by a synchronous consumer.

Concept:

```go
go provider()
consumer()
```

Example:

```go
var (
  provider         = make(chan int)
  consumer         = make(chan int)
  ctx, cancel      = context.WithTimeout(context.Background, time.Second)
)
defer cancel()

// provider
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

// follow-up logic
```

The whole process has a blocking nature due to the last consumer for-loop. This is desired because we want to ensure the program completes the task before quitting.
It is also simple to handle errors, like how I return/break when the channel is not OK. For example, in the consumer loop:

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

// follow-up logic
if err != nil {
    log.Error(err)
}
```

At last, may we all handle channels well.
