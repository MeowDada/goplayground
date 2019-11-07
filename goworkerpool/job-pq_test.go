package workerpool

import (
	"testing"
)

func TestPQPush(t *testing.T) {

	pq := NewJobPriorityQueue()

	obj1 := NewJobTestObj(1,3)
	obj2 := NewJobTestObj(2,3)
	obj3 := NewJobTestObj(3,5)
	obj4 := NewJobTestObj(3,1)
	obj5 := NewJobTestObj(3,4)
	obj6 := NewJobTestObj(30,500)
	obj7 := NewJobTestObj(20,300)
	obj8 := NewJobTestObj(10,10)

	pq.Push(&obj1)
	pq.Push(&obj2)
	pq.Push(&obj3)
	pq.Push(&obj4)
	pq.Push(&obj5)
	pq.Push(&obj6)
	pq.Push(&obj7)
	pq.Push(&obj8)

	want := obj6.Priority()
	get := pq.Pop().Priority()

	if want != get {
		t.Errorf("expect %v, but get %v", want, get)
	}
}

func TestPQPop(t *testing.T) {

	pq := NewJobPriorityQueue()

	obj1 := NewJobTestObj(1,3)
	obj2 := NewJobTestObj(2,3)
	obj3 := NewJobTestObj(3,5)
	obj4 := NewJobTestObj(3,1)
	obj5 := NewJobTestObj(3,4)
	obj6 := NewJobTestObj(30,500)
	obj7 := NewJobTestObj(20,300)
	obj8 := NewJobTestObj(10,10)

	pq.Push(&obj1)
	pq.Push(&obj2)
	pq.Push(&obj3)
	pq.Push(&obj4)
	pq.Push(&obj5)
	pq.Push(&obj6)
	pq.Push(&obj7)
	pq.Push(&obj8)

	pq.Pop()
	pq.Pop()

	want := 6
	get := pq.Len()

	if want != get {
		t.Errorf("expect %v, but get %v", want, get)
	}
}

func TestPQLen(t *testing.T) {

	pq := NewJobPriorityQueue()

	want := 0
	get := pq.Len()

	if want != get {
		t.Errorf("expect %v, but get %v", want, get)
	}
}