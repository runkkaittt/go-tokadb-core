package store

import (
	"sync"
)

type Bucket struct {
	Name string
	Data map[string]any
	Mu   sync.RWMutex
}

func NewBucket(name string) *Bucket {
	return &Bucket{
		Name: name,
		Data: make(map[string]any),
	}
}
