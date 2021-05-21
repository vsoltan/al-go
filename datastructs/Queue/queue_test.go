package queue

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	q := New()

	if isEmpty := q.Size() == 0; !isEmpty {
		t.Errorf("Constructed stack has non-zero size")
	}
}

// func TestEnqueue(t *testing.T) {

// }

// func TestDequeue(t *testing.T) {

// }

// func Test
