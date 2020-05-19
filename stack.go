// Package stack aims to provide a fast, simple, thread safe, and generic golang stack library
// to the public. It uses as few linking libraries as possible, and the ones it does are
// common standard libraries for a smaller code footprint and greater performance.

// Each Stack is generically compatible, meaning that any data can be added and retrieved to/from a Stack.
// Locks are also used on Stack objects when modifying the Stack through using the Pop or Push methods, so
// that Stack objects are thread safe.
package stack

// Minimal imports are used, and those that are used are
// present in the standard library.
import (
	// Used for throwing errors
	"errors"
	// Used to lock and unlock the stack when attempting to modify it
	"sync"
)

// Stack stores the lock and data members of the stack.
type Stack struct {
	lock   sync.Mutex
	member []interface{}
}

// NewStack is the constructor method that initializes a new, empty stack and returns it.
func NewStack() *Stack {
	return &Stack{member: make([]interface{}, 0)}
}

// Push will add a new value to the top of the stack.
func (s *Stack) Push(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.member = append(s.member, val)
}

// Pop attempts to remove a value from the top of the stack. It will return an error if the stack is empty, or
// the value of the top element if it isn't.
func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.member)
	if l == 0 {
		return nil, errors.New("Pop from Empty Stack")
	}

	retVal := s.member[l-1]
	s.member = s.member[:l-1]
	return retVal, nil
}

// Peek looks at, but does not remove, the top element of the stack. It returns an error if
// the stack is empty, or the value of the top element if it isn't.
func (s *Stack) Peek() (interface{}, error) {
	l := len(s.member)
	if l == 0 {
		return nil, errors.New("Peek from Empty Stack")
	}
	retVal := s.member[l-1]
	return retVal, nil
}

// Size returns the current size of the stack. If the stack is empty, it will simply return
// 0, not an error.
func (s *Stack) Size() int {
	l := len(s.member)
	return l
}

// Drain removes all elements that are currently in the stack.
func (s *Stack) Drain() {
	s.member = nil
}
