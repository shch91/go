package main

import (
	"fmt"
	"time"
)

func main()  {
	x := [3]int{1, 2, 3}

	var times= [5]int{1,5,7,8,9}
	for range times {
		fmt.Println()
	}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	now:=time.Now()
	str:=now.Format("20060102")
	fmt.Println(str)

	fmt.Println(x)
}
