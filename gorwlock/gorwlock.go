package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
	"strings"
)

type Writer struct {
	mu *sync.RWMutex
	ch chan int
	id int
	wg *sync.WaitGroup
}

func (w *Writer) DoStuff() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	w.mu.Lock()
	defer w.wg.Done()
	defer w.mu.Unlock()
	w.ch <- 1
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	w.ch <- -1
}

type Reader struct {
	mu *sync.RWMutex
	ch chan int
	id int
	wg *sync.WaitGroup
}

func (r *Reader) DoStuff() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	r.mu.RLock()
	defer r.wg.Done()
	defer r.mu.RUnlock()
	r.ch <- 1
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	r.ch <- -1
}

func main() {

	rand.Seed(time.Now().Unix())

	var rwlock sync.RWMutex

	numWriter := 2
	numReader := 5

	writerCh := make(chan int)
	readerCh := make(chan int)

	var wg sync.WaitGroup 

	// Setup printer
	go func(readerCh, writerCh chan int) {

		ws := 0
		rs := 0

		for {
			select {
			case n := <- readerCh:
				rs += n
			case n := <- writerCh:
				ws += n
			}
			fmt.Printf("%s%s\n", strings.Repeat("R", rs), strings.Repeat("W", ws))
		}
	}(readerCh, writerCh)

	// Setup writers
	for i := 0 ; i < numWriter ; i++ {
		wg.Add(1)
		go func(lock *sync.RWMutex, ch chan int, id int, wg *sync.WaitGroup){
			writer := Writer{lock, ch, id, wg}
			for {
				writer.DoStuff()
			}
		}(&rwlock, writerCh, i, &wg)
	}

	// Setup Readers
	for i := 0 ; i < numReader ; i++ {
		wg.Add(1)
		go func(lock *sync.RWMutex, ch chan int, id int, wg *sync.WaitGroup){
			reader := Reader{lock, ch, id, wg}
			for {
				reader.DoStuff()
			}
		}(&rwlock, readerCh, i, &wg)
	}

	wg.Wait()
}
