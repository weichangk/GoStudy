package main

import "fmt"

func main() {
	// 结构体：是由一系列具有相同类型或不同类型的数据构成的数据集合
	// 结构体成员是由一系列的成员变量构成，这些成员变量被称为“字段”

	chushihuajiegouti()

	structIsValue()

	anonymousStruct()

	nestStruct()
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

func structIsValue() {
	fmt.Println("\n结构体是值类型")
	p1 := Person{"zhangsan", 18, "shenzhen"}
	fmt.Println(p1)                 // {zhangsan 18 shenzhen}
	fmt.Printf("%p, %T\n", &p1, p1) // 0xc00013c510, main.Person

	p2 := p1
	fmt.Println(p2)                 // {zhangsan 18 shenzhen}
	fmt.Printf("%p, %T\n", &p2, p2) // 0xc0000c05a0, main.Person // 地址不一样了

	p2.name = "lisi"
	fmt.Println(p1) // {zhangsan 18 shenzhen}
	fmt.Println(p2) // {lisi 18 shenzhen} //值类型深拷贝

	fmt.Println("\n定义结构体指针")
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)                  // &{zhangsan 18 shenzhen}
	fmt.Printf("%p, %T\n", &pp1, pp1) // 0xc000006030, *main.Person
	fmt.Println(*pp1)                 // {zhangsan 18 shenzhen}

	(*pp1).name = "wangwu"
	pp1.age = 35
	fmt.Println(*pp1) // {wangwu 35 shenzhen}
	fmt.Println(p1)   // {wangwu 35 shenzhen}

	// new函数， go语言专门用于创建某种类型的指针(不是nil)的函数
	pp2 := new(Person)
	fmt.Printf("%p, %T\n", pp2, pp2) // 0xc0000c0750, *main.Person
	*&pp2.name = "haha"
	pp2.age = 100
	fmt.Println(*pp2) // {haha 100 }
}

func anonymousStruct() {
	fmt.Println("\n匿名结构体")
	s1 := struct {
		name string
		age  int
	}{
		name: "张三",
		age:  18,
	}
	fmt.Println(s1) // {张三 18}

	fmt.Println("\n匿名结构体匿名字段")
	s2 := struct {
		string // 每种类型最多只能有个一默认字段
		int
	}{
		"hello",
		18,
	}
	fmt.Println(s2) // {hello 18}

}

type address struct {
	city, state string
}
type people struct {
	name    string
	age     int
	address address
}

type people2 struct {
	name    string
	age     int
	address *address
}

func nestStruct() {
	fmt.Println("\n嵌套结构体")
	// 一个结构体可能包含一个字段，而这个字段也可以是一个结构体，这种结构体被称为嵌套结构体

	var p people
	p.name = "里斯"
	p.age = 10
	var a = address{
		city:  "shenzhen",
		state: "city",
	}
	p.address = a
	fmt.Println(p)      // {里斯 10 {shenzhen city}}
	a.city = "shanghai" //值类型数据传递
	fmt.Println(p)      // {里斯 10 {shenzhen city}}

	var p2 = people2{
		name: "gogogo",
		age:  8,
	}
	p2.address = &a //指针类型引用传递
	a.city = "北京"
	fmt.Println(p2.address.city) // 北京
}
