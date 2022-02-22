package main

import "fmt"

func main() {
	fmt.Println("slice 切片")

	// Go 语言切片是对数组的抽象。 Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
	// 切片是一种方便、灵活且强大的包装器。切片本身没有任何数据。它们只是对现有数组的引用。
	// 切片与数组相比，不需要设定长度，在[]中不用设定值，相对来说比较自由

	// 从概念上面来说slice像一个结构体，这个结构体包含了三个元素：
	// 指针，指向数组中slice指定的开始位置
	// 长度，即slice的长度
	// 最大长度，也就是slice开始位置到数组的最后位置的长度

	sliceSyntax()

	arrayToSlice()

	sliceIsClassType()

	sliceCopy()
}

func sliceSyntax() {
	// 定义切片 var identifier []type
	// 切片不需要说明长度。 或使用make()函数来创建切片
	var slice1 []int = make([]int, 10)
	fmt.Println(slice1)

	slice2 := make([]string, 3)
	fmt.Println(slice2)
	fmt.Printf("%T\n", slice2)

	// 定义并初始化
	slice3 := []int{1, 2, 3}
	fmt.Println(slice3)
	fmt.Printf("%T\n", slice3)
	fmt.Println("-------------------")

	slice4 := make([]int, 3, 5)
	fmt.Printf("slice4 cap is %d len is %d\n", cap(slice4), len(slice4))
	fmt.Println(slice4) // 0 0 0
	// slice4[4] = 1       //error 只能操作长度内的元素，添加长度外容量内的元素应该使用append
	slice4 = append(slice4, 1)
	fmt.Println(slice4)
	fmt.Println(slice4)
	slice4 = append(slice4, 2)
	slice4 = append(slice4, 3)
	fmt.Println(slice4)
	fmt.Printf("slice4 cap is %d len is %d\n", cap(slice4), len(slice4)) //向同切片向容器添加元素时，如果容量不够则会自动翻倍
	slice4 = append(slice4, 4, 5, 6)
	fmt.Println(slice4)
	slice3[0], slice3[1], slice3[2] = 11, 22, 33
	slice4 = append(slice4, slice3...) //整个切片追加
	fmt.Println(slice4)
	fmt.Println("-------------------")

	//切片的遍历语法和数组一样的
}

func arrayToSlice() {
	// 通过数组初始化切片
	// s := arr[startIndex:endIndex]
	// 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片（前闭后开），长度为endIndex-startIndex
	a := [5]int{76, 77, 78, 79, 80}
	s1 := a[0:]
	fmt.Println(s1)
	fmt.Printf("a point id %p s1 point is %p\n", &a, s1)

	s2 := a[:]
	fmt.Println(s2)
	fmt.Printf("a point id %p s2 point is %p\n", &a, s2)

	s3 := a[2:4]
	fmt.Println(s3)
	fmt.Printf("a point id %p s3 point is %p, s3 len is %d cap is %d\n", &a, s3, len(s3), cap(s3))
	fmt.Println("-------------------")

	// 修改切片 slice没有自己的任何数据。它只是底层数组的一个表示。对slice所做的任何修改都将反映在底层数组中。
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr) //array before [57 89 90 82 100 78 67 69 59]
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr) // array after [57 89 91 83 101 78 67 69 59]
	fmt.Println("-------------------")

	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(b)
	slice1 := b[:5]  //1-5
	slice2 := b[3:8] //4-8
	slice3 := b[5:]  //6-10
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	fmt.Println("更改数组内容")
	b[4] = 100
	fmt.Println(b)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	fmt.Println("更改切片内容")
	slice2[2] = 200
	fmt.Println(b)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	fmt.Println("切片追加内容")
	slice1 = append(slice1, 1, 1, 1, 1)
	fmt.Println(b)      //[1 2 3 4 100 1 1 1 1 10]
	fmt.Println(slice1) //[1 2 3 4 100 1 1 1 1]
	fmt.Println(slice2) //[4 100 1 1 1]
	fmt.Println(slice3) //[1 1 1 1 10]

	fmt.Println("切片追加内容操作数组容量时会指向新的底层数组")
	fmt.Println(b)                                               //[1 2 3 4 100 1 1 1 1 10]
	fmt.Println(slice1)                                          //[1 2 3 4 100 1 1 1 1]
	fmt.Println(slice2)                                          //[4 100 1 1 1]
	fmt.Println(slice3)                                          //[1 1 1 1 10]
	fmt.Printf("b point id %p slice1 point is %p\n", &b, slice1) //b point id 0xc0000b0230 slice1 point is 0xc0000b0230
	slice1 = append(slice1, 2, 2, 2, 2)
	fmt.Println("---")
	fmt.Println(b)                                               //[1 2 3 4 100 1 1 1 1 10]
	fmt.Println(slice1)                                          //[1 2 3 4 100 1 1 1 1 2 2 2 2]
	fmt.Println(slice2)                                          //[4 100 1 1 1]
	fmt.Println(slice3)                                          //[1 1 1 1 10]
	fmt.Printf("b point id %p slice1 point is %p\n", &b, slice1) //b point id 0xc0000b0230 slice1 point is 0xc0000d80a0
}

func sliceIsClassType() {
	fmt.Println("----------sliceIsClassType-----------")
	//切片是引用类型， 数组是值类型

	// 每个切片引用一个底层数组
	// 切片本身不存储任何数据，都是切片指向的底层数组储存，所以修改切片则是通过切片修改数据
	// 当向切片添加数据时，如果没有超过容量则直接添加，如果超过容量则自动成倍扩容
	// 切篇一旦扩容，就是重新指向一个新的底层数组

	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("s1 point is %p cap is %d, len is %d\n", s1, cap(s1), len(s1))

	s1 = append(s1, 4, 5, 6)
	fmt.Println(s1)
	fmt.Printf("s2 point is %p cap is %d, len is %d\n", s1, cap(s1), len(s1))

	s2 := append(s1, 4, 5, 6)
	fmt.Println(s2)
	fmt.Printf("s2 point is %p cap is %d, len is %d\n", s2, cap(s2), len(s2))
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

}

func sliceCopy() {
	fmt.Println("\n深拷贝和浅拷贝测试")
	// 深拷贝
	// 拷贝的是数据本身。值类型的数据，默认都是深拷贝：array, int, float, string, bool, struct

	// 浅拷贝
	// 拷贝的是数据的地址。导致多个变量指向同一块内存/引用类型的数据默认都是浅拷贝：slice, map

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0) //len=0 cap=0
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1) // [1 2 3 4]
	fmt.Println(s2) // [1 2 3 4]

	s1[0] = 100
	fmt.Println(s1) // [100 2 3 4] // 深拷贝了，指的不是同一块内存
	fmt.Println(s2) // [1 2 3 4]

	s3 := append(s1, s2...) // 深拷贝
	s1[1] = 111
	s2[2] = 666
	fmt.Println(s1) // [100 111 3 4]
	fmt.Println(s2) // [1 2 666 4]
	fmt.Println(s3) // [100 2 3 4 1 2 3 4] // 不受影响

	fmt.Println("\ncopy() 函数测试")
	// copy() 函数可以实现slice 的深拷贝 copy方法是不会建立两个切片的联系的

	s4 := []int{7, 8, 9}
	fmt.Println(s2) // [1 2 666 4]
	fmt.Println(s4) // [7 8 9]

	// copy(s2, s4)    // 从源 s4 拷贝到目标 s2
	// fmt.Println(s2) // [7 8 9 4]
	// fmt.Println(s4) // [7 8 9]
	// copy(s4, s2)
	// fmt.Println(s2) // [1 2 666 4]
	// fmt.Println(s4) // [1 2 666]
	copy(s4, s2[1:3])
	fmt.Println(s2) // [1 2 666 4]
	fmt.Println(s4) // [2 666 9]
}
