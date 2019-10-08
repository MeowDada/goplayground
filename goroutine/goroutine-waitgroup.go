package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type routine struct {
	id       int        // routine ID
	duration int        // sleeping duration of this routine
}

// sleepRoutine makes this go routine sleep for a duration time in range [ 1, maxDuration ]
func sleepRoutine(id int, wg *sync.WaitGroup, timeCh chan routine, maxDuration int) {
	
	// notify the wait group when return
	defer wg.Done()

	// generate a random number and used it as sleep time
	// the sleeping duraction will be located at range [ 1, maxDuration ]
	randomNumber := rand.Intn(maxDuration)
	duration := time.Duration(randomNumber + 1)*time.Second

	// make this go routine sleep for a duration time
	time.Sleep(duration)

	// notify the channel that this routine has been waken up
	timeCh <- routine{ id:id, duration:randomNumber+1 }
}

// broadcastRoutine keeps brodcasting whenever any sleep routine wakes up
func broadcastRoutine(timeCh chan routine, quitCh chan bool) {
	for {
		select {

		// whenever the time channel recieve the wake up signal
		case routine := <- timeCh:
			fmt.Printf("routine<%d> has been slept for %d seconds\n", routine.id, routine.duration)

		// whenever the quit channel recieve the quit signal
		case <-quitCh:
			fmt.Println("closing brodcast routine")
			return
		}
	}
}

// main will generate serveral sleep routines and ensure that it will not exit
// until all of the sleep routines are done. But it is not guaranteed the boardcast
// routine will be finished normally before the program exited
func main() {

	// using current time as a random seed to ensure whenever we execute
	// this program will get different random results
	rand.Seed(time.Now().UTC().UnixNano())

	// use sync.WaitGroup to synchornize the go routines
	var wg sync.WaitGroup

	maxSleepDuration := 5
	maxRoutineNumber := 10

	// initialize the time channel
	timeCh := make(chan routine)
	defer close(timeCh)
	
	// initialize the quit channel
	quitCh := make(chan bool)
	defer close(quitCh)

	// start the broadcast routine
	go broadcastRoutine(timeCh, quitCh)

	// start adding sleep routines
	for i := 0 ; i < maxRoutineNumber; i++ {

		// waitGroup.Add(1) will increment the counting value by 1
		wg.Add(1)

		// start a sleep routine. it will invoke wg.Done() when it return,
		// which can help us to decrement the counting value by 1
		go sleepRoutine(i, &wg, timeCh, maxSleepDuration)
	}

	// waitGroup.Wait() will block current thread until the counting value is equal to zero
	wg.Wait()

	// after all of the sleep routines are done, notify the quit channel to 
	// close the broadcast routine
	quitCh <- true
	
	fmt.Println("All routines are finished")
}