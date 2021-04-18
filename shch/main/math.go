package main

import (
	"errors"
	"fmt"
	"math"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return math.Sqrt(f), nil
}

func Fact(k int) int {
	if k == 1 {
		return 1
	}
	return k * Fact(k-1)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//找出数组中只出现一次的数
func singleNumber(nums []int) int {
	var seenOnce, seenTwice = 0, 0
	for _, v := range nums {
		seenOnce = ^seenTwice & (seenOnce ^ v)
		seenTwice = ^seenOnce & (seenTwice ^ v)
	}
	return seenOnce
}

func main() {

	fmt.Println(singleNumber([]int{-1, -1, -1, -2}))
	var pos = 1
	fmt.Println(^pos)
	//fmt.Println(Fact(10))
}
