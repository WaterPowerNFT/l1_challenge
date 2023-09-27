package main

import (
	"fmt"
	_ "fmt"
)

func createHugeString(_ int) string {
	some_str := "1234567890qwertyuiopasdfghjkl;'zxcvbnm,./qwertyuiop[]asdfg1234567890qwertyuiopasdfghjkl;'zxcvbnm,./qwertyuiop[]asdfg1234567890qwertyuiopasdfghjkl;'zxcvbnm,./qwertyuiop[]asdfg1234567890qwertyuiopasdfghjkl;'zxcvbnm,./qwertyuiop[]asdfg1234567890qwertyuiopasdfghjkl;'zxcvbnm,./qwertyuiop[]asdfg1234567890qwertyuiopasdfghjkl;'zxcvbnm,./qwertyuiop[]asdfg1234567890qwertyuiopasdfghjkl;'zxcvbnm,./qwertyuiop[]asdfg"
	return some_str
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//если возвращается длинная строка, то проблем не возникает. Если возвращается строка короче 100, то завершается с ошибкой
	//также в большом проекте могут возникнуть корки. Поэтому лучше всего проверять длину полученной строки, либо сделать через минимум, как я
	justString = v[:min(len(v), 100)]
}

func main() {
	someFunc()
	fmt.Println(justString)
}
