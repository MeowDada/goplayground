package workerpool

import (
	"container/heap"
)

// JobPriorityQueue implements JobSelector interface, it could be applied to
// a JobContainer which supports picking a job by calling JobSelector.Pop() and
// appending a job by invoking JobSelector.Push(job).
type JobPriorityQueue struct {
	h JobMaxHeap
}

func NewJobPriorityQueue() JobPriorityQueue {
	h := JobMaxHeap{}
	heap.Init(&h)
	return JobPriorityQueue {
		h :h,
	}
}

func (pq *JobPriorityQueue) Push(job Job) {
	heap.Push(&pq.h, job)
}

func (pq *JobPriorityQueue) Pop() Job {
	return heap.Pop(&pq.h).(Job)
}

func (pq *JobPriorityQueue) Len() int {
	return pq.h.Len()
}