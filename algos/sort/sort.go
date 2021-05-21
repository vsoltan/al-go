package sort

// Swap swaps the values of indices i and j in slice list
func Swap(list []int, i int, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}

// BubbleSort classic implementation
func BubbleSort(list []int) {
	for i := 0; i < cap(list); i++ {
		for j := 0; j+1 < cap(list)-i; j++ {
			if list[j] > list[j+1] {
				Swap(list, j, j+1)
			}
		}
	}
}

func InsertionSort(list []int) {

}

func QuickSort(list []int) {
	// pivot
}

// MergeSort wrapper function
func MergeSort(list []int) []int {
	return SortSlice(list)
}

// MergeSorted takes two sorted slices and combines them into one sorted slice
func MergeSorted(a, b []int) []int {
	var sorted = make([]int, len(a)+len(b))
	insertIdx, aIdx, bIdx := 0, 0, 0

	for aIdx < len(a) && bIdx < len(b) {
		if a[aIdx] < b[bIdx] {
			sorted[insertIdx] = a[aIdx]
			aIdx++
		} else {
			sorted[insertIdx] = b[bIdx]
			bIdx++
		}
		insertIdx++
	}

	for aIdx < len(a) {
		sorted[insertIdx] = a[aIdx]
		aIdx++
		insertIdx++
	}

	for bIdx < len(b) {
		sorted[insertIdx] = b[bIdx]
		bIdx++
		insertIdx++
	}
	return sorted
}

// SortSlice takes a slice and sorts it using a divide and conquer approach
func SortSlice(list []int) []int {
	if len(list) == 1 {
		return list
	}
	if len(list) == 2 {
		if list[0] > list[1] {
			Swap(list, 0, 1)
		}
		return list
	}
	split := len(list) / 2
	left := list[:split]
	right := list[split:]
	return MergeSorted(SortSlice(left), SortSlice(right))
}
