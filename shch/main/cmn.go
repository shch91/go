package main

import (
	"fmt"
	"math"
	"sort"
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
		//去掉最后选择的那个元素
		sel = sel[:len(sel)-1]
	}
}

var (
	temp []int
	ans  [][]int
)

//递增子序列,nums递增
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

//部件采购
func purchasePlans(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	var ans = 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			ans += right - left
			left++

		}
		ans %= 1e9 + 7
	}
	return ans % (1e9 + 7)
}

//方格画
func paintingPlan(n int, k int) int {
	if k == 0 || k == n*n {
		return 1
	}
	var ret = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i*n+j*n-i*j == k {
				ret += calCmn(n, i) * calCmn(n, j)
			}
		}
	}
	return ret
}

func calCmn(m, n int) int {
	if n == 0 || m == n {
		return 1
	}
	return com(m) / com(n) / com(m-n)
}

func com(n int) int {
	var ret = 1
	for n > 0 {
		ret *= n
		n--
	}
	return ret
}

func main() {

	nums := []int{2, 2, 1, 9}
	purchasePlans(nums, 10)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cmn(arr, sel, 10, 8)
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
