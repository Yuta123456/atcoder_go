package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s[i])
	}
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &t[i])
	}
	// fmt.Println(t, s)
	ans := make([]int, n)
	start := -1
	startValue := int(10e10)
	for i := 0; i < n; i++ {
		if t[i] < startValue {
			startValue = t[i]
			start = i
		}
	}
	ans[start] = t[start]
	for index := 1; index < n; index++ {
		i := (start + index) % n
		pre_i := (i + n - 1) % n
		ans[i] = min(ans[pre_i]+s[pre_i], t[i])
	}
	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
