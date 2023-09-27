package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	massive := []int64{2, 4, 6, 8, 10}

	channel_x := make(chan int64)
	channel_x2 := make(chan int64)

	mas_len := len(massive)
	wg.Add(3)
	//write to channelx
	go func(length int) {
		for i := 0; i < mas_len; i += 1 {
			channel_x <- massive[i]
		}
		wg.Done()
	}(mas_len)

	//write to channel x2
	go func(length int) {
		for i := 0; i < mas_len; i += 1 {
			elem_from_ch1 := <-channel_x
			channel_x2 <- elem_from_ch1 * 2
		}
		wg.Done()
	}(mas_len)

	go func(length int) {
		for i := 0; i < mas_len; i += 1 {
			elem_from_ch2 := <-channel_x2
			fmt.Println(elem_from_ch2)
		}
		wg.Done()
	}(mas_len)

	wg.Wait()
}
