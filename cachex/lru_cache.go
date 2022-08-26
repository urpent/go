package cache

import (
	"sync"

	"github.com/urpent/go/linklist"
)

type keyValuePair[K, V any] struct {
	Key   K
	Value V
}

// lruCache is Least Recently Used (LRU) cache.
type lruCache[K, V any] struct {
	cacheMap   map[any]*linklist.Node[keyValuePair[K, V]]
	mux        sync.Mutex
	linkedList linklist.DoublyLinkedList[keyValuePair[K, V]]
	capacity   int
}

func NewLRUCache[K, V any](capacity int) lruCache[K, V] {
	return lruCache[K, V]{capacity: capacity,
		cacheMap: make(map[any]*linklist.Node[keyValuePair[K, V]], capacity+1),
	}
}

func (c *lruCache[K, V]) Set(key K, value V) {
	c.mux.Lock()
	defer c.mux.Unlock()

	item, existed := c.cacheMap[key]
	if existed {
		c.linkedList.MoveNodeToFront(item)
		return
	}

	node := c.linkedList.AddFront(keyValuePair[K, V]{Key: key, Value: value})
	c.cacheMap[key] = node

	// If max capacity reached
	if len(c.cacheMap) > c.capacity {
		node := c.linkedList.RemoveLast()
		delete(c.cacheMap, node.Data.Key)
	}

	return
}

func (c *lruCache[K, V]) Get(key K) (value V, ok bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	var item *linklist.Node[keyValuePair[K, V]]
	item, ok = c.cacheMap[key]
	if !ok {
		return
	}

	c.linkedList.MoveNodeToFront(item)

	return item.Data.Value, true
}
