// Package stack aims to provide a fast, simple, thread safe, and user friendly stack library.
// Stacks that are composed of primitive types (int, float64, string, bool) inherit from their
// parent GenericStacks, which can be of any data type. GenericStacks must be type asserted
// when retrieving their values, however, primitive Stacks include wrapping functions that will
// easily return the intended types for you. This way, completeness and user friendliness
// (through avoiding user-side type assertions when possible) is achieved.

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

var typesStringify = map[int]string{
	Int:    "int",
	Float:  "float",
	String: "string",
	Bool:   "bool",
}

// GenericStack stores the lock and data members
type GenericStack struct {
	lock   sync.Mutex
	member []interface{}
}

// Stack inherits GenericStack's previous fields, and adds a dataType field to store what primitive data type the stack
// should accept
type Stack struct {
	GenericStack
	dataType int
}

// NewGenericStack is the constructor method that initializes a new, empty generic stack. Generic stacks can accept any
// data type, however, a type assertion will have to be included in a user-side program to retrieve values from the stack
func NewGenericStack() *GenericStack {
	return &GenericStack{member: make([]interface{}, 0)}
}

// NewStack is the constructor method that initializes a new, empty stack of a specific type and returns it
func NewStack(dataType int) *Stack {
	return &Stack{GenericStack: GenericStack{member: make([]interface{}, 0)}, dataType: dataType}
}

// Push (GenericStack method) will add a new generic value to the top of the stack.
func (s *GenericStack) Push(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.member = append(s.member, val)
}

// Push (Stack method) will add a new value to the top of the stack, and also checks for type congruency.
// Returns an error if an attempted Push of a new element is not the same type as the declared Stack
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
func (s *GenericStack) Pop() (interface{}, error) {
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
func (s *GenericStack) Peek() (interface{}, error) {
	l := len(s.member)
	if l == 0 {
		return nil, errors.New("Peek(): Attempted peek from empty stack")
	}
	retVal := s.member[l-1]
	return retVal, nil
}

// Size returns the current size of the stack. If the stack is empty, it will simply return
// 0, not an error
func (s *GenericStack) Size() int {
	l := len(s.member)
	return l
}

// Drain removes all elements that are currently in the stack.
func (s *GenericStack) Drain() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.member = nil
}

// checkType checks an incoming member value from Push() and
// validates if it matches the declared type from NewStack().
// returns an error if the types do not match
func checkType(s *Stack, member interface{}) error {
	e := errors.New("Push(): expected: [" + typesStringify[s.dataType] + "]; received: [" + reflect.TypeOf(member).String() + "]")
	switch member.(type) {
	case int:
		if s.dataType != Int {
			return e
		}
	case float64:
		if s.dataType != Float {
			return e
		}
	case string:
		if s.dataType != String {
			return e
		}
	case bool:
		if s.dataType != Bool {
			return e
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
