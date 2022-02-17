package main

import "fmt"

func main() {
	variableFunc()
	constantFunc()
}

func variableFunc() {
	// 1. 声明一个变量
	// 指定变量类型，声明后若不赋值，使用默认值
	var q string
	q = "q"
	fmt.Println(q)

	// 根据值自行判定变量类型(类型推断Type inference)
	var w = 100
	fmt.Println(w)

	// 省略var, 注意 := 左侧的变量不应该是已经声明过的(多个变量同时声明时，至少保证一个是新变量)，否则会导致编译错误(简短声明)。
	// 这种方式它只能被用在函数体内，而不可以用于全局变量的声明与赋值
	// 在同一个作用域中，已存在同名的变量，则之后的声明初始化，则退化为赋值操作。但这个前提是，最少要有一个新的变量被定义，且在同一作用域
	e := true
	fmt.Println(e)
	ee, e := 200, false
	fmt.Println(ee, e)

	// 2. 多变量声明
	// 以逗号分隔，声明与赋值分开，若不赋值，存在默认值
	var r, t, y float32
	r, t, y = 10, 20, 30
	fmt.Println(r, t, y)

	// 直接赋值，下面的变量类型可以是不同的类型
	var u, i, o = true, "iii", 10
	fmt.Println(u, i, o)

	// 集合类型
	var (
		firstName = "hello"
		lastName  = "world"
		isSB      bool
	)
	isSB = true
	fmt.Println(firstName, lastName, isSB)

	// 3. 注意事项
	// 变量必须先定义才能使用, 单纯赋值不被使用会报异常
	// go语言是静态语言，要求变量的类型和赋值的类型必须一致。
	// 变量名不能冲突。(同一个作用于域内不能冲突)
	// 简短定义方式，左边的变量名至少有一个是新的
	// 简短定义方式，不能定义全局变量。
	// 变量的零值。也叫默认值。
	// 变量定义了就要使用，否则无法通过编译
}

func constantFunc() {
	// 1. 常量声明
	// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
	const AA string = "aa"       // 显式类型定义
	const BB float32 = 10.1      // 隐式类型定义
	const CC, DD = "hello", true //多重赋值
	fmt.Println(AA, BB, CC, DD)

	// 常量可以作为枚举，常量组
	// 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	const (
		UNkNOWN = 0
		FEMALE  = 1
		MALE    = 2
		OTHER
	)
	fmt.Println(UNkNOWN, FEMALE, MALE, OTHER)

	// 常量的注意事项
	// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
	// 不曾使用的常量，在编译的时候，是不会报错的
	// 显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，变量是可以是不同的类型值

	// 2. iota
	// iota，特殊常量，可以认为是一个可以被编译器修改的常量
	// iota 可以被用作枚举值
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	// iota的注意事项
	// 如果中断iota自增，则必须显式恢复。且后续自增值按行序递增
	// 自增默认是int类型，可以自行进行显示指定类型
	// 数字常量不会分配存储空间，无须像变量那样通过内存寻址来取值，因此无法获取地址
}
