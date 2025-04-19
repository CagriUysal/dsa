package greedy

// assumes finishTimes are sorted
func RecursiveActivitySelection(startTimes []int, finishTimes []int) []int {
	n := len(startTimes)
	result := []int{}
	if n > 0 {
		result = append(result, 0)
	}

	return solveRecursive(startTimes, finishTimes, 0, n, result)
}

func solveRecursive(startTimes []int, finishTimes []int, k int, n int, result []int) []int {
	m := k + 1
	for m < n && startTimes[m] < finishTimes[k] {
		m += 1
	}

	result = append(result, m)
	if m < n {
		return solveRecursive(startTimes, finishTimes, m, n, result)
	}

	return result
}

// assumes finishTimes are sorted
func ActivitySelection(startTimes []int, finishTimes []int) []int {
	i := 0
	result := []int{i}

	for j := i + 1; j < len(startTimes); j++ {
		if startTimes[j] < finishTimes[i] {
			continue
		}
		i = j
		result = append(result, j)
	}

	return result
}
