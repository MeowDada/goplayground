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
	TimeOut()           time.Time
	State()             TaskState
	AssignedWorkerID()  WorkerID
	Jobs()              JobContainer
	Callback()          TaskCallback
	Priority()          int
}