package main

//猜数字 二分查找
func guessNumber(n int) int {
	var left, right = 1, n
	var mid = left + (right-left)/2
	var g = guess(mid)
	for g != 0 {
		if g > 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = left + (right-left)/2
		g = guess(mid)
	}
	return mid
}


func guess(num int) int {
	return 0
}

