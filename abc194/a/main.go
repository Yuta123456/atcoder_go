package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a+b >= 15 && b >= 8 {
		fmt.Println(1)
	} else if a+b >= 10 && b >= 3 {
		fmt.Println(2)
	} else if b+a >= 3 {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}
