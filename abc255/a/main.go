package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var R int
	var C int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &R)
	fmt.Fscan(r, &C)
	array := make([][]int, 2)
	for i := 0; i < 2; i++ {
		array[i] = make([]int, 2)
		fmt.Fscan(r, &array[i][0], &array[i][1])
	}
	fmt.Println(array[R-1][C-1])

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
