package stack

import (
	"testing"
)

func TestPop(t *testing.T) {
	s := NewStack()
	// should error
	got, err := s.Pop()
	// test if it doesn't
	if err == nil {
		t.Errorf("s.Pop() - expected error, got nil\n")
	}
	s.Push(1)
	got, _ = s.Pop()
	val := got.(int)
	if val != 1 {
		t.Errorf("s.Pop() - expected 1, got %d\n", val)
	}
}

func TestPeek(t *testing.T) {
	s := NewStack()
	// should error
	got, err := s.Peek()
	// test if it doesn't
	if err == nil {
		t.Errorf("s.Peek() - expected error, got nil\n")
	}
	s.Push(5)
	got, _ = s.Peek()
	val := got.(int)
	if val != 5 {
		t.Errorf("s.Peek() - expected 5, got %d\n", val)
	}
}

func TestSize(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val := s.Size()
	if val != 3 {
		t.Errorf("s.Size() - expected 3, got %d\n", val)
	}
}

func TestDrain(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Drain()
	val := s.Size()
	if val != 0 {
		t.Errorf("s.Drain() - expected 0, got %d\n", val)
	}
}
