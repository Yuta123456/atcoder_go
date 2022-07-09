package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	var n int
	fmt.Scan(&n)
	taska := make([]int, n)
	taskb := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&taska[i], &taskb[i])
	}
	ans := int(10e9)
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			if a == b {
				ans = min(ans, taska[a]+taskb[b])
			} else {
				ans = min(ans, max(taska[a], taskb[b]))
			}
			// fmt.Println(a, b, ans)
		}
	}
	fmt.Println(ans)
}
