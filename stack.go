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
	// Used to get variable types as string representation
	"reflect"
	// Used to lock and unlock the stack when attempting to modify it
	"sync"
)

const (
	Int = iota
	Float
	String
	Bool
)

// Stack stores the lock and data members of the stack.
type Stack struct {
	lock     sync.Mutex
	member   []interface{}
	dataType int
}

// NewStack is the constructor method that initializes a new, empty stack and returns it.
func NewStack(dataType int) *Stack {
	return &Stack{member: make([]interface{}, 0), dataType: dataType}
}

// Push will add a new value to the top of the stack.
func (s *Stack) Push(val interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if err := checkType(s, val); err != nil {
		return err
	}

	s.member = append(s.member, val)
	return nil
}

// Pop attempts to remove a value from the top of the stack. It will return an error if the stack is empty, or
// the value of the top element if it isn't.
func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.member)
	if l == 0 {
		return nil, errors.New("Pop(): Attempted pop from empty stack")
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
		return nil, errors.New("Peek(): Attempted peek from empty stack")
	}
	retVal := s.member[l-1]
	return retVal, nil
}

// Size returns the current size of the stack. If the stack is empty, it will simply return
// 0, not an error
func (s *Stack) Size() int {
	l := len(s.member)
	return l
}

// Drain removes all elements that are currently in the stack.
func (s *Stack) Drain() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.member = nil
}

// checkType checks an incoming member value from Push() and
// validates if it matches the declared type from NewStack().
// returns an error if the types do not match
func checkType(s *Stack, member interface{}) error {
	switch member.(type) {
	case int:
		if s.dataType != Int {
			return errors.New("Push(): expected int, but received " + reflect.TypeOf(member).String())
		}
	case float64:
		if s.dataType != Float {
			return errors.New("Push(): expected float64, but received " + reflect.TypeOf(member).String())
		}
	case string:
		if s.dataType != String {
			return errors.New("Push(): expected string, but received " + reflect.TypeOf(member).String())
		}
	case bool:
		if s.dataType != Bool {
			return errors.New("Push(): expected bool, but received " + reflect.TypeOf(member).String())
		}
	default:
		return errors.New("Push(): unknown data type received")
	}
	return nil
}

// ToInt is intended to be a wrapping function around Pop() or Peek()
// so that the int variable can be explicitely returned to the user.
// Eliminates the need for a user side type assertion
func ToInt(memberReturned interface{}, err error) (int, error) {
	if err != nil {
		return 0, err
	} else {
		return memberReturned.(int), err
	}
}

// ToFloat behaves equivalently to ToInt, but operates with
// float64 types
func ToFloat(memberReturned interface{}, err error) (float64, error) {
	if err != nil {
		return 0.0, err
	} else {
		return memberReturned.(float64), err
	}
}

// ToString behaves equivalently to ToInt, but operates with
// string types
func ToString(memberReturned interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	} else {
		return memberReturned.(string), err
	}
}

// ToBool behaves equivalently to ToInt, but operates with
// bool types
func ToBool(memberReturned interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	} else {
		return memberReturned.(bool), err
	}
}
