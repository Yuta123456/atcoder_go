package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c, d := float64(a), float64(b)
	fmt.Println((c - d) * 100 / c)
}
