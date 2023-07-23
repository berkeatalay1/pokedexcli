package pokecache

import (
	"sync"
	"time"
)

// CacheEntry -
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache -
type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.Mutex
}
