package main

import (
	"fmt"
)

func CalcSquare(index int, value int64, ptr *[]int64) {
	(*ptr)[index] = value * value
}

func main() {
	massive := []int64{2, 4, 6, 8, 10}
	for i, j := range massive {
		go CalcSquare(i, j, &massive)
	}
	fmt.Println(massive)
}
