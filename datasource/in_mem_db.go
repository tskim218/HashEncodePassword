package datasource

import (
	"log"
	"sync"
	"time"
)

type inMemoryDB struct {
	// m   map[string][]byte
	inc uint64
	m map[uint64] string
	lck sync.RWMutex
}

// NewInMemoryDB creates a new DB implementation that stores all data in memory.
// All operations are concurrency safe
func NewInMemoryDB() DB {
	// return &inMemoryDB{m: make(map[string][]byte)}
	return &inMemoryDB{inc: 0,m: make(map[uint64]string)}
}

// Get is the interface implementation
// func (d *inMemoryDB) Get(key string) ([]byte, error) {
//	func (d *inMemoryDB) Get(key string) (string, error) {
func (d *inMemoryDB) Get(key uint64) (string, error) {
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
func (d *inMemoryDB) Set(key uint64, val string) error {
	time.Sleep(5 * time.Second)
	log.Printf("inserting %s", key)
	d.lck.Lock()
	defer d.lck.Unlock()
	//d.inc++
	d.m[key] = val
	log.Printf("key: %d, val: %s", key, d.m[key])
	return nil
}

func (d *inMemoryDB) GetId() uint64 {
	d.lck.Lock()
	d.inc++
	d.lck.Unlock()
	return d.inc
}