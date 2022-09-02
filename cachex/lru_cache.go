package cachex

import (
	"sync"

	"github.com/urpent/go/linklist"
)

type keyValuePair[K, V any] struct {
	Key   K
	Value V
}

// LRUCache is Least Recently Used (LRU) cache.
type LRUCache[K, V any] struct {
	cacheMap   map[any]*linklist.Node[keyValuePair[K, V]]
	mux        sync.Mutex
	linkedList linklist.DoublyLinkedList[keyValuePair[K, V]]
	capacity   int
}

func NewLRUCache[K, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		cacheMap: make(map[any]*linklist.Node[keyValuePair[K, V]], capacity+1),
	}
}

func (c *LRUCache[K, V]) Set(key K, value V) {
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
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
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

func (c *LRUCache[K, V]) Delete(key K) {
	c.mux.Lock()
	defer c.mux.Unlock()

	item, ok := c.cacheMap[key]
	if !ok {
		return
	}

	c.linkedList.Remove(item)
	delete(c.cacheMap, key)
}

func (c *LRUCache[K, V]) Clear() {
	cacheMap := make(map[any]*linklist.Node[keyValuePair[K, V]], c.capacity+1)
	linkedList := linklist.NewDoublyLinkedList[keyValuePair[K, V]]()

	c.mux.Lock()
	c.cacheMap = cacheMap
	c.linkedList = linkedList
	c.mux.Unlock()
}

func (c *LRUCache[K, V]) Len() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return len(c.cacheMap)
}
