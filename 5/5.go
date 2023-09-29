package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	N := 5
	var i int
	N *= 1000000000
	responseTimeout := time.Duration(N)
	go func() {
		for {
			v, ok := <-channel
			if ok != false {
				fmt.Println(v)
			} else {
				break
			}
		}
	}()
	deadline := time.Now().Add(responseTimeout)
	for time.Now().Before(deadline) {
		channel <- i
		i += 1
	}
	close(channel)
}
