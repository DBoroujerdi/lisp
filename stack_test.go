package main

import "testing"

func TestStack(t *testing.T) {
	var s = new(Stack)
	s.push(5)

	v, _ := s.pop()

	if v != 5 {
		t.Error("Expected 5 got ", v)
	}

	e := s.isEmpty()
	if !e {
		t.Error("Expected empty stack after pop()")
	}
}

func TestStackMultiple(t *testing.T) {
	var s = new(Stack)
	s.push(5)
	s.push(6)
	s.push(7)

	e_1, _ := s.pop()
	if e_1 != 7 {
		t.Error("Expected 7 got ", e_1)
	}

	e_2, _ := s.pop()
	if e_2 != 6 {
		t.Error("Expected 6 got ", e_2)
	}

	e_3, _ := s.pop()
	if e_3 != 5 {
		t.Error("Expected 5 got ", e_3)
	}

	e := s.isEmpty()
	if !e {
		t.Error("Expected empty stack after pop()")
	}
}

func TestPopOnEmptyStack(t *testing.T) {
	var s = new(Stack)
	val, err := s.pop()

	if val != nil {
		t.Error("Top item on empty stack expected nil but was ", val)
	}

	if err == nil {
		t.Error("Error should not be nil when popping on empty stack")
	}

	if err.Error() != "Nothing on the stack to pop!" {
		t.Error("Error message not as expected. Actual =", err.Error())
	}
}
