package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircularBufferCache(t *testing.T) {
	size := 5
	cache := NewCircularBufferCache[int, string](size)

	// Assert that the buffer and values map are initialized correctly
	assert.NotNil(t, cache.buffer)
	assert.NotNil(t, cache.values)
	assert.Equal(t, size, len(cache.buffer))
	assert.Empty(t, cache.values)

	// Assert that the index is initialized to 0
	assert.Equal(t, 0, cache.index)
}

func TestCircularBufferCache_Get(t *testing.T) {
	cache := NewCircularBufferCache[int, string](5)

	// Add some values to the cache
	cache.Set(1, "one")
	cache.Set(2, "two")
	cache.Set(3, "three")

	// Test getting existing keys
	value, ok := cache.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "one", value)

	value, ok = cache.Get(2)
	assert.True(t, ok)
	assert.Equal(t, "two", value)

	value, ok = cache.Get(3)
	assert.True(t, ok)
	assert.Equal(t, "three", value)

	// Test getting non-existing key
	value, ok = cache.Get(4)
	assert.False(t, ok)
	assert.Equal(t, "", value)
}
func TestCircularBufferCache_Set(t *testing.T) {
	cache := NewCircularBufferCache[int, string](3)

	// Add some values to the cache
	cache.Set(1, "one")
	cache.Set(2, "two")
	cache.Set(3, "three")

	// Assert that the values are set correctly
	value, ok := cache.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "one", value)

	value, ok = cache.Get(2)
	assert.True(t, ok)
	assert.Equal(t, "two", value)

	value, ok = cache.Get(3)
	assert.True(t, ok)
	assert.Equal(t, "three", value)

	// Add another value, which should replace the oldest value in the buffer
	cache.Set(4, "four")

	// Assert that the oldest value is replaced
	value, ok = cache.Get(1)
	assert.False(t, ok)
	assert.Equal(t, "", value)

	// Assert that the new value is set correctly
	value, ok = cache.Get(4)
	assert.True(t, ok)
	assert.Equal(t, "four", value)
}
