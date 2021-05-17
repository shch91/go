package main

import "fmt"

//Nim游戏
func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println(canWinNim(55432))
	fmt.Println("fa")
}
