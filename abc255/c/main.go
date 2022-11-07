package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var x, a, d, n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &x, &a, &d, &n)
	last := a + d*(n-1)
	// 初項 <= 末項の場合
	if a <= last {
		// 初項 > X or 末項 < X の場合
		if d == 0 {
			fmt.Println(int(math.Abs(float64(a - x))))
			return
		}
		if a >= x {
			fmt.Println(a - x)
			return
		}

		if last <= x {
			fmt.Println(x - last)
			return
		}

		shift_a := x - a
		mod_d := ((shift_a % d) + d) % d
		fmt.Println(min(d-mod_d, mod_d))
	} else {
		os.Exit(1)
		// 初項 > X or 末項 < X の場合
		if last > x || d == 0 {
			fmt.Println(last - x)
			return
		}

		if a < x {
			fmt.Println(x - a)
			return
		}

		shift_a := x - a
		mod_d := shift_a % d
		fmt.Println(min(d-mod_d, mod_d))
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
