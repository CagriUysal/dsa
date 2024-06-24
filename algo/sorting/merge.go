package sorting

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

// assumes left and right arrays are sorted
func merge(left, right []int) []int {
	i, j := 0, 0
	sorted := make([]int, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i += 1
			continue
		}

		sorted = append(sorted, right[j])
		j += 1
	}

	for i < len(left) {
		sorted = append(sorted, left[i])
		i += 1
	}

	for j < len(right) {
		sorted = append(sorted, right[j])
		j += 1
	}

	return sorted
}
