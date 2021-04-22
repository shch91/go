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
	color[node]=red
	for _,nei := range graph[node] {
		if !dfsVisit(nei, !red, graph, color) {
			return false
		}
	}
	return true
}

//最小交换次数
func minSwap(A []int, B []int) int {

}

func main() {


	//but output: 2 2

	fmt.Printf("%+v \n", possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 3}}))

	fmt.Printf("%+v \n", findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
