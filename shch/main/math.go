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

//自除数
func selfDividingNumbers(left int, right int) []int {
	var ret []int
	for k := left; k <= right; k++ {
		c := k
		for c > 0 {
			t := c % 10
			if t == 0 {
				break
			}
			if k%t != 0 {
				break
			}
			c /= 10
		}
		if c == 0 {
			ret = append(ret, k)
		}
	}
	return ret
}

//判定二的幂函数
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}

//计算二进制中1的个数
func count1(val int) int {
	var res = 0
	for val != 0 {
		val = val & (val - 1)
		res++
	}
	return res
}

//蚂蚁掉落时刻
func getLastMoment(n int, left []int, right []int) int {
	var max = math.MinInt32
	for _, val := range left {
		if val > max {
			max = val
		}
	}
	for _, val := range right {
		if n-val > max {
			max = n - val
		}
	}
	return max
}

//基数数量
func countOdds(low int, high int) int {

	if low%2 == 0 {
		low += 1
	}
	if high%2 == 0 {
		high -= 1
	}

	return (high-low)/2 + 1
}

func main() {
	ret := selfDividingNumbers(1, 22)
	fmt.Println(ret)
	fmt.Println(singleNumber([]int{-1, -1, -1, -2}))
	var pos = 1
	fmt.Println(^pos)
	//fmt.Println(Fact(10))
}
