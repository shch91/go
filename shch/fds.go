package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return math.Sqrt(f),nil
}

func main() {
	for a := 1; a < 10; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%d * %d =%d  ", b, a, a*b)
		}
		fmt.Println()
	}
	var a int = 10
	fmt.Printf("变量的地址: %x\n", &a)
	var result,err=Sqrt(3213)
	if err==nil{
		fmt.Println(result)
	}

}
