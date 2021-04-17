package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
	"unsafe"
)



type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
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

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if reflect.DeepEqual(got, want) {
		fmt.Println("got == want")
	}
	fmt.Println(unsafe.Sizeof(float64(0)))

	for a := 1; a < 10; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%d * %d =%d  ", a, b, a*b)
		}
		fmt.Println()
	}

	fmt.Println(Fact(10))
}
