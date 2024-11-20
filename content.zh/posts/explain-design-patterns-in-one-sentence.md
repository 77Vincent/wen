---
title: 用Golang理解几种设计模式
date: 2023-10-29T02:01:58+05:30
tags: [computer-science, golang, english]
categories: study 
canonicalUrl: https://wenstudy.com/posts/explain-design-patterns-one-sentence/
---

大部分模式都有一个共同点，那就是通过接口连接生产者和消费者。接口的抽象是关键，它反映了对底层类型的深刻理解。知道这一点可以简化事情，无论我们处理的是什么场景，或者使用的是什么语言。

至于为什么现在有这些模式，还有多少尚未被发现，以及背后的数学本质是什么，我也还无法回答。

## 创建型 - 实例如何创建

### 工厂

工厂方法负责创建一系列实现同一接口的实例。返回类型是接口而不是具体类。

```go
// type A, B, C all implement I
type I interface {}

func factory(string t) I {
  switch t {
    case "A":
      return new A()
    case "B":
      return new B()
    case "C":
      return new C()
  }
}
```

### 单例

在应用程序生命周期中只有一个实例。

```go
type A struct{}

var (
  one *A
  m   = &sync.Mutex{}
)

func Get() {
  // prevents unnecessary lock
  if one == nil {
    m.Lock()
    defer m.Unlock()

    // prevents unnecessary creation
    if one == nil {
      one = &A{}
      return one
    }
    return one
  }
  return one
}
```

### 原型

实例通过对一个母体克隆而不是实例化获得。它们的类实现了一个公共接口，该接口定义了一个克隆方法，因此返回类型是接口而不是具体类。

```go
type I interface {
  clone() I
}

type A struct {}

func (a *A) clone() I {
  tmp := *a
  return &tmp
}

func main() {
  var a, a2 I
  a  = &A{}
  a2 = a.clone()
}
```

## 结构型 - 如何组织对象

### 桥接

类A将任务委托给类B。它们共享一个公共接口。

```go
type I interface {
  do() string
}

type A struct {}
func (a A) do() string {
  return "hello"
}

type B struct {
  i I
}
func (b B) do() {
  return b.i.do()
}

func main() {
  b:= B{
    i: A{}
  }
  
  b.do() // => "hello"
}
```

### 适配器

类A通过适配器与不兼容的类B进行交互。

```go
type A interface {
  produce() int
}

type B interface {
  consume(string s)
}

type AtoBAdapter struct {
  A
}

func (a AtoBAdapter) produce() string {
  return strconv.Itoa(a.A.produce())
}

func main() {
  a := AtoBAdapter{
    A: A{}
  }
  b := B{}

  b.consume(a.produce())
}
```

### 组合

如果业务模型是递归的，可以使用树形结构。

```go
type I interface {
  do()
}

type Component struct {
  children []I // can be component or node
}

func (c Component) do() {
  for _, child := range r.children {
    child.do()
  }
}

type Node struct {}

func (n Node) do() {}

func main() {
  c := Component{
    children: []I{
      Component{},
      Node{},
    }
  }
}
```

### 立面

为单一入口定义了新接口。在其后，是一个涉及许多类的复杂子系统，但新的立面（控制面板）掩盖了背后的细节。

```go
// the only interface for user interaction
type I interface {
  foo()
  bar()
}

type A interface {
  a1()
  a2()
}

type B interface {
  b1()
}

type system struct {
  A
  B
}
func (s system) foo() {
  s.A.a1()
  s.B.b1()
}
func (s system) bar() {
  s.A.a2()
}

func main() {
  s I := system{
    A: A{}
    B: B{} 
  }

  s.foo()
  s.bar()
}
```

## 行为型 - 如何产生变化

### 策略

上下文从多个实现了共同接口的类中选择一个实例。

```go
// type A, B, C all implement I
type I interface {
  do()
}

func main() {
  var executor I

  switch strategy {
    case "A":
      executor := new A()
    case "B":
      executor := new B()
    case "C":
      executor := new C()
  }

  executor.do()
}
```

### 模板

A类是B C D类的模板，可以在定义时重写以产生变化（而不是在运行时）。

```go
type interface I {
  foo()
  bar()
}

// a default implementation of I
type template struct {}
func (t template) foo() {}
func (t template) bar() {}

// only overwrites foo
type A struct {
  template
}
func (a A) foo() {
  fmt.Print("a foo")
}
func (a A) bar() {
  a.template.bar()
}

// only overwrites bar
type B struct {
  template
}
func (b B) foo() {
  b.template.foo()
}
func (b B) bar() {
  fmt.Print("b bar")
}
```

### 装饰

B装饰A以改变结果，但不改变A的原始行为。

```go
type I interface {
  do()
}

type A struct {}

func (a A) do() int {
  return 1
}

type B struct {
  i I
}

func (b B) do() int {
  return b.i.do() + 1
}

func main() {
  a := A{}
  b := B{
    i: a
  }

  a.do() // => 1
  b.do() // => 2
}
```
