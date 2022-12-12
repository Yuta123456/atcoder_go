package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b float64
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &a, &b)
	c := float64(2) / float64(3)
	x := int(math.Pow(a, c)/(math.Pow(2, c)*math.Pow(b, c)) - 1)
	ans := 10e18
	for i := -1.0; i <= 1; i++ {
		ans = min(ans, f(a, b, float64(x)+i))
	}
	fmt.Println(ans)
}

func f(a, b, x float64) float64 {
	return a/math.Sqrt(1+x) + b*x
}
func min(x, y float64) float64 {
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
