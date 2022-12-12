package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, t int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &t)
	a := make([]int, n)
	acc := make([]int, n+1)
	a_sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
		a_sum += a[i]
		acc[i+1] = a_sum
	}
	t = t % a_sum
	// fmt.Println(t, acc)
	for i := 0; i < n+1; i++ {
		if acc[i] > t {
			fmt.Println(i, t-acc[i-1])
			return
		}
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
