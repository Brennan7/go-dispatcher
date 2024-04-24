package dispatcher

import (
	"log"
)

// worker represents a concurrent worker that processes jobs.
type worker struct {
	id         int           // Unique identifier for the worker.
	workerPool chan chan Job // A channel on which this worker will send its job channel when ready.
	jobChannel chan Job      // A channel for receiving jobs to process.
}

// newWorker initializes a new instance of a Worker.
func newWorker(id int, workerPool chan chan Job) worker {
	return worker{
		id:         id,
		workerPool: workerPool,
		jobChannel: make(chan Job), // Initialize the job channel for receiving jobs.
	}
}

// start begins the worker's loop in a new goroutine, waiting for jobs and quit signals.
func (w worker) start() {
	go func() {
		for {
			// The worker registers its jobChannel to the workerPool, signaling it's ready for work.
			// This operation can block if the pool is currently full and no dispatcher is taking channels out.
			w.workerPool <- w.jobChannel

			select {
			case job, ok := <-w.jobChannel:
				if !ok {
					log.Printf("Worker %d stopping due to jobChannel closure\n", w.id)
					return
				}
				// Once a job is received, the worker is now processing a job.
				log.Printf("Worker %d started job\n", w.id)
				// Process the job using the job's Process method which must be implemented by the job type.
				if err := job.Process(); err != nil {
					log.Printf("Worker %d error processing job: %s\n", w.id, err)
				}
				log.Printf("Worker %d finished job\n", w.id)
			}
		}
	}()
}
