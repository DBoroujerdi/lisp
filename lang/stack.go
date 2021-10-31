package lisp

import "errors"

type Stack struct {
	top *Element
}

type Element struct {
	next  *Element
	value interface{}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("Nothing on the stack to pop!")
	}

	v := s.top.value
	newTop := s.top.next
	s.top = newTop
	return v, nil
}

func (s *Stack) Push(v interface{}) {
	el := new(Element)
	el.value = v
	el.next = s.top
	s.top = el
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Size() int {
	size := 0
	var next *Element = s.top

	for next != nil {
		size++
		next = next.next
	}

	return size
}
