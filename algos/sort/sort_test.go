package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	toSort := []int{5, 3, 4, 1, 2}
	expected := []int{1, 2, 3, 4, 5}
	BubbleSort(toSort)

	if got := reflect.DeepEqual(toSort, expected); !got {
		t.Errorf("Sorted array is different from expected; is %v", toSort)
	}
}

func TestMergeSort(t *testing.T) {
	toSort := []int{5, 3, 4, 1, 2}
	expected := []int{1, 2, 3, 4, 5}
	sorted := MergeSort(toSort)

	if got := reflect.DeepEqual(sorted, expected); !got {
		t.Errorf("Sorted array is different from expected; is %v", sorted)
	}
}

func TestMergeSort2(t *testing.T) {
	toSort := []int{4, 1, 2}
	expected := []int{1, 2, 4}
	sorted := MergeSort(toSort)

	if got := reflect.DeepEqual(sorted, expected); !got {
		t.Errorf("Sorted array is different from expected; is %v", sorted)
	}
}
