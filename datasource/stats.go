package datasource

type Stats interface {
	// Add hash handler worker
	AddWorker()

	// Get the total number of hash handler worker
	GetTotal() int64

	// Add process time of the worker
	AddTime(procTime int64)

	// Get process time
	GetTotalTime() int64
}
