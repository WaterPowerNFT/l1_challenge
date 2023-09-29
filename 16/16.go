package main

import (
	"fmt"
)

func qsort(mas_to_sort []int) []int {
	if len(mas_to_sort) < 2 {
		return mas_to_sort
	}

	left, right := 0, len(mas_to_sort)-1

	pivot := right / 2

	mas_to_sort[pivot], mas_to_sort[right] = mas_to_sort[right], mas_to_sort[pivot]

	for i := range mas_to_sort {
		if mas_to_sort[i] < mas_to_sort[right] {
			mas_to_sort[left], mas_to_sort[i] = mas_to_sort[i], mas_to_sort[left]
			left++
		}
	}

	mas_to_sort[left], mas_to_sort[right] = mas_to_sort[right], mas_to_sort[left]

	qsort(mas_to_sort[:left])
	qsort(mas_to_sort[left+1:])

	return mas_to_sort
}

func main() {

	massive := []int{1, 5, 2, 41, 231, 6}
	qsort(massive)
	fmt.Println(massive)
}
