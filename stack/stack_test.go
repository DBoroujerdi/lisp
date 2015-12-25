package stack

import "testing"

func TestStack(t *testing.T) {
	var s = new(Stack)
	s.Push(5)

	v, _ := s.Pop()

	if v != 5 {
		t.Error("Expected 5 got ", v)
	}

	e := s.IsEmpty()
	if !e {
		t.Error("Expected empty stack after pop()")
	}
}

func TestStackMultiple(t *testing.T) {
	var s = new(Stack)
	s.Push(5)
	s.Push(6)
	s.Push(7)

	e_1, _ := s.Pop()
	if e_1 != 7 {
		t.Error("Expected 7 got ", e_1)
	}

	e_2, _ := s.Pop()
	if e_2 != 6 {
		t.Error("Expected 6 got ", e_2)
	}

	e_3, _ := s.Pop()
	if e_3 != 5 {
		t.Error("Expected 5 got ", e_3)
	}

	e := s.IsEmpty()
	if !e {
		t.Error("Expected empty stack after pop()")
	}
}

func TestPopOnEmptyStack(t *testing.T) {
	var s = new(Stack)
	val, err := s.Pop()

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

func TestStackSize(t *testing.T) {
	var s = new(Stack)

	if s.Size() != 0 {
		t.Error("Stack expected to be empty but had size =", s.Size())
	}

	s.Push('c')

	if s.Size() != 1 {
		t.Error("Stack expected to be size 1, but was=", s.Size())
	}

	s.Push('u')

	if s.Size() != 2 {
		t.Error("Stack expected to be size 2, but was=", s.Size())
	}

	s.Push('n')

	if s.Size() != 3 {
		t.Error("Stack expected to be size 3, but was=", s.Size())
	}

	s.Push('t')

	if s.Size() != 4 {
		t.Error("Stack expected to be size 4, but was=", s.Size())
	}
}
