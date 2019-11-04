package main

import (
	"fmt"
)

type resolver interface {
	Resolve()
}

// A foo type struct which implements the resolver interface
type foo struct {}

// overrided function from the interface, resolver
func (f foo) Resolve() {
	fmt.Println("foo: Resolve()")
}

func main() {

	var x interface{}

	x = &foo{}

	// Assert x implements resolver interface
	if v, ok := x.(resolver); ok {
		v.Resolve()
	}
}