package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, m int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	adjacent_list := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		adjacent_list[a-1] = append(adjacent_list[a-1], b-1)
		adjacent_list[b-1] = append(adjacent_list[b-1], a-1)
	}
	for i := 0; i < n; i++ {
		sort.Slice(adjacent_list[i], func(i2, j int) bool {
			return adjacent_list[i][i2] < adjacent_list[i][j]
		})
		fmt.Print(len(adjacent_list[i]), " ")
		for j := 0; j < len(adjacent_list[i]); j++ {
			fmt.Print(adjacent_list[i][j]+1, " ")
		}
		fmt.Println()
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
