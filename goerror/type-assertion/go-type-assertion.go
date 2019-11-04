package main

import (
	"fmt"
)

func main() {

	var x interface{}

	// panic: interface conversion: interface{} is nil, not string
	// s := x.(string)

	// dose not panic: it could handle the case if interface x cannot be convert to string
	if s, ok := x.(string); ok {
		fmt.Println(s)
		return
	}
}