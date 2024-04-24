package main

import (
	"fmt"
	"go-dispatcher/dispatcher"
	"log"
	"time"
)

// PrintJob defines a job type that simply prints a message.
type PrintJob struct {
	Message string
}

// Process implements the Job interface's Process method for PrintJob.
// This method is called by a worker when the job is dequeued.
func (p PrintJob) Process() error {
	fmt.Println(p.Message) // Output the message to the console.
	return nil             // Return no error.
}

func main() {
	// Create a new dispatcher with a capacity for 10 workers and a queue size of 50.
	goDispatcher := dispatcher.New(10, 50)
	goDispatcher.Start() // Start the dispatcher and all its workers.

	// Simulate job submissions to the jobQueue.
	for i := 1; i <= 100; i++ {
		goDispatcher.AddJob(PrintJob{Message: fmt.Sprintf("Processing job %d", i)})
	}

	// Wait for a specified duration before stopping the dispatcher.
	time.Sleep(10 * time.Second) // Wait for 10 seconds

	log.Println("Automatically stopping the dispatcher after 10 seconds")
	goDispatcher.Stop() // Stop the dispatcher and all its workers gracefully.
}
