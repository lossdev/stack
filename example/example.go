package main

import (
	"fmt"
	"github.com/lossdev/stack"
	"log"
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
