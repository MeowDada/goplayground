package pool

import (
	"fmt"
)

type Dispatcher struct {
	WorkCh chan Work
	QuitCh chan bool
}

var WorkerChannel = make(chan chan Work)

func StartDispatcher(workerCount int) Dispatcher {
	
	var workers []Worker
	input := make(chan Work)
	end := make(chan bool)
	dispatcher := Dispatcher { WorkCh:input, QuitCh:end }
	
	for i := 0; i < workerCount; i++ {
		fmt.Println("starting worker: ", i)
		worker := Worker {
			ID:       i,
			WorkerCh: WorkerChannel,
			WorkCh:   make(chan Work),
			QuitCh:   make(chan bool),
		}
		worker.Start()
		workers = append(workers, worker)
	}

	go func() {
		for {
			select {
			case <- end:
				for _, w := range workers {
					w.Stop()
				}
				return
			case work := <- input:
				worker := <- WorkerChannel
				worker <- work
			}
		}
	}()
	return dispatcher
}

