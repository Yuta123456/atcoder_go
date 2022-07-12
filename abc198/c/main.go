package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var r, x, y int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &r, &x, &y)
	dist := math.Sqrt(float64(x*x + y*y))
	// fmt.Println(dist)
	if isInt(dist) && int(dist)%r == 0 {
		fmt.Println(int(dist) / r)
	} else {
		cnt := int(dist) / r
		if cnt <= 0 {
			cnt += 1
		}
		fmt.Println(cnt + 1)
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
func isInt(x float64) bool {
	y := float64(int(x))
	return math.Abs(y-x) < 10e-18
}
