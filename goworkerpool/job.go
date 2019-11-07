package workerpool

import (
	"time"
)

type JobID uint64

type Job interface {
	ID()            JobID
	CreateTiming()  time.Time
	StartTiming()   time.Time
	EndTiming()     time.Time
	ExecutionTime() time.Duration
	Execute()       JobResult
	Abort()
	Priority()      int
}