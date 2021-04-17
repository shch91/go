package main

import "fmt"

/**
 * 返回矩阵的最大子矩阵 (左上角 右下角)
 */
func getMaxMatrix(matrix [][]int) []int {
	var row, col = len(matrix), len(matrix[0])

	var result []int
	var max = ^int(^uint(0) >> 1)
	for i := 0; i < row; i++ {
		for j := i + 1; j <= row; j++ {
			arrTmp := ijSum(matrix, i, j, col)
			sum, s, e := maxSubArray(arrTmp)
			if sum > max {
				max = sum
				result = []int{i, s, j - 1, e}
			}
		}
	}
	return result
}

/**
 * 第i行与第j-1行的和
 */
func ijSum(matrix [][]int, i, j, col int) []int {
	var ret []int
	for c := 0; c < col; c++ {
		var sum = 0
		for k := i; k < j; k++ {
			sum += matrix[k][c]
		}
		ret = append(ret, sum)
	}
	return ret
}

/**
 *一维数组最大子数组和
 */
func maxSubArray(arr []int) (int, int, int) {
	var sum, max, begin, start, end = 0, ^int(^uint(0) >> 1), 0, 0, 0
	for index, v := range arr {
		sum += v
		if sum <= v {
			sum = v
			begin = index
		}

		if sum > max {
			max = sum
			start = begin
			end = index
		}
	}
	return max, start, end
}

func main() {
	sum, s, e := maxSubArray([]int{2, -2, 4, 5, -9})
	fmt.Printf("sum=%d,start=%d,end=%d\n", sum, s, e)

	fmt.Printf("%+v \n", getMaxMatrix([][]int{{9, -8, 1, 3, -2}, {-3, 7, 6, -2, 4}, {6, -4, -4, 8, -7}}))
}
