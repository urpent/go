package cache

import (
	"testing"

	"github.com/urpent/go/ut"
)

func Test_LRUCache(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		cache := NewLRUCache[int, int](2)

		cache.Set(1, 1)
		cache.Set(2, 2)

		ut.AssertEqual(t, len(cache.cacheMap), 2)

		result, _ := cache.Get(1)
		ut.AssertEqual(t, 1, result) // ok return 1
		cache.Set(3, 3)              // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
		_, ok := cache.Get(2)        // 2 not found
		ut.AssertEqual(t, false, ok)
		ut.AssertEqual(t, len(cache.cacheMap), 2)

		cache.Set(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
		_, ok = cache.Get(1)
		ut.AssertEqual(t, false, ok)

		result, _ = cache.Get(3)
		ut.AssertEqual(t, 3, result)

		result, _ = cache.Get(4)
		ut.AssertEqual(t, 4, result)

		ut.AssertEqual(t, len(cache.cacheMap), 2)
		ut.AssertEqual(t, cache.linkedList.Len(), 2)
	})
}
