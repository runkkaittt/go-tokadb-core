package store

import (
	"sync"
)

type Bucket struct {
	Name string
	Data map[string]any
	Mu   sync.RWMutex
}

type Store struct {
	Buckets map[string]*Bucket
	mu      sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		Buckets: make(map[string]*Bucket),
	}
}

func (s *Store) NewBucket(name string) *Bucket {
	s.mu.Lock()
	defer s.mu.Unlock()

	if bucket, exists := s.Buckets[name]; exists {
		return bucket
	}

	bucket := &Bucket{
		Name: name,
		Data: make(map[string]any),
	}
	s.Buckets[name] = bucket
	return bucket
}

func (s *Store) GetBucket(name string) *Bucket {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, bucket := range s.Buckets {
		if bucket.Name == name {
			return bucket
		}
	}
	return nil
}

func (s *Store) RemoveBucket(name string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.Buckets[name]; !exists {
		return false
	}

	delete(s.Buckets, name)
	return true
}
