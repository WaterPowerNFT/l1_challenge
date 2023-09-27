package main

import (
	"fmt"
	"time"
)

// можно еще сделать с активным ожиданием, но максимально не охото такое решение показывать...
func SleepSeconds(seconds_to_sleep int) {
	<-time.After(time.Second * time.Duration(seconds_to_sleep))
}

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC850))
	sleep_time := 10

	fmt.Printf("sleep for %d seconds\n", sleep_time)

	SleepSeconds(sleep_time)
	t = time.Now()
	fmt.Println(t.Format(time.RFC850))
}
