package sorting

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i += 1 {
		for j := 0; j < n-i-1; j += 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
