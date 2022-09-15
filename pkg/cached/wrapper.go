package cached

import (
	"fmt"
	"git.mills.io/prologic/bitcask"
	"github.com/zenthangplus/goccm"
	"os"
	"sync"
)

var (
	MaxFileSize  = 1024 * 1024 * 100
	MaxKeySize   = uint32(1024 * 1024 * 5)
	MaxValueSize = uint64(1024 * 1024 * 50)
	DATABASE     = map[string]*DataCache{}
	mu           sync.Mutex
)

type DataCache struct {
	db *bitcask.Bitcask
}

type Item struct {
	Key   []byte
	Value []byte
}

// Connect init data service
func Connect(path string) (*DataCache, error) {
	mu.Lock()
	defer mu.Unlock()
	if DATABASE[path] == nil {
		tmpPath := fmt.Sprintf("%s/%s", os.TempDir(), path)
		base, err := bitcask.Open(tmpPath, bitcask.WithMaxDatafileSize(MaxFileSize), bitcask.WithMaxValueSize(MaxValueSize), bitcask.WithMaxKeySize(MaxKeySize), bitcask.WithAutoRecovery(true))
		if err != nil {
			return nil, err
		}
		DATABASE[path] = &DataCache{
			db: base,
		}
	}
	return DATABASE[path], nil
}

// Add object to data
func (data *DataCache) Add(key []byte, value []byte) error {
	return data.db.Put(key, value)
}

// Get object by key
func (data *DataCache) Get(key []byte) ([]byte, error) {
	return data.db.Get(key)
}

// Remove object by key
func (data *DataCache) Remove(key []byte) error {
	return data.db.Delete(key)
}

// GetList data
func (data *DataCache) GetList() (<-chan Item, error) {
	total := 0
	manager := goccm.New(10)
	channel := make(chan Item)
	go func(inputCache *DataCache) {
		defer close(channel)
		inputCache.db.Fold(func(key []byte) error {
			manager.Wait()
			total++
			go func(resp chan Item) {
				defer manager.Done()
				get, err := inputCache.db.Get(key)
				if err != nil {
					return
				}
				resp <- Item{
					Key:   key,
					Value: get,
				}
			}(channel)
			return nil
		})
		fmt.Printf("Total scan: %d\n", total)
		manager.WaitAllDone()
	}(data)
	return channel, nil
}
