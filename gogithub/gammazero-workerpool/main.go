package main

import (
	"fmt"

	"github.com/gammazero/workerpool"
)

func main() {
	wp := workerpool.New(2)
	requests := []string{"alpha","beta","gamma","delta","epsilon"}

	for _, req := range requests {
		r := req
		wp.Submit(func() {
			fmt.Println("Handling request:", r)	
		})
	}
	wp.StopWait()
}