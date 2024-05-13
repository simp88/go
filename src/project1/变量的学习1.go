package main

import "fmt"

//全局变量的使用:任何一个方法都可以使用全局变量

var n5 = "class"

func main() {
	//局部变量的使用
	//先声明变量，在进行赋值
	var age int
	age = 18
	fmt.Println("age = ", age)
	//声明和赋值一起使用
	var years int = 22
	fmt.Println("小明今年", years, "岁了")
	//变量为0的输出
	var n1 int
	fmt.Print(n1)

	//fmt.Println("\n")

	//声明多个变量
	fmt.Println("声明多个变量")
	var n2, n3, n4 = 2, 3, 4
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Print(n4, "\n")

	//fmt.Println('\n')
	//使用全局变量
	fmt.Println("使用全局变量")
	fmt.Println("班级的英文这么拼写", n5)

	//在函数内部，可以通过花括号建立一个代码块，代码块彼此之间的变量作用域是相互独立的。例如下面的代码

}
