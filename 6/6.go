package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(6)
	//1 горутина отработала заложенные ей команды
	go func() {
		fmt.Println("goroutine 1 finished tasks")
		wg.Done()
	}()
	//2 возникло внешнее условие в цикле, которое привело к завершению
	var i int
	go func() {
		for i != 1 {
			time.Sleep(1 * time.Second)
		}
		fmt.Println("goroutine 2 finished tasks")
		wg.Done()
	}()
	time.Sleep(2 * time.Second)
	i = 1

	//3 прослушиваемый канал отправил спецсимвол
	j := 0
	ch_go3 := make(chan int)
	go func() {
		temp := <-ch_go3
		for temp != 10 {
			temp = <-ch_go3
		}
		fmt.Println("goroutine 3 finished tasks")
		wg.Done()
	}()
	for j != 11 {
		ch_go3 <- max(j, 5)
		j += 1
	}

	//4 создать канал выхода
	ch_go4 := make(chan bool)
	go func() {
		var temp int
		for {
			select {
			case <-ch_go4:
				fmt.Printf(" temp is %d, goroutine 4 finishhed tasks\n", temp)
				wg.Done()
				return
			default:
				temp += 1
			}
		}
	}()
	time.Sleep(1 * time.Second)
	ch_go4 <- true

	//5 закрыть прослушиваемый канал
	ch_go5 := make(chan int)
	go func() {
		var answer int
		for {
			v, ok := <-ch_go5
			if !ok {
				fmt.Printf("answer is %d, goroutine 5 finishhed tasks\n", answer)
				wg.Done()
				return
			} else {
				answer += v
			}
		}
	}()
	ch_go5 <- 5
	time.Sleep(500 * time.Millisecond)
	ch_go5 <- 15
	time.Sleep(500 * time.Millisecond)
	close(ch_go5)
	//6 создать 2 канала. С одного ждать выхода, с другого выполнять операции
	ch_go6_data := make(chan int)
	ch_go6_exit := make(chan bool)
	go func() {
		var answer int
		var temp int
		for {
			select {
			case <-ch_go6_exit:
				fmt.Printf("answer is %d, goroutine 6 finishhed tasks\n", answer)
				close(ch_go6_data)
				close(ch_go6_exit)
				wg.Done()
				return
			case temp = <-ch_go6_data:
				answer += temp
			}
		}
	}()
	ch_go6_data <- 5
	time.Sleep(500 * time.Millisecond)
	ch_go6_data <- 125
	time.Sleep(500 * time.Millisecond)
	ch_go6_exit <- true
	wg.Wait()
}
