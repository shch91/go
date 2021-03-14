package main

import (
	"errors"
	"fmt"
	"math"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func levelOrder(node *TreeNode)  {
	if node == nil{
		return
	}

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

func main() {
	for a := 1; a < 10; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%d * %d =%d  ", b, a, a*b)
		}
		fmt.Println()
	}

	var result, err = Sqrt(3213)
	if err == nil {
		fmt.Println(result)
	}
	fmt.Println(Fact(10))
}
