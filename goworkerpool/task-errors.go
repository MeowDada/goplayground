package workerpool

type TaskResultErrorCode int
type TaskResultErrorCodeMap map[TaskResultErrorCode] TaskResultError

const (
	ErrEmptyTask        TaskResultErrorCode = iota
	ErrTaskNotImplement
	ErrNonExistErrorCode
	ErrTaskTimeOut
	ErrJobFailed
)

var taskResultErrorCodeMap = TaskResultErrorCodeMap {
	ErrEmptyTask : TaskResultError{
		errorType:   ErrorTypeIgnore,
		description: "There is no job of this task",
	},
	ErrTaskNotImplement : TaskResultError {
		errorType:   ErrorTypeNormal,
		description: "There are some functions in this task not yet been implemented",
	},
	ErrNonExistErrorCode : TaskResultError {
		errorType:   ErrorTypeInternal,
		description: "The error code dose not exist",
	},
	ErrTaskTimeOut : TaskResultError {
		errorType:   ErrorTypeNormal,
		description: "The task has exceed exection time limit",
	},
	ErrJobFailed : TaskResultError {
		errorType:   ErrorTypeNormal,
		description: "A job failed so the task has been halt and exited",
	},
}

func ErrorCodeToResult(code TaskResultErrorCode) TaskResultError {
	if r, ok := taskResultErrorCodeMap[code]; ok {
		return r
	}
	return taskResultErrorCodeMap[ErrNonExistErrorCode]
}

