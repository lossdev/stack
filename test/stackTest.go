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
