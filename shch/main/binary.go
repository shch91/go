package main

import "fmt"

//二进制距离
func binaryGap(n int) int {
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

func main() {
	fmt.Println(binaryGap(22))
}
