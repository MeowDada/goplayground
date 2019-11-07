package workerpool

import (
	"time"
)

type WorkerID    uint32
type WorkerState int
type WorkerpoolChannel chan Task
type ResultChannel chan TaskResult

const (
	StateCreate  WorkerState = iota
	StateReady
	StateSleep   
	StateBusy
	StateIdle
	StateExiting
	StateExit
)

type Worker struct {
	id                WorkerID
	state             WorkerState
	interval          time.Duration
	processingChannel WorkerpoolChannel
	resultChannel     ResultChannel
	quitChannel       chan bool
}

func NewWorker(id WorkerID, interval time.Duration, processingChannel WorkerpoolChannel, resultChannel ResultChannel) Worker {
	return Worker {
		id:    id,
		state: StateCreate,
		interval: interval,
		processingChannel: processingChannel,
		resultChannel: resultChannel,
	}
}

// Do handles the task which is extracted from the job container of the task
func (w *Worker) Do(task Task) TaskResult {

	taskResult := TaskResult{ id: w.id, }

	// Pop job from task's job container and execute it
	container := task.Jobs()

	for container.Len() > 0 {
		job := container.Pop()
		result := job.Execute()

		// Append the job result to the task results
		taskResult.Append(result)
	}

	return taskResult
}

// Start starts a worker go routine
func (w *Worker) Start() {
	go w.WorkerRoutine()
}

// WorkerRoutine will keep trying to work when receving new Task from workerpool,
// not until it receiving exiting signal will it exit.
func (w *Worker) WorkerRoutine() {

	// inifinite for loop till receving exiting signal
	for {
		w.state = StateReady
		select {
		// While this worker receving new task from the worker pool
		case task := <-w.processingChannel:
			w.state = StateBusy
			result := w.Do(task)
			w.resultChannel<-result
			w.state = StateIdle

		// While this worker receving exiting signal from the worker pool
		case <-w.quitChannel:
			w.state = StateExiting
			defer func() {w.state = StateExit}()
			return

		// While this worker doesn't receiving any new signal from the worker pool,
		// make it sleep for a while to save the CPU usage
		default:
			w.state = StateSleep
			time.Sleep(w.interval)
		}
	}
}