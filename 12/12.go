package main

import (
	"fmt"
)

func FillSet(map_to_fil *map[string]bool, array_strs *[]string) {
	for _, item := range *array_strs {
		if !((*map_to_fil)[item]) {
			(*map_to_fil)[item] = true
		}
	}
}

func FillMultiSet(map_to_fil *map[string]int, array_strs *[]string) {
	for _, item := range *array_strs {
		(*map_to_fil)[item] += 1
	}
}

func main() {
	common_set := map[string]bool{}
	multi_set := map[string]int{}
	array_strs := []string{"cat", "cat", "dog", "cat", "tree"}
	FillMultiSet(&multi_set, &array_strs)
	FillSet(&common_set, &array_strs)
	fmt.Println(common_set)
	fmt.Println()
	fmt.Println(multi_set)
}
