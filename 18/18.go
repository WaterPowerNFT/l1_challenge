package main

import (
	"fmt"
	"sync"
)

type IncStruct struct {
	field uint64
}

func (this *IncStruct) IncByOne() {
	this.field += 1
}

func main() {
	var wg sync.WaitGroup
	var var_to_inc IncStruct
	for i := 0; i < 100; i += 1 {
		wg.Add(1)
		go func() {
			var_to_inc.IncByOne()
			wg.Done()
		}()
	}

	wg.Wait()
	//при плохом распаралеливании потоков здесь может быть любое число до 100
	//у меня всегда выводится 100 без использования мютексов
	fmt.Println(var_to_inc.field)
}
