package datasource

import (
	"log"
	"sync"
	"time"
)

type inMemoryDB struct {
	id uint64
	m map[uint64] string
	lck sync.RWMutex
}

// NewInMemoryDB creates a new DB implementation that stores all data in memory.
// All operations are concurrency safe
func NewInMemoryDB() DB {
	return &inMemoryDB{id: 0,m: make(map[uint64]string)}
}

// Get is the interface implementation
func (d *inMemoryDB) Get(id uint64) (string, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()
	v, ok := d.m[id]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

// Set is the interface implementation
func (d *inMemoryDB) Set(id uint64, encode string) error {
	time.Sleep(5 * time.Second)
	d.lck.Lock()
	defer d.lck.Unlock()
	d.m[id] = encode
	log.Printf("id: %d, encode: %s", id, d.m[id])
	return nil
}

func (d *inMemoryDB) GetId() uint64 {
	d.lck.Lock()
	d.id++
	d.lck.Unlock()
	return d.id
}