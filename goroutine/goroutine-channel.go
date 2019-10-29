package main

import (
	"fmt"
)

func GoRoutine(routineID int, quitCh chan bool, quitDoneCh chan bool) {
	
	fmt.Printf("goroutine<%d> start\n", routineID)
	<- quitCh

	fmt.Printf("goroutine<%d> ends\n", routineID)
	quitDoneCh <- true
}

func main() {

	gloablQuitCh := make(chan bool)
	globalQuitDoneCh := make(chan bool)
	numRoutines := 10
	
	for i := 0 ; i < numRoutines ; i++ {
		go GoRoutine(i, gloablQuitCh, globalQuitDoneCh)
	}
	
	for i := 0 ; i < numRoutines ; i++ {
		gloablQuitCh <- true
	}

	leftRoutine := 0
	for _ = range globalQuitDoneCh {
		leftRoutine += 1
		if leftRoutine == numRoutines {
			fmt.Println("All the goroutine has been closed!")
			return
		}
	}
}