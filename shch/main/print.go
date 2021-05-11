package main

import "fmt"

/*
%v 打印结构体
%+V 打印带有字段的结构体
%T 打印对象类型
%t 打印布尔值
%d 打印整型数，十进制输出，如果d前面有数字，表示控制输出宽度，默认使用空白填充，%05d 会在不满5位时填充0
%b 打印整型数，二进制输出
%c 打印整型数，字符输出（如果有）
%o 打印整型数，八进制输出，如果x前面带有#表示带有零的八进制
%x 打印整型数，十六进制输出，如果x前面带有#表示带有0x的十六进制
%f 打印浮点数，正常输出，如果f前面有x.y 表示控制宽度和输出小数点位数，要对其，再加-，例如 %-10.5f
%e,%E 打印浮点数，科学记数法输出
%s 打印字符串
%q 打印字符串，带有引号输出
%x 打印字符串，使用base-16输出其编码
%p 打印指针
%U 打印Unicode字符
%#U 打印带字符的Unicode
*/
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
		{Age: 1, Name: "111"},
		{Age: 2, Name: "222"},
	}
	fmt.Printf("打印结构体指针数组1：%v \n", arr1)

	// 结构体指针的数组2
	arr2 := []*teacher{
		{High: 170, Sex: 17},
		{High: 180, Sex: 18},
	}
	fmt.Printf("打印结构体指针数组2：%v \n", arr2)

	a := &student{Age: 1, Name: "xiaoming"}

	fmt.Printf("a=%v	\n", a)
	fmt.Printf("a=%+v	\n", a)
	fmt.Printf("a=%#v	\n", a)
}
