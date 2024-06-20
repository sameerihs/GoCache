package cache

import (
	"container/list"
	"sync"
)

type Pair struct {
	key   string
	value interface{}
}

type LRUCache struct {
	capacity int
	list     *list.List
	exists   map[string]*list.Element
	mu       sync.RWMutex
}

// NewLRUCache initializes a new LRU cache with the given capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		exists:   make(map[string]*list.Element),
	}
}

// Get retrieves the value of the given key from the cache
func (c *LRUCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if node, ok := c.exists[key]; ok {
		val := node.Value.(Pair).value
		c.list.MoveToFront(node)
		return val, nil
	}

	return nil, ErrKeyNotFound
}

// Set inserts a key-value pair into the cache, evicting the least recently used item if necessary
func (c *LRUCache) Set(key string, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.exists[key]; ok {
		// If the key already exists, update its value and move it to the front
		c.list.MoveToFront(node)
		node.Value = Pair{key: key, value: value}
		return nil
	}

	// Evict the least recently used item if the cache is at capacity
	if c.list.Len() == c.capacity {
		idx := c.list.Back().Value.(Pair).key
		delete(c.exists, idx)
		c.list.Remove(c.list.Back())
	}

	// Insert the new item at the front
	ptr := c.list.PushFront(Pair{
		key:   key,
		value: value,
	})
	c.exists[key] = ptr

	return nil
}

// Delete removes the key-value pair associated with the given key from the cache
func (c *LRUCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.exists[key]; ok {
		delete(c.exists, key)
		c.list.Remove(node)
		return nil
	}

	return ErrKeyNotFound
}
