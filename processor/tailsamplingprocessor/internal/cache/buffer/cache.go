// Package buffer implements a cache backed by a configurable circular buffer.
package buffer

import (
	"sync"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/cache"
)

var _ cache.Cache[int, string] = &CircularBufferCache[int, string]{}

type CircularBufferCache[T comparable, V any] struct {
	// The buffer is a circular buffer that holds the most recent items.
	buffer []T
	// The values map holds the values for the keys in the buffer.
	values map[T]V
	// The index is the current index in the buffer.
	index int
	// The mutex is used to synchronize access to the buffer and values map.
	mutex sync.RWMutex
}

// NewCircularBufferCache creates a new CircularBufferCache with the given size.
func NewCircularBufferCache[T comparable, V any](size int) *CircularBufferCache[T, V] {
	return &CircularBufferCache[T, V]{
		buffer: make([]T, size),
		values: make(map[T]V),
	}
}

// Get returns the value for the given key if it exists in the cache.
func (c *CircularBufferCache[T, V]) Get(key T) (V, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, ok := c.values[key]
	return value, ok
}

// Set sets the value for the given key.
func (c *CircularBufferCache[T, V]) Set(key T, value V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.index = (c.index + 1) % len(c.buffer)
	if _, ok := c.values[c.buffer[c.index]]; ok {
		c.Delete(c.buffer[c.index])
	}

	c.buffer[c.index] = key
	c.values[key] = value
}

// Delete deletes the value for the given key.
func (c *CircularBufferCache[T, V]) Delete(key T) {
	delete(c.values, key)
}
