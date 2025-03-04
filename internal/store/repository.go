package store

func (b *Bucket) Set(key string, value any) {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	b.Data[key] = value
}

func (b *Bucket) Get(key string) (any, bool) {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	val, ok := b.Data[key]
	return val, ok
}

func (b *Bucket) Delete(key string) {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	delete(b.Data, key)
}
