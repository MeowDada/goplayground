package main

import (
	"fmt"
	"sync"
)

type MyPool struct {
	pool sync.Pool
}

func New(count int) *MyPool {
	
	i := count

	mp := &MyPool{
		pool: sync.Pool{
			New: func() interface{} {
				i++
				return i
			},
		},
	}

	for i := 0 ; i < count ; i++ {
		mp.Put(i)
	}

	return mp
}

func (mp *MyPool) Get() interface{} {
	return mp.pool.Get()
}

func (mp *MyPool) Put(something interface{}) {
	mp.pool.Put(something)
}

func main() {

	mp := New(5)

	for i := 0 ; i < 10 ; i++ {
		fmt.Println(mp.Get())
	}
}