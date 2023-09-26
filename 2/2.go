package main

import (
	"fmt"
	"sync"
)

func CalcSquare(index int, value int64, ptr *[]int64) {
	(*ptr)[index] = value * value
}

func main() {
	var wg sync.WaitGroup
	massive := []int64{2, 4, 6, 8, 10}
	for i, j := range massive {
		wg.Add(1)
		go func(i int, j int64) {
			CalcSquare(i, j, &massive)
			wg.Done()
		}(i, j)
	}
	wg.Wait()
	fmt.Println(massive)
}
