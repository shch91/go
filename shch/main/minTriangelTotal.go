package main

import (
	"fmt"
)

//三角行的最下路径和
func minimumTotal(triangle [][]int) int {
	var l = len(triangle)
	var dp = make([]int, l+1)

	for i := l - 1; i >= 0; i-- {
		for j := 0; j <=i; j++ {
			if dp[j] < dp[j+1] {
				dp[j] += triangle[i][j]
			} else {
				dp[j] = dp[j+1] + triangle[i][j]
			}
		}
	}

	return dp[0]
}

func main() {
	//[[-1],[2,3],[1,-1,-3]]
	fmt.Println(minimumTotal([][]int{{-1}, {2, 3}, {1, -1, -3}}))
}
