package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	flag := make(map[int]bool)
	for i := 2; i < int(10e6); i++ {
		if flag[i] {
			// fmt.Println(i, " skip")
			continue
		}
		for p := 2; int(math.Pow(float64(i), float64(p))) <= n; p++ {
			// fmt.Println(i, p)
			flag[int(math.Pow(float64(i), float64(p)))] = true
		}
	}
	// fmt.Println(flag)
	fmt.Println(n - len(flag))
}
