package main

import (
	"fmt"
)

func main() {
	/*
		程序的流程控制结构一共有三种：顺序结构，选择结构，循环结构。
		顺序结构：从上向下，逐行执行。
		选择结构：条件满足，某些代码才会执行。0-1次
		​分支语句：if，switch，select
		循环结构：条件满足，某些代码会被反复的执行多次。0-N次
		​循环语句：for
	*/

	ifFunc()
	switchFunc()
}

func ifFunc() {
	if true {
		fmt.Println("1")
	}

	a := "aaa"
	if a == "" {
		fmt.Println(a + " 2")
	} else {
		fmt.Println(a + " 3")
	}

	if a == "a" {
		fmt.Println(a + " 4")
	} else if a == "aa" {
		fmt.Println(a + " 5")
	} else {
		fmt.Println(a + " 6")
	}

	// if 变体
	// 如果其中包含一个可选的语句组件(在评估条件之前执行)。
	// 需要注意的是，num的定义在if里，那么只能够在该if..else语句块中使用，否则编译器会报错的。
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

func switchFunc() {
	// switch是一个条件语句，它计算表达式并将其与可能匹配的列表进行比较，并根据匹配执行代码块。它可以被认为是一种惯用的方式来写多个if else子句。
	// switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。 switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break。
	// 而如果switch没有表达式，它会匹配true
	// Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。

	var grade = "a"
	var marks = 80
	switch marks {
	case 90:
		grade = "a"
	case 80:
		grade = "b"
	case 70, 60:
		grade = "c"
	default:
		grade = "d"
	}
	fmt.Println(grade)

	switch grade {
	case "a":
		fmt.Println("优秀")
	case "b", "c":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 如需贯通后续的case，就添加fallthrough
	switch x := 100; x {
	default:
		fmt.Println(x)
	case 100:
		x += 1
		fmt.Println(x)
		fallthrough
	case 200:
		x += 1
		fmt.Println(x)
	case 300:
		x += 1
		fmt.Println(x)
	}

	// case中的表达式是可选的，可以省略。如果该表达式被省略，则被认为是switch true，并且每个case表达式都被计算为true，并执行相应的代码块。
	switch {
	case marks >= 90:
		fmt.Println("优秀")
	case marks >= 70 && marks < 90:
		fmt.Println("良好")
	case marks >= 60 && marks < 70:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// switch的注意事项
	// case后的常量值不能重复
	// case后可以有多个常量值
	// fallthrough应该是某个case的最后一行。如果它出现在中间的某个地方，编译器就会抛出错误。

	// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
