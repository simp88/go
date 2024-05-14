package main

/*func main() {
	a, b := 1, 10
	if a > b {
		b++
	} else {
		b--
	}
	fmt.Println(a)
	fmt.Println(b)
}*/

/*func main() {
	score := 900
	var ans string
	if score >= 100 {
		ans = "S"
	} else if score >= 90 && score < 100 {
		ans = "A"
	} else if score >= 80 && score < 90 {
		ans = "B"
	} else if score >= 70 && score < 80 {
		ans = "C"
	} else if score > 60 && score < 70 {
		ans = "E"
	}
	fmt.Println(ans)
}
*/

/*func main() {
	str := "as"
	switch str {
	case "a":
		str += "a"
		str += "c"
	case "b":
		str += "bb"
		str += "aaaa"
	default: // 当所有case都不匹配后，就会执行default分支
		str += "CCCC"
	}
	fmt.Println(str)
}
*/

func main() {
	switch num := f(); {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
		fallthrough
	case num < 0:
		num += num
	}
}
func f() int {
	return 1
}
