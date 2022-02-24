package main

import "fmt"

func main() {
	fmt.Println("\nstrig 类型测试")

	stringUse()

	// 访问strings包，可以有很多操作string的函数。

	// 访问strconv包，可以实现string和其他数值类型之间的转换。
}

func stringUse() {
	fmt.Println("\nstrig 使用测试")

	// Go 中的字符串是一个字节的切片。可以通过将其内容封装在“”中来创建字符串。Go 中的字符串是 Unicode 兼容的，并且是 UTF-8 编码的。
	s1 := "hello world"
	fmt.Println(s1)

	// 访问字符串中的单个字节
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}
	for i := 0; i < len(s1); i++ {
		fmt.Printf("\n%d %c", s1[i], s1[i])
	}
}

// string 常用操作类 待补充
