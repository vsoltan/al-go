package stack

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	s := New()

	if is_empty := s.Size() == 0; !is_empty {
		t.Errorf("Constructed stack has non-zero size")
	}
}

func TestAppend(t *testing.T) {
	s := New()

	s.Append("Hello")
	s.Append("World")
	s.Append("Foo")
	s.Append("Bar")

	if got := s.Size(); got != 4 {
		t.Errorf("Size = %d; want 4", got)
	}

	if got, ok := s.Peek(); !ok {
		t.Errorf("Tried peeking from empty stack!")
	} else if got != "Bar" {
		t.Errorf("Top = %s; want \"Bar\"", got)
	}
}

func TestPop(t *testing.T) {
	s := New()

	s.Append("Foo")
	s.Append("Bar")

	if got, ok := s.Pop(); !ok {
		t.Errorf("Tried popping from empty stack!")
	} else if got != "Bar" {
		t.Errorf("Top = %s; want \"Bar\"", got)
	}

	if got, ok := s.Pop(); !ok {
		t.Errorf("Tried popping from empty stack!")
	} else if got != "Foo" {
		t.Errorf("Top = %s; want \"Foo\"", got)
	}

	if is_empty := s.IsEmpty(); !is_empty {
		t.Errorf("Size: %d; want 0", s.Size())
	}
}
