package datasource

// Worker is to keep tracking the status of all workers and server
// like how many are running, what the status or server running or shutting down
type Worker interface {
	// When a worker is running, increment by 1
	IncWorker()

	// When the worker is done, decrement by 1
	DecWorker()

	// Change the shut down status to true.
	ChangeShutDownStatus() error

	// Get the status the shut down
	ShutDownStatus() (bool, error)

	// Wait until all worker are down during the shutting down process
	WaitingForWorkersDone() error
}