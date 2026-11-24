package main

import "fmt"

// CanPartition reports whether the array splits into two subsets of equal
// sum: possible iff some subset reaches total/2, with an odd total failing
// immediately. Boolean subset-sum DP, swept backwards for single use.
func CanPartition(nums []int) bool {
	total := 0
	for _, v := range nums {
		total += v
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range nums {
		for s := target; s >= v; s-- {
			if dp[s-v] {
				dp[s] = true
			}
		}
		if dp[target] {
			return true // early exit once the half-sum is reachable
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(CanPartition([]int{1, 5, 11, 5}))
	fmt.Println(CanPartition([]int{1, 2, 3, 5}))
}
