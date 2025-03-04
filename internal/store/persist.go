package store

import (
	"encoding/json"
	"io"
	"os"
)

func (b *Bucket) SaveToFile() error {
	b.Mu.RLock()
	defer b.Mu.RUnlock()

	file, err := os.Create("buckets/" + b.Name + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(b.Data); err != nil {
		return err
	}

	return nil
}

func (b *Bucket) LoadFromFile() error {
	b.Mu.Lock()
	defer b.Mu.Unlock()

	file, err := os.Open("buckets/" + b.Name + ".json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&b.Data); err != nil && err != io.EOF {
		return err
	}

	return nil
}
