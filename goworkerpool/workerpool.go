package workerpool

type WorkerPool struct {
	maxWorkers int
	interval   int
	processingChannel WorkerpoolChannel
	resultChannel     ResultChannel
}