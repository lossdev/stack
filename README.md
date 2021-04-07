# stack

[![GoDoc](https://godoc.org/github.com/lossdev/stack?status.png)](http://godoc.org/github.com/lossdev/stack)
[![Go Report Card](https://goreportcard.com/badge/github.com/lossdev/stack)](https://goreportcard.com/report/github.com/lossdev/stack)
[![Go Coverage](https://codecov.io/github/lossdev/stack/coverage.svg)](https://codecov.io/github/lossdev/stack/)

Lightweight, Simple, Quick, Thread-Safe Golang Stack Implementation


## Purpose

Provide a fast, thread safe, and generic Golang Stack API with minimal external linkage
and maximum performance and usability.


## Installation

``` bash
go get -d -v github.com/lossdev/stack
```

## Example

``` Go
package main

import (
    "github.com/lossdev/stack"
    "log"
    "fmt"
)

type foo struct {
    bar string
    baz bool
}

func main() {
    // declare a new Stack 's' with int type (stack.Int)
    s := stack.NewStack(stack.Int)
    if err := s.Push(1); err != nil {
        log.Println(err)
    }
    if recv, err := stack.ToInt(s.Peek()); err != nil {
        log.Println(err)
    } else {
        fmt.Println(recv)
    }
    // Adding a member of a different type than what s is declared as will error
    if err := s.Push("Hello, World!"); err != nil {
        log.Println(err)
    }
    gs := stack.NewGenericStack()
    f := foo{"Hello, World!", true}
    gs.Push(f)
    if recv, err := gs.Peek(); err != nil {
        log.Println(err)
    } else {
        // type assertion needed
        frecv := recv.(foo)
        fmt.Printf("GenericStack: {%s, %t}\n", frecv.bar, frecv.baz)
    }
}
```

``` bash
$ go run example.go
1
2021/04/07 13:31:52 Push(): expected: [int]; received: [string]
GenericStack: {Hello, World!, true}
```
