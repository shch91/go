package main

import "fmt"


// 结构体1，因为没有String()函数，所以该结构体指针的数组打印时会输出一串内存地址
type student struct {
	Age  int32
	Name string
}

// 结构体2，因为有了String()函数，所以该结构体指针的数组打印时会调用String()函数
type teacher struct {
	High int32
	Sex  int32
}
func (t *teacher) String() string {
	return fmt.Sprintf("{High:%d,Sex:%d}", t.High, t.Sex)
}

func main() {
	// 结构体指针的数组1
	arr1 := []*student{
		&student{Age: 1, Name: "111"},
		&student{Age: 2, Name: "222"},
	}
	fmt.Printf("打印结构体指针数组1：%v \n", arr1)

	// 结构体指针的数组2
	arr2 := []*teacher{
		&teacher{High: 170, Sex: 17},
		&teacher{High: 180, Sex: 18},
	}
	fmt.Printf("打印结构体指针数组2：%v \n", arr2)


	a := &student{Age: 1, Name: "xiaoming"}

	fmt.Printf("a=%v	\n", a)
	fmt.Printf("a=%+v	\n", a)
	fmt.Printf("a=%#v	\n", a)
}
