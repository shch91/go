package util



//最小元素
func Min(args ...int) int {
	var m = args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

//最大元素
func Max(args ...int) int {
	var m = args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}
