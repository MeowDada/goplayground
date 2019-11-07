package workerpool

import (
	"testing"
	"container/heap"
)

func TestLen(t *testing.T) {
	
	obj1 := NewJobTestObj(1,3)
	obj2 := NewJobTestObj(2,3)
	obj3 := NewJobTestObj(3,4)

	// normal case for length = 3
	h := JobMaxHeap{}
	heap.Init(&h)
	heap.Push(&h, &obj1)
	heap.Push(&h, &obj2)
	heap.Push(&h, &obj3)
	want := 3
	get := h.Len()
	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}

	// boundary case for length = 0
	h = JobMaxHeap{ }
	want = 0
	get = h.Len()
	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}
}

func TestPush(t *testing.T) {

	obj1 := NewJobTestObj(1,3)
	obj2 := NewJobTestObj(2,3)
	obj3 := NewJobTestObj(3,5)
	obj4 := NewJobTestObj(3,1)
	obj5 := NewJobTestObj(3,4)
	obj6 := NewJobTestObj(30,500)
	obj7 := NewJobTestObj(20,300)
	obj8 := NewJobTestObj(10,10)


	// normal case for length = 3
	h := JobMaxHeap{}
	heap.Init(&h)
	heap.Push(&h, &obj1)
	heap.Push(&h, &obj2)
	heap.Push(&h, &obj3)
	heap.Push(&h, &obj4)
	heap.Push(&h, &obj5)
	heap.Push(&h, &obj6)
	heap.Push(&h, &obj7)
	heap.Push(&h, &obj8)
	want := &obj6
	get := heap.Pop(&h)
	if get != want {
		t.Errorf("expect %v, but get %v", want, get)
	}
}

func TestPop(t *testing.T) {

	obj1 := NewJobTestObj(1,200)
	obj2 := NewJobTestObj(2,100)
	obj3 := NewJobTestObj(3,5)
	obj4 := NewJobTestObj(3,1)
	obj5 := NewJobTestObj(3,4)

	// normal case for length = 3
	h := JobMaxHeap{}
	heap.Init(&h)
	heap.Push(&h, &obj1)
	heap.Push(&h, &obj2)
	heap.Push(&h, &obj3)
	heap.Push(&h, &obj4)
	heap.Push(&h, &obj5)
	want := &obj1
	get := heap.Pop(&h)
	if want != get {
		t.Errorf("expect %v, but get %v", want, get)
	}
}