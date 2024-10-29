package dynamicprogramming

import (
	"math"
)

// navie approach has O(2^n)
func RodCutNaive(prices []int, n int) int {
	if n == 0 {
		return 0
	}

	maxRevenue := math.Inf(-1)
	for i := 0; i < n; i++ {
		maxRevenue = math.Max(maxRevenue, float64(prices[i]+RodCutNaive(prices, n-i-1)))
	}

	return int(maxRevenue)
}

func MemoizedRodCut(prices []int, n int) int {
	memoizedPrices := make([]int, n+1)

	return memoizedRodCutAux(prices, n, memoizedPrices)
}

func memoizedRodCutAux(prices []int, n int, mem []int) int {
	if n == 0 {
		return 0
	}

	if mem[n] > 0 {
		return mem[n]
	}

	maxRevenue := math.Inf(-1)
	for i := 0; i < n; i++ {
		maxRevenue = math.Max(maxRevenue, float64(prices[i]+memoizedRodCutAux(prices, n-i-1, mem)))
	}

	mem[n] = int(maxRevenue)
	return mem[n]
}

func BottomUpRobCut(prices []int, n int) int {
	revenues := make([]int, n+1)
	for j := 1; j <= n; j++ {
		q := math.Inf(-1)
		for i := 1; i <= j; i++ {
			q = math.Max(q, float64(prices[i-1]+revenues[j-i]))
		}
		revenues[j] = int(q)
	}

	return revenues[n]
}
