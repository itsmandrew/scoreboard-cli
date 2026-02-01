package sports

import (
	"sync"
	"time"
)

// Cache is a generic thread-safe cache with TTL support
type Cache[T any] struct {
	mu       sync.RWMutex
	data     T
	timestamp time.Time
	ttl      time.Duration
}

// NewCache creates a new cache with the specified TTL
func NewCache[T any](ttl time.Duration) *Cache[T] {
	return &Cache[T]{
		ttl: ttl,
	}
}

// Get returns cached data if it's still valid
func (c *Cache[T]) Get() (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var zero T
	// If timestamp is zero, cache has never been set
	if !c.timestamp.IsZero() && time.Since(c.timestamp) < c.ttl {
		return c.data, true
	}
	return zero, false
}

// Set updates the cache with fresh data
func (c *Cache[T]) Set(data T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = data
	c.timestamp = time.Now()
}
