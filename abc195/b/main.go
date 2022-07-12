package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, w int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &a, &b, &w)
	w = w * 1000
	o_min := (w + a - 1) / a
	o_max := (w + b - 1) / b
	fmt.Println(o_min, o_max)

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
