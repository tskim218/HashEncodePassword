package datasource

import "errors"

var (
	ErrNotFound = errors.New("id not found")
)

// DB is the interface to a simple key/value store
type DB interface {
	// Get returns the value for the given key, ErrNotFound if the id doesn't exist,
	// or another error if the get failed
	// Get(key string) ([]byte, error)
	Get(Id uint64) (string, error)
	// Set sets the value for the given key. Returns an error if the set failed.
	// If non-nil error is returned, the value was not updated
	// Set(key string, val []byte) error
	Set(Id uint64, encode string) error

	// Create an id, increment by 1 and returns it.
	GetId() uint64
}