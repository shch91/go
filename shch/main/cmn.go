package main

import (
	"fmt"
	"math"
)

var (
	result [][]int
	sel    []int
)

//选择元素cmn arr原数组,sel已选择小标
func cmn(arr, sel []int, m, n int) {
	if n == 0 {
		var val = make([]int, len(sel))
		copy(val, sel)
		result = append(result, val)
		return
	}
	//下标选择
	for i := n; i <= m; i++ {
		sel = append(sel, arr[i-1])
		//后续再选n-1个
		cmn(arr, sel, i-1, n-1)
		//选择完后清空
		sel = sel[:len(sel)-1]
	}
}

var (
	temp []int
	ans  [][]int
)

//递增子序列
func findSubsequences(nums []int) [][]int {
	ans = [][]int{}
	dfsSel(nums, 0, math.MinInt32)
	return ans
}

func dfsSel(arr []int, cur, last int) {
	//递归结束
	if cur == len(arr) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}

	if arr[cur] >= last {
		temp = append(temp, arr[cur])
		dfsSel(arr, cur+1, arr[cur])
		temp = temp[:len(temp)-1]
	}
	//不相同则可选
	if arr[cur] != last {
		dfsSel(arr, cur+1, last)
	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cmn(arr, sel, 10, 2)
	for _, v := range result {
		for _, a := range v {
			fmt.Printf("%d \t", a)
		}
		fmt.Println()
	}
	fmt.Println()
	findSubsequences([]int{4, 3, 2})
	for _, v := range ans {
		for _, a := range v {
			fmt.Printf("%d \t", a)
		}
		fmt.Println()
	}

}
