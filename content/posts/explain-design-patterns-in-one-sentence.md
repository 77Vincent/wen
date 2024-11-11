---
title: Explain design patterns in one sentence
date: 2023-10-29T02:01:58+05:30
description: With minimum Golang sample code
tags: [computer-science, golang, english]
featured_image: https://wen-images.s3.ap-northeast-1.amazonaws.com/blog/explain-design-patterns-in-one-sentence/explain-design-patterns-in-one-sentence.webp
categories: study 
canonicalUrl: https://wenstudy.com/posts/explain-design-patterns-one-sentence/
---

![image of common pattern](https://wen-images.s3.ap-northeast-1.amazonaws.com/blog/explain-design-patterns-in-one-sentence/explain-design-patterns-in-one-sentence.webp "common pattern")

Most of these patterns have one thing in common, which is to connect producers and consumers through the interface. The abstraction of the interface is the key. It reflects a deep understanding of the underlying type. Knowing this can simply things, no matter what scenario we are dealing with, or what language is being used.

As for why currently there are these patterns, how many there are yet to be discovered, and what mathematical nature is behind them, I can't answer now.

# Creational - How instances are created

## Factory

The factory method is in charge of creating a series of instances that all implement one common interface. The return type is the interface instead of a concrete class.

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

## Singleton

One instance throughout the application life-cycle.

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

## Prototype

Instances are obtained by cloning but not instantiating. Their class implements a common interface that defines a clone method, so the return type is an interface but not the concrete class.

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

# Structural - How components collaborate

## Bridge

Class A delegates tasks to class B. They share a common interface.

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

## Adapter

Class A interacts with incompatible class B through an adapter.

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

## Composite

Using a tree-like structure if the business model is recursive.

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

## Facade

A new interface is defined for a single-entry-point control. Behind that, it is a complex subsystem that involves many classes.

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

# Behavioral - How variation is produced

## Strategy

The context chooses one of the instances from multiple classes that all implement a common interface.

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

## Template

Class A is a template for classes B C and D to inherit and overwrite to produce variations during definition (not at runtime).

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

## Decorator

B decorates(wraps) A to mutate the result without changing A's original behavior.

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
