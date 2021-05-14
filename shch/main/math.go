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

//12345678910 找到第n位数字
func findNthDigit(n int) int {
	//当前位数的最高值，最低值，以及值的宽度
	var low, high uint64 = 1, 9
	var width, mul uint64 = 1, 10
	//位置
	var pos uint64 = 0
	for {
		pos += (high - low + 1) * width
		if pos > uint64(n) {
			break
		}
		low = high + 1
		high = low*mul - 1
		width++
	}
	//退回位数
	//往回退的步数
	back := pos - uint64(n)
	remain := back % width
	high -= back / width
	for remain != 0 {
		high /= 10
		remain--
	}
	//取当前末尾数字
	return int(high % 10)
}

//排列硬币
func arrangeCoins(n int) int {
	var row = 0
	for n > row {
		row++
		n -= row
	}
	return row
}

//移除k个元素剩余最小元素
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var stack = []byte{num[0]}

	for i := 1; i < len(num); i++ {
		l := len(stack)
		if l < 1 { //空
			stack = append(stack, num[i])
		} else {
			//大于当前元素的弹出
			for l > 0 && k > 0 && num[i] < stack[l-1] {
				stack = stack[:l-1]
				k--
				l--
			}
			stack = append(stack, num[i])
		}
	}
	for k > 0 {
		k--
		l := len(stack)
		stack = stack[:l-1]
	}
	//去掉0
	for i := 0; i < len(stack); i++ {
		if stack[i] != '0' {
			return string(stack[i:])
		}
	}
	return "0"
}

//各位数字相加只剩一位
func addDigits(num int) int {
	for num > 9 {
		val := 0
		for num != 0 {
			val += num % 10
			num /= 10
		}
		num = val
	}
	return num
}



func main() {
	fmt.Println(removeKdigits("112", 1))
	fmt.Println(arrangeCoins(5))
	fmt.Println(findNthDigit(1))
	ret := selfDividingNumbers(1, 22)
	fmt.Println(ret)
	fmt.Println(singleNumber([]int{-1, -1, -1, -2}))
	var pos = 1
	fmt.Println(^pos)
	//fmt.Println(Fact(10))
}
