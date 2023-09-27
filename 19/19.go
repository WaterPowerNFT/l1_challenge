package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReverseStr(str_to_reverse *string) {
	rune_string := strings.Split(*str_to_reverse, "")
	last := len(rune_string) - 1
	for i := 0; i < (len(rune_string) / 2); i += 1 {
		temp := rune_string[i]
		rune_string[i] = rune_string[last-i]
		rune_string[last-i] = temp
	}
	(*str_to_reverse) = strings.Join(rune_string, "")
}
func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		ReverseStr(&text)
		fmt.Println(text)
	}
}
