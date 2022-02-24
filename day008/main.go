package main

import "fmt"

func main() {

	// map 是 Go 中的内置类型，它将一个值与一个键关联起来。可以使用相应的键检索值。
	// Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值 Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的，也是引用类型
	// 使用 map 过程中需要注意的几点：
	// 1. map 是无序的，每次打印出来的 map 都会不一样，它不能通过index获取，而必须通过 key 获取
	// 2. map 的长度是不固定的，也就是和 slice 一样，也是一种引用类型
	// 3. 内置的 len 函数同样适用于 map，返回 map 拥有的 key 的数量
	// 4. map 的 key 可以是所有可比较的类型，如布尔型、整数型、浮点型、复杂型、字符串型……也可以键。

	mapDeclare()

	mapUse()

	mapDelete()

	mapLen()

	mapIsType()

}

func mapDeclare() {
	fmt.Println("\nmap 声明和初始化测试")
	var m1 map[int]string // 声明变量，默认 map 是 nil
	// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对。 但是 nil 的 slice 可以直接使用

	var m2 = make(map[int]string)                           // 使用 make 函数 不为 nil
	var m3 = map[string]int{"go": 10, "c#": 20, "java": 30} // 声明并初始化
	fmt.Printf("m1: %v is nil = %t\n", m1, m1 == nil)       // m1: map[] is nil = true
	fmt.Printf("m2: %v is nil = %t\n", m2, m2 == nil)       // m2: map[] is nil = false
	fmt.Printf("m3: %v is nil = %t\n", m3, m3 == nil)       // m3: map[c#:20 go:10 java:30] is nil = false
	if m1 == nil {
		m1 = make(map[int]string)
	}
	fmt.Printf("m1: %v is nil = %t\n", m1, m1 == nil) // m1: map[] is nil = false
}

func mapUse() {
	fmt.Println("\nmap 使用测试")
	var m1 = make(map[int]string)
	m1[0] = "hello"
	m1[-1] = "..."
	m1[10] = "world"
	// m1[10] = "worldxxx" // 覆盖前面的
	fmt.Println(m1)    // map[-1:... 0:hello 10:world]
	fmt.Println(m1[0]) // hello

	// 通过 key 访问 map 时，当 key 不存在时，不会报错，会得到 value 值类型的默认值
	fmt.Println(m1[-2]) // ""
	// 可以通过 ok-idiom 获取值知道 key/value 是否存在
	v1, ok := m1[100] // ok 表示 key 是否存在，bool值；v1 表示 map 中 key 对应的 value，如果 key 不存在 v1 为 value 类型值的默认值
	if ok {
		fmt.Printf("m1[100] value = %v\n", v1)
	} else {
		fmt.Printf("m1[100] value 不存在！\n") // m1[100] value 不存在！
	}

	// 修改, 如果 key 不存在则添加数据
	m1[0] = "hello golang"
	fmt.Println(m1[0]) // hello golang
}

func mapDelete() {
	fmt.Println("\n使用 delete() 函数删除 map 测试")
	m1 := map[string]int{"go": 1, ".net": 2, "java": 3}
	fmt.Println(m1) // map[.net:2 go:1 java:3]
	delete(m1, ".net")
	fmt.Println(m1) // map[go:1 java:3]

	delete(m1, "c#") // 删除不存在的 key 不会报错
	fmt.Println(m1)  // map[go:1 java:3]

}

func mapLen() {
	fmt.Println("\nmap 长度测试")
	m1 := map[string]int{"go": 1, ".net": 2, "java": 3}
	fmt.Printf("len: %v\n", len(m1)) // len: 3
}

func mapIsType() {
	fmt.Println("\nmap 是引用类型测试")
	m1 := map[int]string{
		1: "html",
		2: "css",
		3: "jquery",
	}
	m1[100] = "go"
	fmt.Println(m1) // map[1:html 2:css 3:jquery 100:go]
	m11 := m1
	m11[100] = "golang"
	fmt.Println(m1)  // map[1:html 2:css 3:jquery 100:golang]
	fmt.Println(m11) // map[1:html 2:css 3:jquery 100:golang]
	// 与切片相似，映射是引用类型。当将映射分配给一个新变量时，它们都指向相同的内部数据结构。因此一个的变化会反映另一个。

	// map不能使用==操作符进行比较。==只能用来检查map是否为空。否则会报错：invalid operation: map1 == map2 (map can only be comparedto nil)
	// if(m1 == m11){}
}
