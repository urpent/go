package cachex

import (
	"time"

	"github.com/urpent/go/linklist"
)

type Options struct {
	MaxTimeToLive time.Duration
}

type ValueWithExpiry[V any] struct {
	Value     V
	KeyExpiry time.Time
}

// LRUCacheExpiry is Least Recently Used (LRU) cache.
type LRUCacheExpiry[K, V any] struct {
	maxTimeToLive time.Duration
	LRUCache[K, ValueWithExpiry[V]]
}

func NewLRUCacheExpiry[K, V any](capacity int, options Options) *LRUCacheExpiry[K, V] {
	return &LRUCacheExpiry[K, V]{
		maxTimeToLive: options.MaxTimeToLive,
		LRUCache: LRUCache[K, ValueWithExpiry[V]]{
			capacity: capacity,
			cacheMap: make(map[any]*linklist.Node[keyValuePair[K, ValueWithExpiry[V]]], capacity+1),
		},
	}
}

func (c *LRUCacheExpiry[K, V]) Set(key K, value V) {
	c.mux.Lock()
	defer c.mux.Unlock()

	item, existed := c.cacheMap[key]
	if existed {
		item.Data.Value.Value = value
		item.Data.Value.KeyExpiry = time.Now().Add(c.maxTimeToLive)
		c.linkedList.MoveNodeToFront(item)
		return
	}

	valueWithExpiry := ValueWithExpiry[V]{
		KeyExpiry: time.Now().Add(c.maxTimeToLive),
		Value:     value,
	}

	node := c.linkedList.AddFront(keyValuePair[K, ValueWithExpiry[V]]{Key: key, Value: valueWithExpiry})
	c.cacheMap[key] = node

	// If max capacity reached
	if len(c.cacheMap) > c.capacity {
		node := c.linkedList.RemoveLast()
		delete(c.cacheMap, node.Data.Key)
	}
}

func (c *LRUCacheExpiry[K, V]) Get(key K) (value V, ok bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	var item *linklist.Node[keyValuePair[K, ValueWithExpiry[V]]]
	item, ok = c.cacheMap[key]
	if !ok {
		return
	}

	// if key expired
	if time.Now().After(item.Data.Value.KeyExpiry) {
		ok = false
		return
	}

	c.linkedList.MoveNodeToFront(item)

	return item.Data.Value.Value, true
}
