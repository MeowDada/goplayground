package workerpool

type TaskResult struct {
	id      WorkerID
	results []JobResult
	err     TaskResultError
}

func (tr *TaskResult) Error() string {
	return tr.err.Error()
}

func (tr *TaskResult) ID() WorkerID {
	return tr.id
}

func (tr *TaskResult) Results() []JobResult {
	return tr.results
}

func (tr *TaskResult) Err() error {
	return &tr.err
}

func (tr *TaskResult) Append(result JobResult) {
	tr.results = append(tr.results, result)
}