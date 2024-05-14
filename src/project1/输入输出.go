package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*os.Stdout.WriteString("Hello 世界!")*/
	println("")
	println("Hello 世界!")
	fmt.Println("Hello 世界！")

	fmt.Printf("%%%s\n", "hello world")

	fmt.Printf("%b\n", 2<<7-1)

	fmt.Printf("%f\n", 1e2)
	//使用其它进制时，在%与格式化动词之间加上一个空格便可以达到分隔符的效果

	fmt.Println()
	str := "abcdefg"
	fmt.Printf("%x\n", str)
	fmt.Printf("% x\n", str)

	fmt.Printf("%s\n", "hello world")

	fmt.Println("--------------")

	/*	var s, s2 string
		fmt.Scanln(&s, &s2)
		fmt.Println(s, s2)
	*/
	/*	var s, s2, s3 string
		scanf, err := fmt.Scanf("%s %s %s", &s, &s2, &s3)
		if err != nil {
			fmt.Println(scanf, err)
		}
		fmt.Println(s)
		fmt.Println(s2)
		fmt.Println(s3)*/

	/*	//读
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println(scanner.Text())*/

	//写
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hello world! \n")
	writer.Flush()
	fmt.Println(writer.Buffered())

}
