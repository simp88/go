package main

import (
	"fmt"
	"os"
)

func main() {
	os.Stdout.WriteString("Hello 世界!")
	println("")
	println("Hello 世界!")
	fmt.Println("Hello 世界！")
	fmt.Printf("%%%s\n", "hello world")

}
