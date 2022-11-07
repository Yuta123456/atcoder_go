package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s, t int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &s, &t)
	ans := 0
	for a := 0; a <= s; a++ {
		for b := 0; b <= s; b++ {
			for c := 0; c <= s; c++ {
				if a*b*c <= t && a+b+c <= s {
					// fmt.Println(a, b, c)
					ans += 1
				}
			}
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
