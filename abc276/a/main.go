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
	ans := -1
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			ans = i + 1
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
