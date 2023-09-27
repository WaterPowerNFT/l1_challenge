package main

import (
	"fmt"
	"reflect"
)

func main() {
	json_map := map[string]interface{}{"first": 1, "second": "str", "third": false}
	for key, value := range json_map {
		fmt.Printf("for key %s, value %v is type of %s\n", key, value, reflect.TypeOf(value))

	}
}
