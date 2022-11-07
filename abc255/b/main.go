package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, k int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &k)
	lighter := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(r, &lighter[i])
	}
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			fmt.Fscan(r, &points[i][j])
		}
	}
	ans := float64(0)
	for i := 0; i < n; i++ {
		near_dist := float64(10e10)
		for j := 0; j < k; j++ {
			near_dist = min(near_dist, dist(points[i], points[lighter[j]-1]))
		}
		ans = max(ans, near_dist)
	}
	fmt.Println(ans)
}

func min(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}
func max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func dist(a, b []int) float64 {
	dx := math.Abs(float64(a[0]) - float64(b[0]))
	dy := math.Abs(float64(a[1]) - float64(b[1]))
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
