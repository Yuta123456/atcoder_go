package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s, t string
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &s)
	fmt.Fscan(r, &t)
	if len(s) < len(t) {
		fmt.Println("No")
		return
	}

	for start := 0; start < len(s); start++ {
		flag := true
		for i := 0; i < len(t); i++ {
			if start+i >= len(s) {
				flag = false
				break
			}
			if s[start+i] != t[i] {
				flag = false
				break
			}
			if start+i == len(s)-1 {
				break
			}
		}
		if flag {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
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
