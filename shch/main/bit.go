package main

import "fmt"

//二进制距离，二进制表示中为1的最大的距离
func binaryGap(n int) int {

	//上次二进制位为1的位置、
	var last, bin, dis = -1, 1, 0

	for i := 0; i < 64; i++ {
		if bin < n && n&bin != 0 {
			if last >= 0 && i-last > dis {
				dis = i - last
			}
			last = i
		}
		bin <<= 1
	}
	return dis
}

//二进制补数
func findComplement(num int) int {
	var t = num
	//最高位位置
	var high, cur = 0, 1
	for t != 0 {
		if t&1 > 0 {
			high = cur
		}
		t >>= 1
		cur++
	}
	return num ^ (1<<high - 1)
}

//二进制中最低位1的位置所表示的值
func binLowOne(t int) int {
	return t &^ (t - 1)
}
//二进制中最低位1的位置,下标从1开始
func binLowPos(t int) int {
	v := binLowOne(t)
	pos := 1
	for v > 1 {
		pos++
		v >>= 1
	}
	return pos
}

//正整数二进制交替
func hasAlternatingBits(n int) bool {

	for flag, t := n&1, n>>1; t > 0; t >>= 1 {
		if t&1 == flag {
			return false
		}
		flag = flag ^ 1
	}
	return true
}

//二进制中1的个数
func bitCount(n int) int {
	count := 0
	for n != 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

func countBits(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		result = append(result, bitCount(i))
	}
	return result
}

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Printf(" %d = %d  \n", binLowOne(i),binLowPos(i))
	}
	fmt.Println(hasAlternatingBits(6))
	fmt.Println(findComplement(1))
	fmt.Println(binaryGap(22))
}
