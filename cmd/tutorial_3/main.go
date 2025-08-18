package main

import (
	"fmt"
	"errors"
)

func main() {
	printMessage("Hello, World!")
	var result, remainder,err = intDiv(10, 3)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return
	}else{
		fmt.Printf("result: %v, remainder: %v\n", result, remainder)
	}
	fmt.Println("--------------------------------")
	switch remainder {
	case 0:
		fmt.Println("division is exact")
	case 1,2:
		fmt.Println("division is close")
	default:
		fmt.Println("division is not close enough")
	}
}

func printMessage(message string) {
	fmt.Println(message)
}

func intDiv(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("division by zero")
		return 0,0,err
	}
	result = numerator / denominator
	remainder = numerator % denominator
	return result, remainder, nil
}

