# stack
[![GoDoc](https://godoc.org/github.com/bmw417/stack?status.png)](http://godoc.org/github.com/bmw417/stack)

Lightweight, Simple, Quick, Thread-Safe Golang Stack Implementation


## Purpose

Provide a fast, thread safe, and generic Golang Stack API with minimal external linkage
and maximum performance and usability.


## Installation

``` bash

go get -d -v github.com/bmw417/stack

```

## Example

A copy of this file is contained in the `test` directory. Note that since this library uses generics, an explicit type assertion is required.

``` Go

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bmw417/stack"
)

func main() {
	s := stack.NewStack()
	s.Push("foo")
	s.Push("bar")
	s.Push("baz")

	fmt.Println("Size: " + strconv.Itoa(s.Size()))

	val, err := s.Pop()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Pop: " + val.(string))
	val, err = s.Pop()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Pop: " + val.(string))
	val, err = s.Peek()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Peek: " + val.(string))
	val, err = s.Pop()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Pop: " + val.(string))
	fmt.Println("Size: " + strconv.Itoa(s.Size()))

	s.Push("foo")
	s.Push("bar")
	s.Push("baz")
	s.Drain()

	fmt.Println("Size: " + strconv.Itoa(s.Size()))
	s.Push("foo")
	fmt.Println("Size: " + strconv.Itoa(s.Size()))
	val, err = s.Peek()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Peek: " + val.(string))
	s.Drain()
	val, err = s.Peek()
	if err != nil {
		log.Println(err)
	}
}

```

``` bash

$ go run stackTest.go
Size: 3
Pop: baz
Pop: bar
Peek: foo
Pop: foo
Size: 0
Size: 0
Size: 1
Peek: foo
2020/05/19 10:42:17 Peek from Empty Stack

```
