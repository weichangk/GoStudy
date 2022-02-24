package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		函数是执行特定任务的代码块。
		go语言至少有一个main函数

		函数格式
		func funcName(parametername type1, parametername type2) (output1 type1, output2 type2) {
		//这里是处理逻辑代码
		//返回多个值
		return value1, value2
		}

		func：函数由 func 开始声明
		funcName：函数名称，函数名和参数列表一起构成了函数签名。
		parametername type：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
		output1 type1, output2 type2：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
		上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型。
		如果只有一个返回值且不声明返回值变量，那么你可以省略包括返回值的括号（即一个返回值可以不声明返回类型）
		函数体：函数定义的代码集合。
	*/

	var r1 = funcUse(299, 100)
	fmt.Println(r1) // 299

	argParam(1, 2, 3, 4)

	tranParam()

	funcReturn()

	blanklabel()

	deferFunc()

	deferMethods()

	stackdelay()

}

func funcUse(num1, num2 int) int {
	fmt.Println("\n函数使用测试")
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func paramUse() {
	// 形式参数：定义函数时，用于接收外部传入的数据，叫做形式参数，简称形参。
	// 实际参数：调用函数时，传给形参的实际的数据，叫做实际参数，简称实参。
	// 函数调用时函数名称必须匹配，实参与形参必须一一对应：顺序，个数，类型
}

func argParam(arg ...int) {
	fmt.Println("\n函数可变参数使用测试")
	// Go函数支持变参。接受变参的函数是有着不定数量的参数的。为了做到这点，首先需要定义函数使其接受变参
	// arg ...int告诉Go这个函数接受不定数量的参数。注意，这些参数的类型全部是int。在函数体中，变量arg是一个int的slice
	for _, n := range arg {
		fmt.Println(n)
	}
}

func tranParam() {
	fmt.Println("\n函数参数传递使用测试")
	// go语言函数的参数也是存在值传递和引用传递

	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	var r1 = getSquareRoot(9) // 值传递
	fmt.Println(r1)           // 3

	// 引用传递
	// 这就牵扯到了所谓的指针。我们知道，变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存。
	// 只有add1函数知道x变量所在的地址，才能修改x变量的值。所以我们需要将x所在地址&x传入函数，并将函数的参数的类型由int改为*int，即改为指针类型，才能在函数中修改x变量的值。
	// 此时参数仍然是按copy传递的，只是copy的是一个指针
	add1 := func(a *int) int {
		*a += 1
		return *a
	}
	x := 3
	var r2 = add1(&x)
	fmt.Println(r2) // 4
	fmt.Println(x)  // 4

	// 传指针使得多个函数能操作同一个对象。
	// 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。
	// 如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
	// Go 语言中slice，map 类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）
}

func funcReturn() {
	fmt.Println("\n函数返回值测试")
	// 一个函数被调用后，返回给调用处的执行结果，叫做函数的返回值。调用处需要使用变量接收该结果
	// 一个函数可以没有返回值，也可以有一个返回值，也可以有返回多个值。
	swap := func(x, y string) (string, string) {
		return x, y
	}
	a, b := swap("hello", "world")
	fmt.Println(a, b) // hello world
}

func blanklabel() {
	fmt.Println("\n空白标识符使用测试")
	// _ 是 Go 中的空白标识符。它可以代替任何类型的任何值。
	rectProps := func(lenght, width float64) (float64, float64) {
		var area = lenght * width
		var perimeter = (lenght + width) * 2
		return area, perimeter
	}
	var _, p = rectProps(10, 20) // _ 表示忽略返回值
	fmt.Println(p)               // 60
}

func funcScope() {
	fmt.Println("\n函数作用域测试")
	// 作用域：变量可以使用的范围。

	// 一个函数内部定义的变量，就叫做局部变量。变量在哪里定义，就只能在哪个范围使用，超出这个范围，我们认为变量就被销毁了。

	// 一个函数外部定义的变量，就叫做全局变量。所有的函数都可以使用，而且共享这一份数据

	// 函数的本质
	// 函数也是Go语言中的一种数据类型，可以作为另一个函数的参数，也可以作为另一个函数的返回值。
}

func deferFunc() {
	fmt.Println("\ndefer函数测试")
	// 延迟是什么? 即延迟（defer）语句，延迟语句被用于执行一个函数调用，在这个函数之前，延迟语句返回。

	// 可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。
	// 特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题

	/*
		func ReadWrite() bool {
			file.Open("file")
			defer file.Close()
			if failureX {
				return false
			} i
			f failureY {
				return false
			}
			return true
		}

		最后才执行file.Close()
	*/
	a := 1
	b := 2
	defer fmt.Println(b)
	defer fmt.Println("hello defer")
	fmt.Println(a)
}

type person struct {
	firstName string
	lastName  string
}

// 定义方法
func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func deferMethods() {
	fmt.Println("\n延迟一个方法测试")
	// 延迟并不仅仅局限于函数。延迟一个方法调用也是完全合法的。
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}

func stackdelay() {
	fmt.Println("\n堆栈推迟测试")
	// 当一个函数有多个延迟调用时，它们被添加到一个堆栈中，并在Last In First Out（LIFO）后进先出的顺序中执行。
	name := "hello golang"
	fmt.Println(name)
	for _, n := range name {
		defer fmt.Printf("%c", n)
	}
}

func deferNote() {
	// defer注意点
	// 当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行。
	// 当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回。
	// 当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数。
}
