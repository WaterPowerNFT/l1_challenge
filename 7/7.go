package main

import (
	"fmt"
	"sync"
)

// сделать специальный тип для записи данных
type MutexMap struct {
	sync.RWMutex
	data map[string]int
}

// записывать данные через этот метод
func (this *MutexMap) WriteData(key *string, data int) {
	this.Lock()
	this.data[*key] = data
	this.Unlock()
}

func main() {
	var map_with_mutex MutexMap
	map_with_mutex.data = make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(5)
	array := []string{"cat", "dog", "pony", "frog", "crocodile"}
	for i := 0; i < 5; i += 1 {
		go func(index int) {
			map_with_mutex.WriteData(&array[index], index)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(map_with_mutex.data)
}
