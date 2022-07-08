package main

import (
	"fmt"
)

func main() {
	var n, s, d int
	fmt.Scanf("%d %d %d", &n, &s, &d)
	magic := make([][]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		magic[i] = make([]int, 2)
		fmt.Scanf("%d %d", &x, &y)
		magic[i][0] = x
		magic[i][1] = y
	}
	// fmt.Println(magic)
	for i := 0; i < n; i++ {
		if magic[i][0] < s && magic[i][1] > d {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
