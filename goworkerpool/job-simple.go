package workerpool

import (
	"time"
)

type SimpleJob struct {
	id            JobID
	createTiming  time.Time
	startTiming   time.Time
	endTiming     time.Time
	function      SimpleFn
	parameters    []interface{}
	priority      int
	executed      bool
}

type SimpleFn func(...interface{}) (interface{}, error)

// NewSimpleJob creats a SimpleJob object, it will be executed in FIFO order.
// The JobID of this SimpleJob object will also be assigned automatically.
func NewSimpleJob(fn SimpleFn, args ...interface{}) SimpleJob {

	var parameters []interface{}
	
	for _, param := range args {
		parameters = append(parameters, param)
	}

	s := SimpleJob{
		createTiming: time.Now(),
		function:     fn,
		parameters:   parameters,
		priority:     0,
		executed:     false,
	}

	return s
}

func (s *SimpleJob) ID() JobID {
	return s.id
}

func (s *SimpleJob) CreateTiming() time.Time {
	return s.createTiming
}

func (s *SimpleJob) StartTiming() time.Time {
	return s.startTiming
}

func (s *SimpleJob) EndTiming() time.Time {
	return s.endTiming
}

func (s *SimpleJob) ExecutionTime() time.Duration {
	return s.endTiming.Sub(s.startTiming)
}

func (s *SimpleJob) Execute() JobResult {
	result, err := s.function(s.parameters...)
	return JobResult{
		job: s,
		result: result,
		err: err,
	}
}

func (s *SimpleJob) Abort() {

}

func (s *SimpleJob) Priority() int {
	return s.priority
}