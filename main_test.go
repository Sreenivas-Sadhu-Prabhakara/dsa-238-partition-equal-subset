package main

import "testing"

func TestCanPartition(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true}, // {1,5,5} and {11}
		{[]int{1, 2, 3, 5}, false},
		{[]int{2, 2}, true},
		{[]int{1}, false},
		{[]int{}, true}, // two empty subsets, both sum 0
		{[]int{3, 3, 3, 3}, true},
		{[]int{1, 1, 1}, false}, // odd total
		{[]int{100, 100}, true},
	}
	for _, c := range cases {
		if got := CanPartition(c.in); got != c.want {
			t.Errorf("CanPartition(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
