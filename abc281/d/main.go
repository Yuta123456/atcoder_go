package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, d, s int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &s, &d)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	// DP
	// 前からi個みて、余りがjの時の最大値
	dp := map[int]int{}
	dp[toKey(0, 0, 0)] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < d+1; j++ {
			for k := 0; k < s+1; k++ {
				// 今回のものから、次へ動かす
				key := toKey(i, j, k)
				if v, ok := dp[key]; ok {
					// 今回のアイテムを取る場合
					selKey := toKey(i+1, (j+a[i])%d, k+1)
					dp[selKey] = max(dp[selKey], v+a[i])
					// 今回のアイテムを取らない場合
					nSelKey := toKey(i+1, j, k)
					dp[nSelKey] = max(dp[nSelKey], v)
				}
			}
		}
	}
	if _, ok := dp[toKey(n, 0, s)]; !ok {
		fmt.Println(-1)
		return
	}
	fmt.Println(dp[toKey(n, 0, s)])
}

func toKey(i, j, k int) int {
	return i*10000 + j*100 + k
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
