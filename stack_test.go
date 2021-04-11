package stack

import (
	"encoding/json"
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack(Int)
	// should not error
	err := s.Push(1)
	if err != nil {
		t.Errorf("[In TestPush()] s.Push(1) - expected nil error, but got %s\n", err)
	}
	// this should error
	err = s.Push("Hello, World!")
	if err == nil {
		t.Errorf("[In TestPush()] s.Push(\"Hello, World!\") - expected error, but got nil\n")
	}
}

func TestGenericPush(t *testing.T) {
	gs := NewGenericStack()
	a := struct {
		foo string
		bar int
	}{
		foo: "Hello, World!",
		bar: 4,
	}
	gs.Push(a)
}

func TestPop(t *testing.T) {
	s := NewStack(Int)
	// add values
	_ = s.Push(1)
	// should not error
	val, err := s.Pop()
	if err != nil {
		t.Errorf("[In TestPop()] s.Pop() - expected nil error, but got %s\n", err)
	}
	if val.(int) != 1 {
		t.Errorf("[In TestPop()] s.Pop() - expected 1, but got %d\n", val.(int))
	}
	// should error
	_, err = s.Pop()
	if err == nil {
		t.Errorf("[In TestPop()] s.Pop() - expected \"Pop from empty stack\" error, but got nil\n")
	}
}

func TestPeek(t *testing.T) {
	s := NewStack(Int)
	// should error
	_, err := s.Peek()
	// test if it doesn't
	if err == nil {
		t.Errorf("[In TestPeek()] s.Peek() - expected \"Peek from empty stack\" error, but got nil\n")
	}
	s.Push(5)
	val, _ := s.Peek()
	if val.(int) != 5 {
		t.Errorf("[In TestPeek()] s.Peek() - expected 5, but got %d\n", val)
	}
}

func TestSize(t *testing.T) {
	s := NewStack(Int)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	sz := s.Size()
	if sz != 3 {
		t.Errorf("[In TestSize()] s.Size() - expected 3, but got %d\n", sz)
	}
}

func TestDrain(t *testing.T) {
	s := NewStack(Int)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Drain()
	sz := s.Size()
	if sz != 0 {
		t.Errorf("[In TestDrain()] s.Drain() - expected 0, but got %d\n", sz)
	}
}

func TestToInt(t *testing.T) {
	s := NewStack(Int)
	s.Push(1)
	val, err := ToInt(s.Pop())
	// shouldn't error
	if err != nil {
		t.Errorf("[In TestToInt()] ToInt(s.Pop()) - expected nil error, but got %s\n", err)
	}
	if val != 1 {
		t.Errorf("[In TestToInt()] ToInt(s.Pop()) - expected 1, but got %d\n", val)
	}
	// now should error
	_, err = ToInt(s.Peek())
	if err == nil {
		t.Errorf("[In TestToInt()] ToInt(s.Pop()) - expected \"Pop from empty stack\" error, but got nil\n")
	}
}

func TestToFloat(t *testing.T) {
	s := NewStack(Float)
	s.Push(3.14)
	val, err := ToFloat(s.Pop())
	// shouldn't error
	if err != nil {
		t.Errorf("[In TestToFloat()] ToFloat(s.Pop()) - expected nil error, but got %s\n", err)
	}
	if val != 3.14 {
		t.Errorf("[In TestToFloat()] ToFloat(s.Pop()) - expected 3.14, but got %f\n", val)
	}
	// now should error
	_, err = ToFloat(s.Peek())
	if err == nil {
		t.Errorf("[In TestToFloat()] ToFloat(s.Pop()) - expected \"Pop from empty stack\" error, but got nil\n")
	}
}

func TestToString(t *testing.T) {
	s := NewStack(String)
	s.Push("Hello")
	val, err := ToString(s.Pop())
	// shouldn't error
	if err != nil {
		t.Errorf("[In TestToString()] ToString(s.Pop()) - expected nil error, but got %s\n", err)
	}
	if val != "Hello" {
		t.Errorf("[In TestToString()] ToString(s.Pop()) - expected \"Hello\", but got %s\n", val)
	}
	// now should error
	_, err = ToString(s.Peek())
	if err == nil {
		t.Errorf("[In TestToString()] ToString(s.Pop()) - expected \"Pop from empty stack\" error, but got nil\n")
	}
}

func TestToBool(t *testing.T) {
	s := NewStack(Bool)
	s.Push(true)
	val, err := ToBool(s.Pop())
	// shouldn't error
	if err != nil {
		t.Errorf("[In TestToBool()] ToBool(s.Pop()) - expected nil error, but got %s\n", err)
	}
	if val != true {
		t.Errorf("[In TestToBool()] ToBool(s.Pop()) - expected true, but got %t\n", val)
	}
	// now should error
	_, err = ToBool(s.Peek())
	if err == nil {
		t.Errorf("[In TestToBool()] ToBool(s.Pop()) - expected \"Pop from empty stack\" error, but got nil\n")
	}
}

func TestCheckType(t *testing.T) {
	s1 := NewStack(Int)
	s2 := NewStack(Float)
	s3 := NewStack(String)
	s4 := NewStack(Bool)

	// these should all error
	err := s1.Push(3.14)
	if err == nil {
		t.Errorf("[In TestCheckType()] s1.Push() - expected float type assertion error, but got nil")
	}
	err = s2.Push("Hello")
	if err == nil {
		t.Errorf("[In TestCheckType()] s2.Push() - expected string type assertion error, but got nil")
	}
	err = s3.Push(true)
	if err == nil {
		t.Errorf("[In TestCheckType()] s3.Push() - expected bool type assertion error, but got nil")
	}
	err = s4.Push(3)
	if err == nil {
		t.Errorf("[In TestCheckType()] s4.Push() - expected int type assertion error, but got nil")
	}
	// to trigger default case, send JSON encoded data
	s := struct {
		a int
		b string
	}{
		4,
		"FooBar",
	}
	sEnc, _ := json.Marshal(s)
	err = s1.Push(sEnc)
	if err == nil {
		t.Errorf("[In TestCheckType()] s1.Push() - expected JSON type assertion error, but got nil")
	}
}
