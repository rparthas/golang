package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767 + 1
	fmt.Println(intNum)

	var floatNum float32 = 3.402823466 * 2
	fmt.Println(floatNum)

	var str string = `Hello, 
	World!`
	fmt.Println(len(str))

	str = "γ"
	fmt.Println(utf8.RuneCountInString(str))

	a, b := 10, 20
	fmt.Println(a + b)

	const pi = 3.14
	fmt.Println(pi)
}
