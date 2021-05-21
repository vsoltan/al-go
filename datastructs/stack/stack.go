package stack

import (
	"github.com/vsoltan/al-go/datastructs/linkedlist"
)

type Stack struct {
	list *linkedlist.List
}

func New() *Stack {
	return &Stack{list: linkedlist.New()}
}

func (s *Stack) Append(i interface{}) {
	s.list.AppendFront(i)
}

func (s *Stack) Peek() (value interface{}, ok bool) {
	if s.Size() <= 0 {
		return
	}
	value = s.list.Front().Value()
	ok = true
	return
}

func (s *Stack) Pop() (value interface{}, ok bool) {
	if s.Size() <= 0 {
		return
	}
	value, ok = s.list.RemoveFront()
	return
}

func (s *Stack) Size() int {
	return s.list.Length()
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
