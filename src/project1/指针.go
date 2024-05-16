package main

import "fmt"

/*
//变量是一种使用方便的占位符，用于引用计算机内存地址。
//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
func main() {
	var a int = 10
	fmt.Printf("变量的地址是：%x\n", &a)
}
*/

/*
//如何使用指针
func main() {
	var a int = 20 //声明实际变量
	var ip *int    //声明指针变量

	ip = &a //指针变量的存储地址

	fmt.Printf("a变量的地址是：%x\n", ip)

	//指针变量的存储地址
	fmt.Printf("ip变量存储的指针地址：%x\n", ip)

	//使用指针访问值
	fmt.Printf("*ip 变量的值：%d\n", *ip)
}
*/

// 空指针。一个指针的变量通常缩写为ptr
func main() {
	var ptr *int
	fmt.Printf("ptr的值为： %x\n", ptr)
}
