package workerpool

import (
	"testing"
)

func TestJobContainerPush(t *testing.T) {
	
	pq := NewJobPriorityQueue()
	jc := NewJobContainer(&pq)

	obj1 := NewJobTestObj(1,3)
	obj2 := NewJobTestObj(2,3)
	obj3 := NewJobTestObj(3,5)
	obj4 := NewJobTestObj(3,1)
	obj5 := NewJobTestObj(3,10)
	obj6 := NewJobTestObj(30,500)
	obj7 := NewJobTestObj(20,300)
	obj8 := NewJobTestObj(10,4)

	jc.Push(&obj1)
	jc.Push(&obj2)
	jc.Push(&obj3)
	jc.Push(&obj4)
	jc.Push(&obj5)
	jc.Push(&obj6)
	jc.Push(&obj7)
	jc.Push(&obj8)

	want := obj6.Priority()
	get := jc.Pop().Priority()

	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}

	want = obj7.Priority()
	get = jc.Pop().Priority()

	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}

	want = obj5.Priority()
	get = jc.Pop().Priority()

	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}
}

func TestJobContainerLen(t *testing.T) {

	pq := NewJobPriorityQueue()
	jc := NewJobContainer(&pq)

	obj1 := NewJobTestObj(1,3)
	obj2 := NewJobTestObj(2,3)
	obj3 := NewJobTestObj(3,5)
	obj4 := NewJobTestObj(3,1)
	obj5 := NewJobTestObj(3,10)
	obj6 := NewJobTestObj(30,500)
	obj7 := NewJobTestObj(20,300)
	obj8 := NewJobTestObj(10,4)

	jc.Push(&obj1)
	jc.Push(&obj2)
	jc.Push(&obj3)
	jc.Push(&obj4)
	jc.Push(&obj5)
	jc.Push(&obj6)
	jc.Push(&obj7)
	jc.Push(&obj8)

	want := 8
	get := jc.Len()

	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}
}

func TestJobContainerPop(t *testing.T) {

	pq := NewJobPriorityQueue()
	jc := NewJobContainer(&pq)

	obj1 := NewJobTestObj(1,3)
	obj2 := NewJobTestObj(2,3)
	obj3 := NewJobTestObj(3,5)
	obj4 := NewJobTestObj(3,1)
	obj5 := NewJobTestObj(3,10)
	obj6 := NewJobTestObj(30,500)
	obj7 := NewJobTestObj(20,300)
	obj8 := NewJobTestObj(10,4)

	jc.Push(&obj1)
	jc.Push(&obj2)
	jc.Push(&obj3)
	jc.Push(&obj4)
	jc.Push(&obj5)
	jc.Push(&obj6)
	jc.Push(&obj7)
	jc.Push(&obj8)

	want := obj6.Priority()
	get := jc.Pop().Priority()

	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}

	want = obj7.Priority()
	get = jc.Pop().Priority()

	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}

	want = obj5.Priority()
	get = jc.Pop().Priority()

	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}
}