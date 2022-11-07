package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &p[i])
	}
	m := n + 1
	index := -1
	for i := 0; i < n; i++ {
		if m < p[n-i-1] {
			index = n - i - 1
			m = p[n-i-1]
			break
		}
		m = p[n-i-1]
	}

	var ans []int
	for i := 0; i < index; i++ {
		ans = append(ans, p[i])
	}
	change := -1
	for i := index; i < n; i++ {
		if m > p[i] && change < p[i] {
			change = p[i]
		}
	}
	ans = append(ans, change)
	rem := p[index:]
	rem = remove(change, rem)
	sort.Slice(rem, func(i, j int) bool {
		return rem[i] > rem[j]
	})
	// fmt.Println(rem)
	ans = append(ans, rem...)
	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i], " ")
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

func remove(target int, array []int) []int {
	res := []int{}
	for i := 0; i < len(array); i++ {
		if array[i] != target {
			res = append(res, array[i])
		}
	}
	return res
}
