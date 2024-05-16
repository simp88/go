package main

import "fmt"

/*func main() {
	var n [10]int //n是一个长度为10的数组
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}*/

func main() {
	var i, j, k int

	//声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//输出数组元素
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	//输出每个数组的值
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//将索引1和3的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3)
	}
}
