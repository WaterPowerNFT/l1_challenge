package main

import (
	"fmt"
)

func SetIntersection(first_set *map[string]bool, second_set *map[string]bool) map[string]bool {
	inter_set := map[string]bool{}
	for key, _ := range *first_set {
		if (*second_set)[key] {
			inter_set[key] = true
		}
	}
	return inter_set
}

func main() {
	set1 := map[string]bool{"John": true, "Mary": true, "Bob": true, "Dondo": true}
	set2 := map[string]bool{"Mary": true, "Vasya": true, "Dondo": true}
	intersection_sets := SetIntersection(&set1, &set2)
	fmt.Println(intersection_sets)
}
