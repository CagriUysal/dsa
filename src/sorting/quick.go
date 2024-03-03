package sorting

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	q := partition(arr)
	QuickSort(arr[:q])
	QuickSort(arr[q+1:])
	return arr
}

// retuns the partition index
func partition(arr []int) int {
	last := len(arr) - 1
	pivot := arr[last]

	i := 0 // where the pivot goes
	for j := 0; j < last; j += 1 {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i += 1
		}
	}

	arr[i], arr[last] = arr[last], arr[i]
	return i
}
