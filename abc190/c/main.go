package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m, k int
	fmt.Scanf("%d %d", &n, &m)
	condition := make([][]int, m)
	for i := 0; i < m; i++ {
		condition[i] = make([]int, 2)
		fmt.Scan(&condition[i][0], &condition[i][1])
	}
	// fmt.Println(condition)
	fmt.Scan(&k)
	// fmt.Println(k)
	ball := make([][]int, k)
	for i := 0; i < k; i++ {
		ball[i] = make([]int, 2)
		fmt.Scan(&ball[i][0], &ball[i][1])
	}
	fmt.Println(solve(condition, ball, make([]bool, 0), n))
}

func solve(condition [][]int, ball [][]int, place []bool, n int) int {
	if len(place) == len(ball) {
		return cnt_sat_cond(condition, ball, place, n)
	}
	place = append(place, true)
	t := solve(condition, ball, place, n)
	place[len(place)-1] = false
	f := solve(condition, ball, place, n)
	return int(math.Max(float64(t), float64(f)))
}

func cnt_sat_cond(condition [][]int, ball [][]int, place []bool, n int) int {
	place_ball := make([]bool, n)
	for i := 0; i < len(place); i++ {
		if place[i] {
			place_ball[ball[i][0]-1] = true
		} else {
			place_ball[ball[i][1]-1] = true
		}
	}
	sat_cond := 0
	for i := 0; i < len(condition); i++ {
		l, r := condition[i][0]-1, condition[i][1]-1
		if place_ball[l] && place_ball[r] {
			sat_cond += 1
		}
	}
	return sat_cond
}
