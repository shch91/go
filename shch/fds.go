package main

import (
	"fmt"
)

func main() {
	for a := 1; a < 10; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%d * %d =%d \t", a, b, a*b)
		}
		fmt.Println()
	}
}
