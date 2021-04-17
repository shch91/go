package main

import "fmt"

//丢失的两个数
func missingTwo(nums []int) []int {
	//长度与缺少的两个数的异或结果
	var l, xor = len(nums), 0
	for _, v := range nums {
		xor ^= v
	}
	for i := 1; i <= l+2; i++ {
		xor ^= i
	}
	//求一个数组的最后一位bit表示的数
	var lastBit,one int = xor & (-xor),0

	for _, v := range nums {
		if lastBit&v>0{
			one^=v
		}
	}
	for i := 1; i <= l+2; i++ {
		 if lastBit&i>0{
		 	one^=i
		 }
	}
	return []int{one,xor^one}

}

func main() {
	fmt.Printf(" miss two %v \n", missingTwo([]int{2,3}))
}
