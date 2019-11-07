package workerpool

// JobMaxHeap implements heap.Interface that could be used as a dependency injection
// target to a JobPriorityQueue struct. 
type JobMaxHeap []Job
func (h JobMaxHeap) Len() int { return len(h) }
func (h JobMaxHeap) Less(i, j int) bool { return h[i].Priority() > h[j].Priority() }
func (h JobMaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *JobMaxHeap) Push(x interface{}) {
	job := x.(Job)
	*h = append(*h, job)
}

func (h *JobMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*h = old[0 : n-1]
	return item
}
