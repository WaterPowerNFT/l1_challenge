package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckAllUniq(str_to_check *string) bool {
	map_to_store := map[rune]bool{}

	for _, elem := range *str_to_check {
		if map_to_store[elem] == true {
			return false
		} else {
			map_to_store[elem] = true
		}
	}

	return true
}

func CheckEngUnic(str_to_check *string) bool {
	map_to_store := map[rune]bool{}

	for _, elem := range *str_to_check {
		if elem >= 65 && elem <= 90 {
			elem = elem + 32
		}
		if map_to_store[elem] == true {
			return false
		} else {
			map_to_store[elem] = true
		}
	}

	return true
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		//fmt.Println(CheckEngUnic(&text))
		//Нижний вариант более масштабируем и удобен (можно и на русскую, и другую расскладку сделать), второй захардкожен под eng (можно добавить
		//еще условий для других языков, какие нужны)
		text = strings.ToLower(text)
		fmt.Println(CheckAllUniq(&text))
	}
}
