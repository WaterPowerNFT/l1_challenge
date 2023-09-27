package main

import (
	"fmt"
	"sync"
)

func CalcSquare(value int64, ch *chan int64) {
	(*ch) <- value * value
}

func main() {
	var wg sync.WaitGroup
	massive := []int64{2, 4, 6, 8, 10}
	channel := make(chan int64)
	var answer int64

	go func() {
		for i := 0; i < len(massive); i += 1 {
			temp := (<-channel)
			answer += temp
		}
		wg.Done()
	}()
	wg.Add(1)
	for _, j := range massive {
		wg.Add(1)
		go func(j int64) {
			CalcSquare(j, &channel)
			wg.Done()
		}(j)
	}
	wg.Wait()
	fmt.Println(answer)
}
