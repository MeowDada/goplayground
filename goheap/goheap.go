package main

import (
	"container/heap"
	"fmt"
)

type IntMinHeap struct {
	arr []int
}

// Len implements sort.Interface
func (h *IntMinHeap) Len() int {
	return len(h.arr)
}

// Less implements sort.Interface
func (h *IntMinHeap) Less(i, j int) bool {
	return h.arr[i] < h.arr[j]
}

// Swap implements sort.Interface
func (h *IntMinHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// Push implements heap.Interface
func (h *IntMinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.arr = append(h.arr, x.(int))
}

// Pop implements heap.Interface
func (h *IntMinHeap) Pop() interface{} {
	old := (*h).arr
	n := len(old)
	x := old[n-1]
	(*h).arr = old[0:n-1]
	return x
}

// Any type that implements Interface interface may be used as a min-heap
func main() {
	h := &IntMinHeap{ 
		arr: []int{2,1,5,4,0},
	}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", h.arr[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}