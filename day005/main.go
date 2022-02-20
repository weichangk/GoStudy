package main

import (
	"fmt"
)

func main() {
	fmt.Println("循环语句")
	forFunc()
}

func forFunc() {
	// 循环语句表示条件满足，可以反复的执行某段代码。for是唯一的循环语句。(Go没有while循环)
	// 语法结构：for init; condition; post { }
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// for循环变体 所有的三个组成部分，即初始化、条件和post都是可选的。
	// for {} 类型 while(){}

	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
	numbers := [5]int{1, 2, 3, 4, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 numbers 的值 = %d\n", i, x)
	}
	fmt.Println()

	// 多层for循环 for循环中又有循环嵌套，就表示多层循环了。

	// 跳出循环的语句
	// break：跳出循环体。break语句用于在结束其正常执行之前突然终止for循环
	for i := 0; i < 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println()

	// continue：跳出一次循环。continue语句用于跳过for循环的当前迭代。在continue语句后面的for循环中的所有代码将不会在当前迭代中执行。循环将继续到下一个迭代。
	for i := 10; i > 0; i-- {
		if i < 5 || i > 8 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// goto：可以无条件地转移到过程中指定的行。
	var a int = 10
LABLE:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LABLE
		}
		fmt.Println(a)
		a++
	}
}
