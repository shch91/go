package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
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

		if flag > 0 && A[i] < A[i-1] { //增
			return false
		} else if flag < 0 && A[i] > A[i-1] { //减
			return false
		} else { //相等
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
	var now, wait = 0, 0
	for _, arr := range customers {
		if arr[0] > now {
			now = arr[0]
		}
		wait += arr[1] + now - arr[0]
		now += arr[1]
	}
	return float64(wait*1.0) / float64(len(customers))
}

func canMakeArithmeticProgression(arr []int) bool {

	sort.Ints(arr)
	if len(arr) < 3 {
		return true
	}
	var d = arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != d {
			return false
		}
	}
	return true
}

func reorderSpaces(text string) string {
	var field = strings.Fields(text)

	var t, sl, has, avg = len(text), 0, 0, 0
	var result string
	for _, str := range field {
		sl += len(str)
	}
	if len(field) > 1 {
		avg = (t - sl) / (len(field) - 1)
	} else {
		avg = t - sl
	}

	for _, str := range field {
		result += str
		for i := 0; i < avg && has < t-sl; i++ {
			result += " "
			has++
		}

	}
	//最后剩下的空
	for has < t-sl {
		result += " "
		has++
	}
	return result
}

func numSpecial(mat [][]int) int {
	var cnt = 0
	var row, col [100]int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 && row[i] == 1 && col[j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

//格雷编码
func grayCode(n int) []int {
	var res []int
	res = append(res, 0)
	var head = 1
	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, head+res[j])
		}
		//二进制位数
		head <<= 1
	}
	return res
}

type NumArray struct {
	s []int
}

func Con(nums []int) NumArray {
	var sum = make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	return NumArray{s: sum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.s[right+1] - this.s[left]
}

func thirdMax(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	l := len(nums)
	if l < 3 {
		if l == 2 {
			return max(nums[0], nums[1])
		}
		return nums[0]
	}
	//降序排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	//第几大元素
	var count = 0
	for i := 1; i < l; i++ {
		if nums[i] != nums[i-1] {
			count++
		}
		if count == 2 {
			return nums[i]
		}
	}
	return nums[0]
}

//找到数组中未出现的数字
func findDisappearedNumbers(nums []int) []int {
	var abs = func(v int) int {
		if v > 0 {
			return v
		} else {
			return -v
		}
	}

	for _, val := range nums {
		if val > 0 {
			nums[val-1] = -abs(nums[val-1])
		} else {
			nums[-val-1] = -abs(nums[-val-1])
		}

	}
	var ans []int
	for i, _ := range nums {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	var ans = 0
	for i, j := 0, 0; i < len(g) && j < len(s); i, j = i+1, j+1 {
		for j < len(s) && g[i] > s[j] {
			j++
		}
		if j < len(s) {
			ans++
		}
	}
	return ans
}

//重塑矩阵
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	var col = len(mat[0])
	var ans [][]int
	for i := 0; i < r; i++ {
		var t []int
		for j := 0; j < c; j++ {
			index := i*c + j
			t = append(t, mat[index/col][index%col])
		}
		ans = append(ans, t)
	}
	return ans
}

//种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	var ans ,prev= 0,-1
	var l = len(flowerbed)
	for i := 0; i < l; i++ {
		if flowerbed[i] == 1 {
			if prev<0{
				ans+=i/2
			}else{
				ans+=(i-prev-2)/2
			}
			prev=i
		}
	}
	if prev<0{
		ans+=(l+1)/2
	}else{
		ans+=(l-prev-1)/2
	}
	return ans>=n
}

func main() {
	findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(thirdMax([]int{4, 7, 5, 3}))
	fmt.Println(grayCode(4))
	reorderSpaces("a b   c d")
	canMakeArithmeticProgression([]int{3, 5, 1})
	fmt.Println(averageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}))
	var arr = make([]int, 4)
	fmt.Println(arr)

	fmt.Printf("%+v \n", possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 3}}))

	fmt.Printf("%+v \n", findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
