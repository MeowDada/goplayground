package main

import (
	"fmt"
	"os"
	"syscall"
	"sync"
	"time"
)

func main() {

	var (
		numGorouintes = 20
		testFile = "test-lock"
	)

	// Create test file
	os.Create(testFile)

	// Use to prevent main function return before not all go routines are finished.
	var wg sync.WaitGroup

	for i := 0 ; i < numGorouintes; i++ {
		wg.Add(1)

		// This routines try open a file and flock it. Print error message if failed.
		go func(id int) {

			// Decrement the semaphore when go routine return.
			defer wg.Done()

			f, err := os.Open(testFile)
			if err != nil {
				fmt.Printf("Goroutine<%d>: failed to open file: %s\n", id, err)
				return
			}

			// Try flocking the file with exclusive mode and non-blocking mode.
			err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX | syscall.LOCK_NB)
			if err != nil {
				fmt.Printf("Goroutine<%d>: failed to flock: %s\n", id, err)
				return
			}
			fmt.Printf("Goroutine<%d>: Flock succeed with fd = %d\n", id, int(f.Fd()))

			// Unlock the file explicitly. ( Close the file descriptor could also unlock. but it is not recommended. )
			defer syscall.Flock(int(f.Fd()), syscall.LOCK_UN)

			// Make it sleep for a while to make others go routine harder to flock the file.
			time.Sleep(time.Duration(1000)*time.Millisecond)

		}(i)
	}

	// Wait for all go routines to finish.
	wg.Wait()
}