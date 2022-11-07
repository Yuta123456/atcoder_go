package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	ans := 0
	// div_2n := make([]int, 0)
	for div := 1; div <= int(math.Sqrt(2.0*float64(n))); div++ {
		if (2*n)%div == 0 {
			a := div
			b := 2 * n / div
			if a%2 != b%2 {
				ans += 2
				if a == b {
					ans -= 1
				}
			}
		}
	}
	fmt.Println(ans)
}
