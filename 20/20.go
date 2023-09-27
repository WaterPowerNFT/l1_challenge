package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReverseStr(str_to_reverse *string) {
	word_string := strings.Split(*str_to_reverse, " ")
	last := len(word_string) - 1
	for i := 0; i < (len(word_string) / 2); i += 1 {
		temp := word_string[i]
		word_string[i] = word_string[last-i]
		word_string[last-i] = temp
	}
	(*str_to_reverse) = strings.Join(word_string, " ")
}
func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-2]
		ReverseStr(&text)
		fmt.Println(text)
	}
}
