package main

import (
	"bufio"
	"fmt"
	"os"
)

func contains(s string, c int) bool {
	for i := 0; i < len(s); i++ {
		if int(s[i]) == c {
			return true
		}
	}
	return false
}
func solve(array []string, k int) int {
	// a ~ z
	res := 0
	for i := 97; i < 123; i++ {
		cnt := 0
		for j := 0; j < len(array); j++ {
			if contains(array[j], i) {
				cnt += 1
			}
		}
		if cnt == k {
			res += 1
		}
	}
	return res
}
func main() {
	var n, k int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &k)
	s_array := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s_array[i])
	}
	ans := 0
	for i := 0; i < (1 << n); i++ {
		select_array := make([]string, 0)
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				select_array = append(select_array, s_array[j])
			}
		}
		cnt := solve(select_array, k)
		ans = max(ans, cnt)
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
