package main

import "fmt"

func main() {
	// 结构体：是由一系列具有相同类型或不同类型的数据构成的数据集合
	// 结构体成员是由一系列的成员变量构成，这些成员变量被称为“字段”

	chushihuajiegouti()
}

// 结构体定义
type Person struct {
	name    string
	age     int
	address string
}

func chushihuajiegouti() {
	fmt.Println("\n初始化结构体")
	//方法一：
	var p1 Person
	fmt.Println(p1) // { 0 }
	p1.name = "hello go"
	p1.age = 10
	p1.address = "shenzhen"
	fmt.Printf("姓名：%s, 年龄：%d, 地址：%s\n", p1.name, p1.age, p1.address) // 姓名：hello go, 年龄：10, 地址：shenzhen

	//方法二：
	p2 := Person{}
	fmt.Println(p2) // { 0 }
	p2.name = "hello go"
	p2.age = 10
	p2.address = "shenzhen"
	fmt.Printf("姓名：%s, 年龄：%d, 地址：%s\n", p2.name, p2.age, p2.address) // 姓名：hello go, 年龄：10, 地址：shenzhen

	//方法三：
	p3 := Person{
		name:    "sss",
		age:     18,
		address: "ddd",
	}
	fmt.Println(p3) // {sss 18 ddd}

	//方法四：
	p4 := Person{"qqq", 100, "eee"}
	fmt.Println(p4) // {qqq 100 eee}

}
