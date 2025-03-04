package store

import (
	"sync"
)

type Database struct {
	Buckets []*Bucket
}

type Bucket struct {
	Name string
	Data map[string]any
	Mu   sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{Buckets: []*Bucket{}}
}

func NewBucket(name string) *Bucket {
	return &Bucket{
		Name: name,
		Data: make(map[string]any),
	}
}
