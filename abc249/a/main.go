package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d, e, f, x int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &a, &b, &c, &d, &e, &f, &x)
	t_c := x / (a + c)
	a_c := x / (d + f)
	takahashi := t_c*a*b + min(a, x-t_c*(a+c))*b
	aoki := a_c*d*e + min(d, x-a_c*(d+f))*e
	// fmt.Println(takahashi, aoki)
	if takahashi > aoki {
		fmt.Println("Takahashi")
	} else if takahashi < aoki {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Draw")
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
