package repository

import "sync"

var once sync.Once
type singleton map[string]bool

var (
	instance singleton
)

// NewRepository Singleton that simulates a database or external process
func NewRepository() singleton{
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}