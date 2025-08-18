package main

import "fmt"

func main() {
	var arr_int [5] int32 = [5] int32 {1,2,3,4,5}	
	fmt.Printf("arr_int %v with length %d and capacity %d\n", arr_int, len(arr_int), cap(arr_int))
	int_arr := [...]int32{1,2,3}
	fmt.Printf("int_arr %v with length %d and capacity %d\n", int_arr, len(int_arr), cap(int_arr))	
	var arr_int_slice []int32 = []int32{1,2,3}
	fmt.Printf("arr_int_slice %v with length %d and capacity %d\n", arr_int_slice, len(arr_int_slice), cap(arr_int_slice))
	arr_int_slice = append(arr_int_slice, 4)
	fmt.Printf("arr_int_slice %v with length %d and capacity %d\n", arr_int_slice, len(arr_int_slice), cap(arr_int_slice))
	
	
}