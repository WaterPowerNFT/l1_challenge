package main

import (
	"fmt"
)

func LBinSearch(array_to_find *[]int, elem int) int {
	l, r := 0, 0
	for l, r = 0, len(*array_to_find)-1; l < r; {
		m := (l + r) / 2
		if elem <= (*array_to_find)[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	if (*array_to_find)[l] == elem {
		return l
	} else {
		return -1
	}
}

func RBinSearch(array_to_find *[]int, elem int) int {
	l, r := 0, 0
	for l, r = 0, len(*array_to_find)-1; l < r; {
		m := (l + r + 1) / 2
		if elem >= (*array_to_find)[m] {
			l = m
		} else {
			r = m - 1
		}
	}
	if (*array_to_find)[l] == elem {
		return l
	} else {
		return -1
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	N := 10
	find_lbin := LBinSearch(&array, N)
	if find_lbin == -1 {
		fmt.Printf("LBinSearch elem %d not in array\n", N)
	} else {
		fmt.Printf("LBinSearch elem %d in array on position %d\n", N, find_lbin)
	}
	find_rbin := RBinSearch(&array, N)
	if find_rbin == -1 {
		fmt.Printf("RBinSearch elem %d not in array\n", N)
	} else {
		fmt.Printf("RBinSearch elem %d in array on position %d\n", N, find_rbin)
	}
}
