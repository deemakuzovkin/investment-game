package cached

import (
	"fmt"
	"git.mills.io/prologic/bitcask"
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
