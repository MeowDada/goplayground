package workerpool

type JobResult struct {
	job     Job
	result  interface{}
	err     error
}

