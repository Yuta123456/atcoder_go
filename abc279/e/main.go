package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, k int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &k, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
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
