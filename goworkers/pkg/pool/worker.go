package pool

import (
	"fmt"
	job "github.com/MeowDada/goplayground/goworkers/pkg/job"
)

type Work struct {
	ID  int
	Job string
}

type Worker struct {
	ID        int
	WorkerCh  chan chan Work
	WorkCh    chan Work
	QuitCh    chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerCh <- w.WorkCh
			select {
			case work := <- w.WorkCh:
				job.DoWork(work.Job, work.ID)
			case <- w.QuitCh:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	fmt.Print("worker<%d> is stopping\n", w.ID)
	w.QuitCh <- true
}