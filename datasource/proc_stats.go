package datasource

import "sync"

type procStats struct {
	total int64
	totalProcessedTime int64
	lck sync.RWMutex
}

func ProcStats() Stats {
	return &procStats{total: 0, totalProcessedTime: 0}
}

func (ps *procStats) AddWorker() {
	ps.lck.Lock()
	ps.total++
	ps.lck.Unlock()
}

func (ps *procStats) GetTotal() int64 {
	ps.lck.Lock()
	total := ps.total
	ps.lck.Unlock()

	return total
}

func (ps *procStats) AddTime(procTime int64) {
	ps.lck.Lock()
	ps.totalProcessedTime += procTime
	ps.lck.Unlock()
}

func (ps *procStats) GetTotalTime() int64 {
	ps.lck.Lock()
	procTime := ps.totalProcessedTime
	ps.lck.Unlock()

	return procTime
}