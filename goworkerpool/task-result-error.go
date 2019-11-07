package workerpool

import (
	"fmt"
	"time"
)

const (
	TaskResultErrorTimeFormat = time.RFC1123
)

type TaskResultError struct {
	errorType   TaskResultErrorType 
	description string
	where       string
	when        time.Time
	err         error
}

func NewTaskResultError(code TaskResultErrorCode, where string, err error) TaskResultError {
	e := ErrorCodeToResult(code)
	e.err = err
	e.where = where
	e.when = time.Now()
	return e
}

func (e *TaskResultError) Error() string {
	return fmt.Sprintf(`Error type: %s
	Location: %s
	Time: %s
	Description: %s
	Error message: %s`, ErrorTypeToString(e.errorType), e.where, e.when.Format(TaskResultErrorTimeFormat), e.description, e.err)
}