package sorting

// assume `arr` only have elements in range [0,k] inclusive
func CountingSort(arr []int, k int) []int {
	count := make([]int, k+1)

	for _, v := range arr {
		count[v] += 1
	}

	for i := 1; i < len(count); i += 1 {
		count[i] += count[i-1]
	}

	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i -= 1 {
		result[count[arr[i]]-1] = arr[i]
		count[arr[i]] -= 1
	}

	return result
}
