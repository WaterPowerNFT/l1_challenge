package main

import (
	"fmt"
)

type new_int int64

func (this *new_int) ChangeBit(bit_num uint) {
	if bit_num < 63 {
		(*this) ^= (1 << bit_num)
	}
}

func main() {
	var int64_variable new_int = 914832859
	fmt.Printf("Number decimal: %d\n Number bytes before: %08b\n", int64_variable, int64_variable)
	int64_variable.ChangeBit(0)
	fmt.Printf("After\n Number bytes before: %08b\n", int64_variable)
}
