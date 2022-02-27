package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n指针")

	// 指针是存储另一个变量的内存地址的变量。
	// 变量是一种使用方便的占位符，用于引用计算机内存地址。
	// 一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。

	aboutPointer()

	arrPonit()

	PonitArr()

	PointPoint()

	funcPoint()

	pointParam()
}

func aboutPointer() {
	fmt.Println("\n认识指针")
	// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	var a int = 10
	fmt.Printf("a 变量的地址是: %x\n", &a) // a 变量的地址是: c0000120d0

	// 声明指针 var var_name *var-type
	var p1 *int
	fmt.Println(p1) // <nil> 空指针
	// Go 空指针 当一个指针被定义后没有分配到任何变量时，它的值为 nil。 nil 指针也称为空指针。
	// nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。 一个指针变量通常缩写为 ptr。

	p1 = &a                            // 指针变量的存储地址
	fmt.Printf("p1 变量的存储地址: %x\n", p1) // p1 变量的存储地址: c0000120d0

	// 获取指针的值，意味着访问指针指向的变量的值。语法是：*a
	fmt.Printf("*p1 变量的值: %d\n", *p1) // *p1 变量的值: 10

	// 操作指针改变变量的数值
	*p1 = 200
	fmt.Println(a) // 200

	// 使用指针传递函数的参数
	var change = func(val *int) {
		*val += 20
	}
	change(p1)
	fmt.Println(a) // 220
	change(&a)
	fmt.Println(a) // 240

	// 不要将一个指向数组的指针传递给函数。使用切片。
	// 虽然将指针传递给一个数组作为函数的参数并对其进行修改，但这并不是实现这一目标的惯用方法。我们有切片。
	var modifyUseArr = func(arr *[3]int) {
		(*arr)[0] = 100
		arr[1] = 200
	}
	arr := [3]int{89, 90, 91}
	modifyUseArr(&arr)
	fmt.Println(arr) // [100 200 91]

	// var modifyUseSlice(sls []int) {
	// 	sls[0] = 1000
	// }
	// modifyUseSlice(arr[:])
	// fmt.Println(arr) // [100 200 91]
}

func arrPonit() {
	fmt.Println("\n数组指针")
	//数组指针是一个指针，一个数组的地址

	//创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1) // [1 2 3 4]

	// 创建一个指针，存储该数组的地址：数组指针
	var p1 *[4]int // 数组指针
	p1 = &arr1
	fmt.Println(p1)         // &[1 2 3 4]
	fmt.Printf("%p\n", p1)  // 0xc000010240 数组arr1的地址
	fmt.Printf("%p\n", &p1) // 0xc000006030 p1指针的地址

	// 根据数组指针，操作数组
	p1[0] = 10 //简化写法，无需*
	(*p1)[1] = 20
	fmt.Println(arr1) // [10 20 3 4]
}

func PonitArr() {
	fmt.Println("\n指针数组")
	a := 1
	b := 2
	c := 3
	d := 4
	arr1 := [4]int{a, b, c, d}
	arr2 := [4]*int{&a, &b, &c, &d} // 指针数组
	fmt.Println(arr1)               // [1 2 3 4]
	fmt.Println(arr2)               // [0xc000012138 0xc000012150 0xc000012158 0xc000012160]
	arr1[0] = 100
	fmt.Println(a)    // 1
	fmt.Println(arr1) // [100 2 3 4]
	for i := 0; i < len(arr2); i++ {
		fmt.Println(*arr2[i])
		// 1
		// 2
		// 3
		// 4
	}
	*arr2[1] = 200
	fmt.Println(b)    // 200
	fmt.Println(arr1) // [100 2 3 4]
	for i := 0; i < len(arr2); i++ {
		fmt.Println(*arr2[i])
		// 1
		// 200
		// 3
		// 4
	}

}

func PointPoint() {
	fmt.Println("\n指针的指针")
	//如果这个指针变量存放的又是另一个只指针的变量的地址，则称这个指针的变量为指向指针的指针变量

	var a int
	var ptr *int
	var pptr **int
	a = 3000
	ptr = &a                                      // 指针ptr地址
	pptr = &ptr                                   // 指向指针 ptr 地址
	fmt.Printf("变量 a = %d\n", a)                  // 变量 a = 3000
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)          // 指针变量 *ptr = 3000
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr) // 指向指针的指针变量 **pptr = 3000

}

func funcPoint() {
	fmt.Println("\n函数指针与指针函数")
	// 函数指针：一个指针，指向了一个函数的指针
	// go 语言中，function 默认看作一个指针，没有*
	var a = func() {
		fmt.Println("我是一个函数指针测试")
	}
	a()

	// 指针函数：一个函数，该函数返回值是一个指针
	var arr1 = func() *[4]int {
		arr := [4]int{1, 2, 3, 4}
		return &arr
	}
	arr2 := arr1()
	fmt.Printf("arr2的类型%T, 地址%p, 数值%v", arr2, &arr2, arr2) // arr2的类型*[4]int, 地址0xc000006038, 数值&[1 2 3 4]

}

func pointParam() {
	fmt.Println("\n指针做函数参数")
	var swap = func(x *int, y *int) {
		var tmp int
		tmp = *x
		*x = *y
		*y = tmp
	}

	var a, b int = 10, 20
	swap(&a, &b)
	fmt.Println(a) // 20
	fmt.Println(b) // 10
}
