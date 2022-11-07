package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	candidates := []int{}
	m_v := amin(a)
	for div := 1; div <= int(math.Sqrt(float64(m_v))); div++ {
		if m_v%div == 0 {
			candidates = append(candidates, div, m_v/div)
		}
	}
	ans := 1000000000
	can := false
	for t := 0; t < len(candidates); t++ {
		cnt := 0
		for i := 0; i < n; i++ {
			if a[i]%candidates[t] != 0 {
				break
			}
			q := a[i] / candidates[t]
			for div := 2; div < 4; div++ {
				for q%div == 0 {
					q = q / div
					cnt += 1
				}
			}
			if q != 1 {
				break
			}
			if i == n-1 {
				can = true
				// fmt.Println(ans, cnt, candidates[t])
				ans = min(ans, cnt)
			}
		}
	}
	if can {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
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

func all_the_same(array []int) bool {
	first := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] != first {
			return false
		}
	}
	return true
}

func amin(array []int) int {
	res := array[0]
	for i := 0; i < len(array); i++ {
		res = min(res, array[i])
	}
	return res
}
