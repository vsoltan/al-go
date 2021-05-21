package linkedlist

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	list := New()
	if nilFront := list.Front() == nil; !nilFront {
		t.Errorf("constructed list has non-nil head")
	}

	if nilBack := list.Back() == nil; !nilBack {
		t.Errorf("constructed list has non-nil tail")
	}

	if isEmpty := list.Length() == 0; !isEmpty {
		t.Errorf("constructed list has non-zero length")
	}
}

func TestAppend(t *testing.T) {
	list := New()

	list.AppendFront(1)
	list.AppendBack(2)
	list.AppendFront(3)

	if got := list.Front().Value(); got != 3 {
		t.Errorf("Head = %d; want 3", got)
	}

	if got := list.Back().Value(); got != 2 {
		t.Errorf("Tail = %d; want 2", got)
	}
}

func TestRemove(t *testing.T) {
	list := New()

	list.AppendFront(1)
	list.AppendBack(2)
	list.AppendFront(3)

	list.RemoveFront()
	list.RemoveBack()

	if got := list.Length(); got != 1 {
		t.Errorf("Length = %d; want 1", got)
	}

	if got := list.Front().Value(); got != 1 {
		t.Errorf("Head = %d; want 1", got)
	}

	if got := list.Back().Value(); got != 1 {
		t.Errorf("Tail = %d; want 1", got)
	}
}
