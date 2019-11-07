package workerpool

type JobSelector interface {
	Len() int
	Push(Job)
	Pop() Job
}





