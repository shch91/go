package main

import (
	"fmt"
	"math"
)

//水缸最少操作次数
func storeWater(bucket []int, vat []int) int {
	//最大的水缸
	var maxK, ans = vat[0], math.MaxInt32
	for _, val := range vat {
		if val > maxK {
			maxK = val
		}
	}
	if maxK == 0 {
		return 0
	}

	//循环加水次数
	for add := 1; add <= maxK; add++ {
		cur := add
		for i := 0; i < len(bucket); i++ {
			//固定加水次数 对于水缸应使用的最小的水桶
			least := vat[i] / add
			if vat[i]%add != 0 {
				least++
			}
			//升级水桶
			if least > bucket[i] {
				cur += least - bucket[i]
			}
		}
		if cur < ans {
			ans = cur
		}
	}
	return ans
}

//最小元素
func min(args ...int) int {
	var m = args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

//最大元素
func max(args ...int) int {
	var m = args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

//乐队站位
func orchestraLayout(num int, xPos int, yPos int) int {
	//所在第几圈
	n := min(xPos, yPos, num-xPos-1, num-yPos-1)
	//前n-1圈元素总数
	kinds := 4*n*num - 4*n - 4*n*(n-1)
	//判断所在的边
	if n == xPos {
		kinds += yPos - n + 1
	} else if n == yPos {
		kinds += (num-2*n)*3 - 3 + num - n - xPos
	} else if n == num-xPos-1 {
		kinds += (num-2*n)*2 - 2 + num - n - yPos
	} else if n == num-yPos-1 {
		kinds += (num-2*n)*1 - 1 + xPos - n + 1
	}

	if kinds%9 == 0 {
		return 9
	} else {
		return kinds % 9
	}
}

func main() {
	fmt.Println(orchestraLayout(4, 1, 2))
}
