package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "résumé"
	println(str)
	for i, s := range str {
		fmt.Printf("%d contains %v", i, s)
		fmt.Println()
	}
	my_rune := []rune(str)
	fmt.Println("-----------------------------------")
	for i, s := range my_rune {
		fmt.Printf("%d contains %v", i, s)
		fmt.Println()
	}
	builder := strings.Builder{}
	builder.WriteString(str)
	builder.WriteString(" of ram")
	fmt.Println(builder.String())
}
