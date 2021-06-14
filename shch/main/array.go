package main

import (
	"fmt"
	"math"
	"shch/main/util"
	"sort"
	"strconv"
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
	var ans, prev = 0, -1
	var l = len(flowerbed)
	for i := 0; i < l; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				ans += i / 2
			} else {
				ans += (i - prev - 2) / 2
			}
			prev = i
		}
	}
	if prev < 0 {
		ans += (l + 1) / 2
	} else {
		ans += (l - prev - 1) / 2
	}
	return ans >= n
}

func maximumProduct(nums []int) int {
	var min1, min2 = math.MaxInt32, math.MaxInt32
	var max1, max2, max3 = math.MinInt32, math.MinInt32, math.MinInt32
	for _, val := range nums {
		//最小的两个数
		if val < min1 {
			min2 = min1
			min1 = val
		} else if val < min2 {
			min2 = val
		}

		//最大的是三个数
		if val > max1 {
			max3 = max2
			max2 = max1
			max1 = val
		} else if val > max2 {
			max3 = max2
			max2 = val
		} else if val > max3 {
			max3 = val
		}
	}
	var a, b = max1 * max2 * max3, max1 * min1 * min2
	if a > b {
		return a
	}
	return b
}

func findMaxAverage(nums []int, k int) float64 {
	var sum = 0
	var result float64
	for t := 0; t < k; t++ {
		sum += nums[t]
	}
	result = float64(sum) / float64(k)
	for t := k; t < len(nums); t++ {
		sum += nums[t] - nums[t-k]
		avg := float64(sum) / float64(k)
		if avg > result {
			result = avg
		}
	}
	return result
}

//找重复与丢失的数字
func findErrorNums(nums []int) []int {
	var l = len(nums)
	var xor, xor0, xor1 = 0, 0, 0
	for i := 1; i <= l; i++ {
		xor ^= i ^ nums[i-1]
	}
	//最小比特位置
	var val = xor & ^(xor - 1)
	for i := 1; i <= l; i++ {
		if nums[i-1]&val != 0 {
			xor1 ^= nums[i-1]
		} else {
			xor0 ^= nums[i-1]
		}
		if i&val != 0 {
			xor1 ^= i
		} else {
			xor0 ^= i
		}
	}
	for i := 0; i < l; i++ {
		if nums[i] == xor0 {
			return []int{xor0, xor1}
		}

	}
	return []int{xor1, xor0}
}

//平滑图片
func imageSmoother(img [][]int) [][]int {
	var ans [][]int
	var r, l = len(img), len(img[0])
	for i := 0; i < r; i++ {
		var row []int
		for j := 0; j < l; j++ {
			sum, cnt := 0, 0
			//每个元素
			for m := -1; m <= 1; m++ {
				for n := -1; n <= 1; n++ {
					if i+m >= 0 && i+m < r && j+n >= 0 && j+n < l {
						sum += img[i+m][j+n]
						cnt++
					}
				}
			}
			row = append(row, sum/cnt)
		}
		ans = append(ans, row)
	}
	return ans
}

//非递减数列
func checkPossibility(nums []int) bool {
	var l, cnt = len(nums), 0
	for i := 0; i < l-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true

}

//最长连续递增子序列
func findLengthOfLCIS(nums []int) int {
	var ans, r = 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i+1] {
			ans++
		} else {
			ans = 1
		}
		if ans > r {
			r = ans
		}
	}
	return r
}

//数组积符号
func arraySign(nums []int) int {
	var ans = 1
	for _, v := range nums {
		if signFunc(v) == 0 {
			return 0
		} else if signFunc(v) < 0 {
			ans = -ans
		}
	}
	return ans
}

func signFunc(p int) int {
	if p > 0 {
		return 1
	} else if p < 0 {
		return -1
	}
	return 0
}

//二分查找
func search(nums []int, target int) int {
	for i, j := 0, len(nums)-1; i <= j; {
		mid := (i + j) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}

//棒球分数
func calPoints(ops []string) int {
	var score []int
	var sum = 0
	for i := 0; i < len(ops); i++ {
		if ops[i] == "+" {
			l := len(score)
			score = append(score, score[l-1]+score[l-2])
		} else if ops[i] == "D" {
			l := len(score)
			score = append(score, 2*score[l-1])
		} else if ops[i] == "C" {
			score = score[:len(score)-1]
		} else {
			v, _ := strconv.Atoi(ops[i])
			score = append(score, v)
		}
	}

	for _, v := range score {
		sum += v
	}
	return sum
}

/**
 * Definition for Employee.
 */
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

//员工重要程度
func getImportance(employees []*Employee, id int) int {
	var ans = 0
	for _, v := range employees {
		//根结点
		if v.Id == id {
			ans += v.Importance
			//下属权重
			for _, c := range v.Subordinates {
				ans += getImportance(employees, c)
			}
		}
	}
	return ans
}

type e struct {
	cnt, l, r int
}

//数组的度
func findShortestSubArray(nums []int) int {
	var m = make(map[int]e)
	var max, len = 0, math.MaxInt32
	for i, v := range nums {
		if item, ok := m[v]; ok {
			item.cnt++
			item.r = i
			m[v] = item
		} else {
			m[v] = e{1, i, i}
		}
	}
	for _, v := range m {
		if v.cnt > max {
			max = v.cnt
			len = v.r - v.l + 1
		} else if v.cnt == max {
			len = util.Min(len, v.r-v.l+1)
		}
	}
	return len
}

//图像渲染
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	dfsFlood(&image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfsFlood(img *[][]int, r, c int, color, newColor int) {
	var row, col = len(*img), len((*img)[0])
	(*img)[r][c] = newColor
	//上
	if r-1 >= 0 && (*img)[r-1][c] == color && color != newColor {
		dfsFlood(img, r-1, c, color, newColor)
	}
	//下
	if r+1 < row && (*img)[r+1][c] == color && color != newColor {
		dfsFlood(img, r+1, c, color, newColor)
	}
	//左
	if c-1 >= 0 && (*img)[r][c-1] == color && color != newColor {
		dfsFlood(img, r, c-1, color, newColor)
	}
	//右
	if c+1 < col && (*img)[r][c+1] == color && color != newColor {
		dfsFlood(img, r, c+1, color, newColor)
	}
}

//反转图像
func flipAndInvertImage(image [][]int) [][]int {
	for _, arr := range image {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			image[i][j] = image[i][j] ^ 1
		}
	}

	return image
}

func largeGroupPositions(s string) [][]int {
	var pre = 0
	var i int
	var ans [][]int
	for i = 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			continue
		}
		if i-pre >= 3 {
			ans = append(ans, []int{pre, i - 1})
		}
		pre = i
	}
	if i-pre >= 3 {
		ans = append(ans, []int{pre, i - 1})
	}
	return ans
}

func arrayRankTransform(arr []int) []int {
	var m = make(map[int]int)
	var t = RemoveReplicaSliceInt(arr)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	var ans []int
	for i := 0; i < len(t); i++ {
		m[t[i]] = i + 1
	}
	for i := 0; i < len(arr); i++ {
		ans = append(ans, m[arr[i]])
	}
	return ans
}

/*
 * slice(int类型)元素去重
 */
func RemoveReplicaSliceInt(slc []int) []int {

	var result []int
	tempMap := make(map[int]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}

func dominantIndex(nums []int) int {
	var max = math.MinInt32
	var index = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != max && max < 2*nums[i] {
			return -1
		}
	}
	return index
}

func isToeplitzMatrix(matrix [][]int) bool {
	var r, l = len(matrix), len(matrix[0])

	for i, j := r-1, 0; i >= 0 && j < l; {
		//循环对脚线元素
		for s := 1; s <= util.Min(r, l); s++ {
			if i+s < r && j+s < l  && matrix[i][j]!=matrix[i+s][j+s]{
             	return  false
			}
		}
		if i==0{
			j++
		}else{
			i--
		}
	}
	return true
}


func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(imageSmoother([][]int{{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}))
	fmt.Printf("%v", findErrorNums([]int{3, 2, 3, 4, 6, 5}))
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
