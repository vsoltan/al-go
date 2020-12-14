package Queue 

import (
	"testing" 
)

func TestConstructor(t *testing.T) {
	q := New() 

	if is_empty := q.Size() == 0; !is_empty {
		t.Errorf("Constructed stack has non-zero size")
	}
}

// func TestEnqueue(t *testing.T) {
	
// }

// func TestDequeue(t *testing.T) {
	
// }

// func Test

