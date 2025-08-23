package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int)
	go db_call(c)
	for v := range c {
		fmt.Println("Received id ", v)
		time.Sleep(1 * time.Second)
	}
}

func db_call(c chan int) {
	defer close(c)
	for id := range []int{1, 2, 3, 4, 5} {
		fmt.Println("Sending id ", id)
		c <- id
	}
	fmt.Println("All ids sent")
}
