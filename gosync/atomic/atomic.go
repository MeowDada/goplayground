package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	numRoutines := 100000

	var value uint64 = 10

	var wg sync.WaitGroup

	fn := func(idx int) {
		if atomic.LoadUint64(&value) == 0 {
			atomic.AddUint64(&value, 1)
		}
		if atomic.AddUint64(&value, ^uint64(0)) == 0{
			val := atomic.LoadUint64(&value)
			if val == 1 {
				fmt.Println(idx, val)
			}
		}
		wg.Done()
	}

	for i := 0 ; i < numRoutines ; i++ {
		wg.Add(1)
		go fn(i)
	}
	wg.Wait()

}