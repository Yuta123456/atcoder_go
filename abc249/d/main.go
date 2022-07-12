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
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := cnt[a[i]]; ok {
			cnt[a[i]] += 1
		} else {
			cnt[a[i]] = 1
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 1; j < int(math.Sqrt(float64(a[i]))+1); j++ {
			if a[i]%j == 0 {
				div := a[i] / j
				c, _ := cnt[div]
				d, _ := cnt[j]
				// fmt.Println(a[i], j, "が", d, "  ", div, "が", c, c*d*2)
				if div == j {
					ans += c * d
				} else {
					ans += c * d * 2
				}
			}
		}
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
