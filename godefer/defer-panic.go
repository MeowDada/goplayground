package main

import "fmt"

// defered function will still be executed even if panic occurs
func main() {
	
	divisor := 0

	defer fmt.Println("Leaving")

	fmt.Println(5/divisor)
}