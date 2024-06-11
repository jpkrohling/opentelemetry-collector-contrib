package cache

type Cache[T comparable, V any] interface {
	// Get returns the value for the given key if it exists in the cache.
	Get(key T) (V, bool)
	// Set sets the value for the given key.
	Set(key T, value V)
	// Delete deletes the value for the given key.
	Delete(key T)
}
