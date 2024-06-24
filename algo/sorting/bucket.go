package sorting

import "math"

// assume uniformly distributed in the range [0, 1)
func BucketSort(arr []float64) []float64 {
	var buckets [][]float64
	for i := 0; i < len(arr); i += 1 {
		buckets = append(buckets, []float64{})
	}

	for _, v := range arr {
		bucketIndex := int(math.Floor(float64(len(arr)) * v))
		buckets[bucketIndex] = append(buckets[bucketIndex], v)
	}

	for _, bucket := range buckets {
		// can use any sorting algorithm here
		insertionSort(bucket)
	}

	var result []float64
	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}

func insertionSort(arr []float64) []float64 {
	for i := 1; i < len(arr); i += 1 {
		for j := i; j > 0 && arr[j] < arr[j-1]; j -= 1 {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}
