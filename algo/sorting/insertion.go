package sorting

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i += 1 {
		for j := i; j > 0 && arr[j] < arr[j-1]; j -= 1 {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}
