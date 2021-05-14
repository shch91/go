package main

import "math"

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
func main() {

}
