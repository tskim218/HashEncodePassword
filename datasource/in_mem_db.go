package datasource

import (
	"log"
	"sync"
	"time"
)

type inMemoryDB struct {
	// m   map[string][]byte
	m map[string] string
	lck sync.RWMutex
}

// NewInMemoryDB creates a new DB implementation that stores all data in memory.
// All operations are concurrency safe
func NewInMemoryDB() DB {
	// return &inMemoryDB{m: make(map[string][]byte)}
	return &inMemoryDB{m: make(map[string]string)}
}

// Get is the interface implementation
// func (d *inMemoryDB) Get(key string) ([]byte, error) {
	func (d *inMemoryDB) Get(key string) (string, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()
	v, ok := d.m[key]
	if !ok {
		// return nil, ErrNotFound
		return "", ErrNotFound
	}
	return v, nil
}

// Set is the interface implementation
func (d *inMemoryDB) Set(key string, val string) error {
	time.Sleep(5 * time.Second)
	log.Printf("inserting %s", key)
	d.lck.Lock()
	defer d.lck.Unlock()
	d.m[key] = val
	return nil
}