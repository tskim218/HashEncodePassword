package datasource

import (
	"log"
	"sync"
	"time"
)

type workerStatus struct {
	shutDownStatus bool
	worker int
	lck sync.RWMutex
}

func WorkerStatus() Worker {
	return &workerStatus{shutDownStatus: false, worker: 0}
}

func (w *workerStatus) IncWorker() {
	w.lck.Lock()
	w.worker++
	w.lck.Unlock()
}

func (w *workerStatus) DecWorker() {
	w.lck.Lock()
	w.worker--
	w.lck.Unlock()
}

func (w *workerStatus) GetWorker() int {
	w.lck.Lock()
	numOfWorker := w.worker
	w.lck.Unlock()

	return numOfWorker
}

func (w *workerStatus) ChangeShutDownStatus() error {
	w.lck.Lock()
	w.shutDownStatus = true
	w.lck.Unlock()

	return nil
}

func (w *workerStatus) ShutDownStatus() (bool, error) {
	w.lck.Lock()
	status := w.shutDownStatus
	w.lck.Unlock()

	return status, nil
}
func (w *workerStatus) WaitingForWorkersDone() error {
	for {
		w.lck.Lock()
		status := w.worker
		log.Printf("Waiting for workers done: %d\n", status)
		if status == 0 {
			w.lck.Unlock()
			break
		}
		w.lck.Unlock()

		time.Sleep(5*time.Second)
	}
	log.Printf("Workers done")

	return nil
}
