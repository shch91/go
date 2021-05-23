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

func main() {

	fmt.Println(findComplement(1))
	fmt.Println(binaryGap(22))
}
