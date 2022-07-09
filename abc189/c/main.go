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
func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans := 0
	for l := 0; l < n; l++ {
		min_orange := a[l]
		for r := l; r < n; r++ {
			min_orange = min(min_orange, a[r])
			ans = max(ans, min_orange*(r-l+1))
		}
	}
	fmt.Println(ans)
}
