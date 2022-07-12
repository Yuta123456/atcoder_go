package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m, h int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &m, &h)
	if h%m == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
