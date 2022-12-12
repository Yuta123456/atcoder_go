package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var h, w int
	var s, t string
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &h, &w)
	S := make([]string, h)
	T := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(r, &s)
		S[i] = s
	}
	for i := 0; i < h; i++ {
		fmt.Fscan(r, &t)
		T[i] = t
	}
	map_s := map[string]int{}
	map_t := map[string]int{}
	for i := 0; i < w; i++ {
		var bs string
		for j := 0; j < h; j++ {
			bs = strings.Join(S[j][i], bs)
		}
		if _, ok := map_s[bs]; ok {
			map_s[bs] += 1
		} else {
			map_s[bs] = 1
		}
	}
	for i := 0; i < w; i++ {
		var bt string
		for j := 0; j < h; j++ {
			bt += T[j][i]
		}
		if _, ok := map_t[bt]; ok {
			map_t[bt] += 1
		} else {
			map_t[bt] = 1
		}
	}
	for key, value := range map_s {
		if v, ok := map_t[key]; !(ok && v == value) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

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
