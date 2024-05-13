package main

import "fmt"

const name string = "jack"

func main() {
	fmt.Println(name)

	const (
		num1 = 1
		num2
		num3
	)
	fmt.Println(num1, num2, num3)

	// iota是一个内置的常量标识符，通常用于表示一个常量声明中的无类型整数序数，一般都是在括号中使用
	const (
		a1 = iota
		a2
		a3
	)
	fmt.Println(a1, a2, a3)

	// 多个常量声明在一行,iota相当于声明常量行数的索引值, 多个常量在一行声明,iota不会增加, 所以d1 = 1, d2 = 2, d3 = 2, d4 = 3
	const (
		d1, d2 = iota + 1, iota + 2
		d3, d4 = iota + 1, iota + 2
	)
	fmt.Println(d1, d2, d3, d4)

	//iota是递增的，第一个常量使用iota值的表达式，根据序号值的变化会自动的赋值给后续的常量，
	//直到用新的iota重置，这个序号其实就是代码的相对行号，是相对于当前分组的起始行号，看下面的例子
	const (
		b1 = iota<<2*3 + 1 //
		b2 = iota<<2*3 + 1
		_
		b4
		b5 = iota
		b6
		_
		b8 = iota
		b9
	)
	fmt.Println(b1, b2, b4, b5, b6, b8, b9)

	const (
		_  = iota
		KB = 1 << (10 * iota) // << 想当于向左移动一定的位数,  此时iota等于1,也就是1 << 10, 1左移10位,也就是10000000000,二进制换算成10进制,也就是1024
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(KB, MB, GB, TB, PB)

}
