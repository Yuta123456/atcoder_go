package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	// fmt.Println(a)
	sum := 0
	ans := 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	for i := 0; i < n; i++ {
		ans += (a[i] * a[i]) * (n - 1)
		ans -= 2 * a[i] * (sum - a[i])
		sum -= a[i]
	}
	fmt.Println(ans)

}
