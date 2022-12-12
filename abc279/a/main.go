package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &s)
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'v' {
			ans += 1
		} else {
			ans += 2
		}
	}
	fmt.Println(ans)
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
