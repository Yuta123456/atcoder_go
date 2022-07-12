package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &s)
	lower_s := strings.ToLower(s)
	upper_s := strings.ToUpper(s)
	s_set := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		k := s[i]
		s_set[k] = true
	}
	if len(s_set) == len(s) &&
		lower_s != s &&
		upper_s != s {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
