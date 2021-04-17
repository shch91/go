package main

import "fmt"

//单调栈递减栈
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for v := range nums {
		result[v] = -1
	}
	var stack []int
	for i := 0; i < 2*n-1; i++ {
		//循环栈顶元素小于当前元数
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			result[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return result
}



func main() {

   for _,v:=range nextGreaterElements([]int{432,4,5,6,7}){
	   fmt.Println(v)
   }
}
