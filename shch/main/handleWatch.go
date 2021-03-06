package main

import (
	"fmt"
	"strconv"
)

//二进制手表
func readBinaryWatch(turnedOn int) []string {
	var result []string
	var minute string
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if count1(i)+count1(j) == turnedOn {
				if j < 10 {
					minute = "0" + strconv.Itoa(j)
				} else {
					minute = strconv.Itoa(j)
				}
				result = append(result, strconv.Itoa(i)+":"+minute)
			}
		}
	}
	return result
}



func main() {
	for _, str := range readBinaryWatch(2) {
		fmt.Printf(str + "\n")
	}
}
