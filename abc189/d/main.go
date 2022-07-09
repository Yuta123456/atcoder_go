package main

import (
	"fmt"
)

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Key struct {
	b bool
	i int
}

func solve(before bool, s []string, index int, memo map[Key]int) int {
	k := Key{before, index}
	if v, ok := memo[k]; ok {
		return v
	}
	if index == -1 {
		return 1
	}
	if s[index] == "OR" {
		if before {
			value := 2*solve(true, s, index-1, memo) + solve(false, s, index-1, memo)
			memo[k] = value
			return value
		} else {
			value := solve(false, s, index-1, memo)
			memo[k] = value
			return value
		}
	}
	if before {
		value := solve(true, s, index-1, memo)
		memo[k] = value
		return solve(true, s, index-1, memo)
	} else {
		value := solve(true, s, index-1, memo) + 2*solve(false, s, index-1, memo)
		memo[k] = value
		return value
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	memo := map[Key]int{}
	fmt.Println(solve(true, s, n-1, memo))
}
