package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &k)
	operation := make([][]int, n)
	for i := 0; i < n; i++ {
		operation[i] = make([]int, 2)
		fmt.Fscan(r, &operation[i][0], &operation[i][1])
	}
	// 一回も操作を無視することなくやってみる
	// もし置き換える方が値が大きくなる操作があれば、
	// 「それ以前の操作」は無視する必要がない
	// 無視したほうが大きくなる操作は、無視したほうが良い。
	// どこを無視するかさえ決めれば、最大値は一意に定まる
	// 無視したほうが良い点は二つ
	// 今より小さくなる置き換えと、負の数の足し算。
	//
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
