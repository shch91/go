package main

import (
	"fmt"
)

func main() {
	var address = "广东省佛山市禅城区祖庙街道卫国路43号地质大厦"
	tel := "18813093802"
	fmt.Printf("地址%s,\t电话：%s \n", address, tel)
	var st = "fdsfafs"
	for index, ch := range st {
		fmt.Printf("index=%d,val:%c \t", index, ch)

	}

}
