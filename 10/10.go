package main

import (
	"fmt"
)

func main() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	d := map[int][]float64{}
	for _, elem := range array {
		new_elem := int(elem / 10)
		new_elem *= 10
		d[new_elem] = append(d[new_elem], elem)
	}
	fmt.Println(d)
}
