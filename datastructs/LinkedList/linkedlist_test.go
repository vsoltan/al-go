package LinkedList 

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	list := NewLinkedList()
	if got := list.front == nil; got != true {
		t.Errorf("constructed list has non-nil head")
	}

	if got := list.back == nil; got != true {
		t.Errorf("constructed list has non-nil tail")
	} 

	if got := list.length; got != 0 {
		t.Errorf("constructed list has non-zero length")
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkedList()

	list.AppendFront(1)
	list.AppendBack(2)
	list.AppendFront(3)

	if got := list.front.data; got != 3 {
		t.Errorf("Head = %d; want 3", got)
	}

	if got := list.back.data; got != 2 {
		t.Errorf("Tail = %d; want 2", got)
	}
}

func TestRemove(t *testing.T) {
	list := NewLinkedList()

	list.AppendFront(1)
	list.AppendBack(2)
	list.AppendFront(3)

	list.RemoveFront()
	list.RemoveBack() 

	if got := list.length; got != 1 {
		t.Errorf("Length = %d; want 1", got)
	}

	if got := list.front.data; got != 1 {
		t.Errorf("Head = %d; want 1", got)
	}

	if got := list.back.data; got != 1 {
		t.Errorf("Tail = %d; want 1", got)
	}
}