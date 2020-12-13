package main

import (
	"fmt"
)

func main() {
	for a := 1; a < 10; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%d * %d =%d  ", b, a, a*b)
		}
		fmt.Println()
	}
}
