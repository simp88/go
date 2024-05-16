package main

import "fmt"

/*// 函数返回两个数的最大值
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)

	fmt.Printf("最大值是: %d\n", ret)

}
*/

func swap(x, y, z string) (string, string, string) {
	return x, y, z
}

func main() {
	a, b, _ := swap("杨", "yang", "z")
	fmt.Println(a, b)
}
