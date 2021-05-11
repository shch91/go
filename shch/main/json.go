package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Country  string
	Province string
}

type Person struct {
	Addr   *Address `json:"address"`
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Salary float64  `json:"salary"`
}

type Student struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

// 对数据进行序列化
func marshalData(data interface{}) []byte {
	data, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("序列化失败：%s\n", err)
	}
	fmt.Printf("序列化结果字节数组：%v\n", data)
	fmt.Printf("序列化结果字符串：%s\n", data)
	return data.([]byte)
}

// 对 struct 进行序列化
func marshalStruct() []byte {
	person := Person{
		Addr: &Address{Country: "ZH",Province: "re"},
		Name:   "佩奇",
		Age:    18,
		Salary: 99.99,
	}
	fmt.Printf("原始的数据：%+v\n", person)
	return marshalData(person)
}

// 反序列化成 struct
func unmarshalStruct(data []byte) {
	// 定义一个 Student 实例
	var person Person

	err := json.Unmarshal([]byte(data), &person)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("反序列化结果：", person)
}

// 对 map 进行序列化
func marshalMap() []byte {
	m := make(map[string]interface{})
	m["name"] = "盲僧"
	m["age"] = 10
	m["hobby"] = [2]string{"篮球", "游泳"}
	fmt.Printf("原始的数据：%v\n", m)
	return marshalData(m)
}

// 对 map 切片进行序列化
func marshalSlice() []byte {
	var slice []map[string]interface{}
	m1 := map[string]interface{}{
		"name": "妖姬",
		"age":  20,
	}
	slice = append(slice, m1)
	m2 := map[string]interface{}{
		"name": "德玛",
		"age":  30,
	}
	slice = append(slice, m2)
	fmt.Printf("原始的数据：%v\n", slice)
	return marshalData(slice)
}



// 反序列化成 map
func unmarshalMap(data []byte) {
	// 定义一个 map
	var m map[string]interface{}
	// 注意：反序列化 map，不需要 make，因为 make 操作被封装到了 Unmarsha 函数中
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("反序列化结果：", m)
}

// 反序列化成 slice
func unmarshalSlice(data []byte) {
	// 定义一个 silce
	var slice []map[string]interface{}
	// 注意：反序列化 slice，不需要 make，因为 make 操作被封装到了 Unmarsha 函数中
	err := json.Unmarshal([]byte(data), &slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("反序列化结果：", slice)
}

func main() {
	// 序列化
	result1 := marshalStruct()
	unmarshalStruct(result1)
	//fmt.Println()
	//result2 := marshalMap()
	//fmt.Println()
	//result3 := marshalSlice()
	//fmt.Println()
	// 反序列化

	//unmarshalMap(result2)
	//unmarshalSlice(result3)
}
