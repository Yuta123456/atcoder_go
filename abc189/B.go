package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scanf("%d %d\n", &n, &x)
	alc := make([][]int, n)
	for i := 0; i < n; i++ {
		var p, v int
		fmt.Scanf("%d %d\n", &p, &v)
		alc[i] = make([]int, 2)
		alc[i][0] = p
		alc[i][1] = v
	}
	take := 0
	for i, val := range alc {
		take += val[0] * val[1]
		if take > x*100 {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}
