package workerpool

import (
	"time"
)

type TaskID uint64
type TaskState int

const (
	TaskStateWait TaskState = iota
	TaskStateRun
	TaskStateSuceess
	TaskStateFailure
	TaskStatePause
	TaskStateExit
)

type Task interface {
	ID()                TaskID
	Description()       string
	StartTiming()       time.Time
	JoinTiming()        time.Time
	EndTiming()         time.Time
	Timeout()           time.Duration
	State()             TaskState
	AssignedWorkerID()  WorkerID
	Jobs()              JobContainer
	Abort()
	IgnoreJobFail()     bool
	Callback()          TaskCallback
	Priority()          int
}