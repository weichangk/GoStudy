package main

import "fmt"

func main() {
	fmt.Println("数组")
	// Go 语言提供了数组类型的数据结构。 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
	// 数组元素可以通过索引（位置）来读取（或者修改），索引从0开始，第一个元素索引为 0，第二个索引为 1，以此类推。数组的下标取值范围是从0开始，到长度减1。
	// 数组一旦定义后，大小不能更改。

	array()

	mulitArray()

	arrayIsValueType()

	sort()
}

func array() {
	// 需要指明数组的大小和存储的数据类型。
	var a [3]int
	fmt.Println(a)

	var b = [3]string{"11", "22", "33"}
	fmt.Println(b)

	var c = [5]int{'A', 'B', 'C', 'D', 'E'} // byte
	fmt.Println(c)

	d := [...]int{10, 11, 12}
	fmt.Println(d)

	e := [5]int{1: 100}
	fmt.Println(e)

	f := [...]int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
	fmt.Println(f)

	// 访问数组元素
	var n [10]int

	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("n[%d] = %d\n", i, n[i])
	}

	// 数组长度 实际存储的数据量
	l := [...]float32{0, 1.1, 2.3, 3.3}
	fmt.Println("l len is ", len(l))

	// 数组容量 能够存储最大数量 因为数组是定长，所以长度等于容量
	ll := [10]int{1, 2, 3}
	fmt.Printf("ll len is %d cap is %d", len(ll), cap(ll))

	// 遍历数组
	for i := 0; i < len(l); i++ {
		fmt.Printf("l[%d] = %f\n", i, l[i])
	}

	// range遍历数组
	for i, x := range ll {
		fmt.Printf("ll[%d] = %d\n", i, x)
	}
}

func mulitArray() {
	// Go 语言支持多维数组，以下为常用的多维数组声明语法方式
	// var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

	a := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(a)
	fmt.Printf("a[1][2] = %d", a[1][2])
}

func arrayIsValueType() {
	// 数组是值类型 Go中的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。
	a := [...]string{"China", "USA", "UK", "HK"}
	b := a
	b[0] = "中国"
	fmt.Println(a, b)

	// 数组的大小是类型的一部分。因此[5]int和[25]int是不同的类型。因此，数组不能被调整大小。不要担心这个限制，因为切片的存在是为了解决这个问题。
	// c := [3]int{5, 78, 8}
	// var d [5]int
	// d = c //not possible since [3]int and [5]int are distinct types

}

func sort() {
	a := [...]int{12, 43, 11, 3, 7, 90, 34, 100, 2}
	fmt.Println(a)
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
