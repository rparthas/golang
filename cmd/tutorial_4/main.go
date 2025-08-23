package main

import "fmt"

func main() {
	var arr_int [5]int32 = [5]int32{1, 2, 3, 4, 5}
	fmt.Printf("arr_int %v with length %d and capacity %d\n", arr_int, len(arr_int), cap(arr_int))
	int_arr := [...]int32{1, 2, 3}
	fmt.Printf("int_arr %v with length %d and capacity %d\n", int_arr, len(int_arr), cap(int_arr))
	var arr_int_slice []int32 = []int32{1, 2, 3}
	fmt.Printf("arr_int_slice %v with length %d and capacity %d\n", arr_int_slice, len(arr_int_slice), cap(arr_int_slice))
	arr_int_slice = append(arr_int_slice, 4)
	fmt.Printf("arr_int_slice %v with length %d and capacity %d\n", arr_int_slice, len(arr_int_slice), cap(arr_int_slice))

	var test_map map[string]int32 = map[string]int32{"one": 1, "two": 2}
	fmt.Printf("map %v with length %d\n", test_map, len(test_map))
	for key, value := range test_map {
		fmt.Printf("key %s value %d\n", key, value)
	}
	var val, ok = test_map["three"]
	fmt.Printf("key three value %d exists %v\n", val, ok)

	for i := range 10 {
		fmt.Printf("i is %d\n", i)
	}
}
