package main

import (
	"fmt"
)

func RemoveElemFromIntSlice(mas *[]int, x int) {
	if x > len(*mas) {
		return
	}
	(*mas) = append((*mas)[:x], (*mas)[x+1:]...)
}

func RemoveElemFromStringSlice(mas *[]string, x int) {
	if x > len(*mas) {
		return
	}
	(*mas) = append((*mas)[:x], (*mas)[x+1:]...)
}

func main() {
	some_str_slice := []string{"1", "232", "1432", "fdmkb", "mdkbmkr"}
	some_int_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("some str slice == %v\n", some_str_slice)
	fmt.Printf("some int slice == %v\n", some_int_slice)
	RemoveElemFromStringSlice(&some_str_slice, 3)
	RemoveElemFromIntSlice(&some_int_slice, 100)
	fmt.Printf("some str slice == %v\n", some_str_slice)
	fmt.Printf("some int slice == %v\n", some_int_slice)
}
