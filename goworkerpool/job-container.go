package workerpool

type JobContainer struct {
	selector JobSelector
}

func NewJobContainer(selector JobSelector) JobContainer {
	return JobContainer {
		selector: selector,
	}
}

func (jc *JobContainer) Len() int {
	return jc.selector.Len()
}

func (jc *JobContainer) Push(job Job) {
	jc.selector.Push(job)
}

func (jc *JobContainer) Pop() Job {
	return jc.selector.Pop()
}

func (jc *JobContainer) Set(selector JobSelector) {
	jc.selector = selector
}