package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cm   map[string]cacheEntry
	mu   sync.Mutex
	tick *time.Ticker
}
