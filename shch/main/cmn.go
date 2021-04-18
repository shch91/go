package main

import "fmt"

var result[][]int

//选择元素cmn
func cmn(arr, sel []int, m, n int) {
	if n == 0 {
		var val []int
		for _, v := range sel {
			val = append(val, v)
		}
		result = append(result, val)
		return
	}
	//下标选择
	for i := n; i <= m; i++ {
		sel[n-1] = arr[i-1]
		//后续再选n-1个
		cmn(arr, sel, i-1, n-1)
	}
}

func main()  {

	arr:= []int{1, 2, 3,4,5,6,7,8,9,10}
	var sel=make([]int,3)

	cmn(arr, sel, 10, 3)
	for _, v := range result {
		for _, a := range v {
			fmt.Printf("%d \t",a)
		}
		fmt.Println()
	}

}
