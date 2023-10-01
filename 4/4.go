package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	N := 5
	chanel_data := make(chan int)
	chanel_exit := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(N + 1)
	//2 канала, потому что удобно. Теоретически потом можно довести до ума и выключать определенных воркеров (добавляет масштабируемость коду)
	for i := 0; i < N; i += 1 {
		go func(num int) {
			summed := 0
			for {
				select {
				case got_num := <-chanel_data:
					summed += got_num
				case <-chanel_exit:
					fmt.Printf("worker %d finished job, was summed num up to %d\n", num, summed)
					wg.Done()
					return
				}
			}
		}(i)
	}
	//1 воркер нужен, чтобы писать данные в канал
	go func() {
		for {
			select {
			case <-chanel_exit:
				fmt.Printf("worker writer finished job\n")
				wg.Done()
				return
			default:
				num_to_send := rand.Int() % 100
				chanel_data <- num_to_send
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	//
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	for i := 0; i <= N; i += 1 {
		chanel_exit <- true
	}
	wg.Wait()
	fmt.Println("program has finished")
}
