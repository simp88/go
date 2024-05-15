package main

import "fmt"

/*//双循环打印九九乘法表
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i <= j {
				fmt.Printf("%d*%d=%3d   ", i, j, i*j)
			}
		}
		fmt.Println()
	}
}*/

/*func main() {
	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println(index, value)
	}
}*/

/*func main() {
	for i := 3; i <= 9; i++ {
		fmt.Printf("i是%d\n", i)
		for j := 0; j <= 9; j++ {
			if i == j {
				break //当i等于j时停止for循环
			}
			fmt.Println(i, j)
		}
	}
}
*/

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("i是%d\n", i)
		for j := 0; j < 8; j++ {
			if i == j {
				continue //当i等于j时将跳过i==j的循环
			}
			fmt.Println(i, j)
		}
	}
}
