package main

import (
	pool "github.com/MeowDada/goplayground/goworkers/pkg/pool"
	job  "github.com/MeowDada/goplayground/goworkers/pkg/job"
)

func main() {
	dispatcher := pool.StartDispatcher(10)

	for i, job := range job.CreateJobs(50) {
		dispatcher.WorkCh <- pool.Work{ Job: job, ID: i }
	}
}