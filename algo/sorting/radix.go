package sorting

// assume `d` is the maximum number of digits that `arr` elements can have
func RadixSort(arr []int, d int) []int {
	for i := 1; i <= d; i += 1 {
		arr = countingSort(arr, i)
	}

	return arr
}

// modified counting sort for radix sort
// makes the counting for `i`th least least significant digit starting from 1
func countingSort(arr []int, i int) []int {
	count := make([]int, 10)

	exp := 1
	for j := 1; j < i; j++ {
		exp *= 10
	}

	for _, v := range arr {
		digit := (v / exp) % 10
		count[digit] += 1
	}

	for i := 1; i < len(count); i += 1 {
		count[i] += count[i-1]
	}

	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i -= 1 {
		digit := (arr[i] / exp) % 10
		result[count[digit]-1] = arr[i]
		count[digit] -= 1
	}

	return result
}
