package workerpool

import (
	"time"
)

type JobTestObj struct {
	id            JobID
	createTiming  time.Time
	startTiming   time.Time
	endTiming     time.Time
	executionTime time.Duration
	priority      int
}

func NewJobTestObj(id JobID, priority int) JobTestObj {
	return JobTestObj {
		id: id,
		priority: priority,
	}
}

func (j *JobTestObj) ID() JobID {
	return j.id
}

func (j *JobTestObj) CreateTiming() time.Time {
	return j.createTiming
}

func (j *JobTestObj) StartTiming() time.Time {
	return j.startTiming
}

func (j *JobTestObj) EndTiming() time.Time {
	return j.endTiming
}

func (j *JobTestObj) ExecutionTime() time.Duration {
	return j.executionTime
}

func (j *JobTestObj) Execute() JobResult {
	return JobResult{}
}

func (j *JobTestObj) Abort() {
	return
}

func (j *JobTestObj) Priority() int {
	return j.priority
}

func (j *JobTestObj) Compare(other JobTestObj) bool {
	return j.id == other.id && j.priority == other.priority && j.createTiming == other.createTiming &&
		j.startTiming == other.startTiming && j.endTiming == other.endTiming && j.executionTime == other.executionTime
}