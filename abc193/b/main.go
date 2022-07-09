package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	var a, p, x, n int
	fmt.Scan(&n)
	r := bufio.NewReader(os.Stdin)
	ans := int(10e9)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a, &p, &x)
		if x > a {
			ans = min(ans, p)
		}
	}
	if ans == int(10e9) {
		fmt.Println(-1)
		return
	}
	fmt.Println(ans)
}
