package algorithm

func QuickSort[T comparable](data []T, compareFunc func(T, T) int) {
	if len(data) <= 1 {
		return
	}
	// partisi
	pivotIdx := len(data) - 1
	pivot := data[pivotIdx]

	p := -1

	for q := 0; q < pivotIdx; q++ {
		if compareFunc(data[q], pivot) <= 0 {
			p++
			data[p], data[q] = data[q], data[p]
		}
	}

	data[p+1], data[pivotIdx] = data[pivotIdx], data[p+1]

	QuickSort(data[0:p+1], compareFunc)
	QuickSort(data[p+1:], compareFunc)
}
