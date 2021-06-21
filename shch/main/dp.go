package main

import (
	"fmt"
	"math"
	"shch/main/util"
)

//动态规划
func minCostClimbingStairs(cost []int) int {
	var dp = make([]int, len(cost)+1)
	dp[1] = 0
	dp[0] = 0

	for i := 2; i <= len(cost); i++ {
		if dp[i-1]+cost[i-1] < dp[i-2]+cost[i-2] {
			dp[i] = dp[i-1] + cost[i-1]
		} else {
			dp[i] = dp[i-2] + cost[i-2]
		}
	}
	return dp[len(cost)]
}

/**
 *硬币找零，无限数量,求几种方式
 */
func changeWay(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for k := coin; k <= amount; k++ {
			dp[k] = dp[k-coin] + dp[k]
		}
	}
	return dp[amount]
}

/**
 * 硬币找零，最小数量
 */
func changeMin(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			remain := i - coin
			if remain >= 0 && dp[remain] != -1 {
				dp[i] = util.Min(dp[remain]+1, dp[i])
			}
		}
		//无法被找开
		if dp[i] == math.MaxInt32 {
			dp[i] = -1
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(changeWay(5, []int{1, 5, 2}))
	fmt.Println(changeMin(5,[]int{1,2,5}))
}
