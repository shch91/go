package main

import (
	"fmt"
	"math"
)

//重复的DNA子串
func findRepeatedDnaSequences(s string) []string {
	var l, n, a = 10, len(s), 4
	if n < l {
		return []string{}
	}

	var al = math.Pow(4, 10)
	var mapToInt = map[string]int{"A": 0, "C": 1, "G": 2, "T": 3}
	var nums = make([]int, n)
	for index, v := range s {
		nums[index] = mapToInt[string(v)]
	}

	var h = 0
	seen, output := make(map[int]struct{}), make(map[string]struct{})

	for start := 0; start < n-l+1; start++ {
		if start != 0 {
			h = h*a - nums[start-1]*int(al) + nums[start+l-1]
		} else {
			for i := 0; i < l; i++ {
				h = h*a + nums[i]
			}
		}

		if _, ok := seen[h]; ok {
			output[s[start:start+l]] = struct{}{}
		}
		seen[h] = struct{}{}
	}

	var ret []string
	for k := range output {
		ret = append(ret, k)
	}
	return ret
}

//寻找旋转数组中最小元素
func findMin(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	left, right := 0, l-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

//分成可能的两组
func possibleBipartition(N int, dislikes [][]int) bool {
	var graph = make([][]int, N+1)
	for _, no := range dislikes {
		graph[no[0]] = append(graph[no[0]], no[1])
		graph[no[1]] = append(graph[no[1]], no[0])
	}
	var color = make(map[int]bool, N)
	for i := 1; i <= N; i++ {
		if _, ok := color[i]; !ok && !dfsVisit(i, true, graph, color) {
			return false
		}
	}
	return true
}

//邻接矩阵的dfs遍历并染色
func dfsVisit(node int, red bool, graph [][]int, color map[int]bool) bool {
	if val, ok := color[node]; ok {
		return val == red
	}
	//染色
	color[node] = red
	for _, nei := range graph[node] {
		if !dfsVisit(nei, !red, graph, color) {
			return false
		}
	}
	return true
}

//最小交换次数
func minSwap(A []int, B []int) int {
	l := len(A)
	var dp = make([][2]int, l)
	//表示第i位置交换，不交换下最小交换次数
	dp[0][0], dp[0][1] = 0, 1

	for i := 1; i < l; i++ {
		if A[i-1] < A[i] && B[i-1] < B[i] {
			if A[i-1] < B[i] && B[i-1] < A[i] {
				dp[i][0] = dp[i-1][0]
				if dp[i-1][0] > dp[i-1][1] {
					dp[i][0] = dp[i-1][1]
				}
				dp[i][1] = dp[i][0] + 1

			} else {
				dp[i][0] = dp[i-1][0]     //当前位置不交换，上一个位置也不交换
				dp[i][1] = dp[i-1][1] + 1 //当前位置交换，上一个位置也必须交换
			}
		} else {
			dp[i][0] = dp[i-1][1]     //当前位置不交换，上一个位置必须交换
			dp[i][1] = dp[i-1][0] + 1 //当前位置交换，上一个位置不交换
		}
	}
	if dp[l-1][0] > dp[l-1][1] {
		return dp[l-1][1]
	}
	return dp[l-1][0]
}

//三维形体投影面积
func projectionArea(grid [][]int) int {
	return 0
}

//单调数列
func isMonotonic(A []int) bool {
	var l = len(A)
	if l <= 1 {
		return true
	}
	//递增标记
	var flag = 0
	if A[1] > A[0] {
		flag = 1
	}
	if A[1] < A[0] {
		flag = -1
	}

	for i := 2; i < l; i++ {
		//增
		if flag > 0 && A[i] < A[i-1] {
			return false
		} else if flag < 0 && A[i] > A[i-1] {
			return false
		} else {
			if A[i] > A[i-1] {
				flag = 1
			} else if A[i] < A[i-1] {
				flag = -1
			}
		}
	}
	return true
}

func minOperations(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var op = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			op += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return op
}

func maxAscendingSum(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	var dp = make([]int, len(nums))
	var ret = nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

//平均等待时间
func averageWaitingTime(customers [][]int) float64 {
	var now,wait = 0,0
	for _,arr := range customers {
		if arr[0]>now{
			now=arr[0]
		}
		wait+=arr[1]+now-arr[0]
		now+=arr[1]
	}
	return float64(wait*1.0)/float64(len(customers))
}

func main() {
	fmt.Println(averageWaitingTime([][]int{{5,2},{5,4},{10,3},{20,1}}))
	var arr = make([]int, 4)
	fmt.Println(arr)

	fmt.Printf("%+v \n", possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 3}}))

	fmt.Printf("%+v \n", findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
