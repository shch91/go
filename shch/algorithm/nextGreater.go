package main

import (
	"fmt"
)

//单调栈
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for v := range nums {
		result[v] = -1
	}
	var stack = []int{}
	for i := 0; i < 2*n-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			result[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return result
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func main() {

	in := make(chan int)
	quitChan := make(chan bool)
	quitChan2 := make(chan bool)

	value := 0

	go func() {
		for i := 0; i < 3; i++ {

			value = value + 1
			in <- value

			fmt.Println("write finish, value ", value)

			//time.Sleep(time.Second)
		}
		quitChan <- true
	}()
	go func() {
		for {
			select {
			case v := <-in:
				fmt.Println("read finish, value ", v)
			case <-quitChan:
				quitChan2 <- true
				return
			}
		}

	}()

	<-quitChan2
	fmt.Println("task is done ")


	/*data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]

	fmt.Println(s)
	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。

	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。*/

}
