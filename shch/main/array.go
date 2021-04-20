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

func main() {
	fmt.Printf("%+v", findRepeatedDnaSequences(""))

	fmt.Printf("%+v", findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
