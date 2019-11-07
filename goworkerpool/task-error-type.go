package workerpool

type TaskResultErrorType int

const (
	ErrorTypeFatal  TaskResultErrorType = iota
	ErrorTypeInternal
	ErrorTypeIgnore
	ErrorTypeNormal 
)

type TaskResultErrorTypeMap map[TaskResultErrorType]string

var taskResultErrorTypeMap = TaskResultErrorTypeMap {
	ErrorTypeFatal:    "Fatal",
	ErrorTypeInternal: "Internal",
	ErrorTypeIgnore:   "Ignore",
	ErrorTypeNormal:   "Normal",
}

func ErrorTypeToString(t TaskResultErrorType) string {
	return taskResultErrorTypeMap[t]
}