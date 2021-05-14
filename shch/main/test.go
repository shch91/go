package main

import (
	"fmt"
	"sync"
	"time"
)

func onceDo() {
	var num int
	sign := make(chan bool)
	var once sync.Once
	//闭包函数
	f := func(ii int) func() {
		return func() {
			num = num + ii*2
			sign <- true
		}
	}
	for i := 0; i < 3; i++ {
		fi := f(i + 1)
		go once.Do(fi)
	}
	for j := 0; j < 3; j++ {
		select {
		case <-sign:
			fmt.Println("Received a signal.")
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
		}
	}
	fmt.Printf("Num: %d.\n", num)
}

func main() {
	onceDo()
}
