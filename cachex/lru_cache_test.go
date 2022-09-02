package cachex

import (
	"testing"

	"github.com/urpent/go/ut"
)

func Test_LRUCache(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		cache := NewLRUCache[int, int](2)

		cache.Set(1, 1)
		cache.Set(2, 2) // cache is full after set 2
		ut.AssertEqual(t, len(cache.cacheMap), 2)

		result, _ := cache.Get(1)
		ut.AssertEqual(t, 1, result) // ok return 1
		cache.Set(3, 3)              // LRU key was 2, evicts key 2, cache is {3=3, 1=1}
		_, ok := cache.Get(2)        // 2 not found
		ut.AssertEqual(t, false, ok)
		ut.AssertEqual(t, len(cache.cacheMap), 2)

		// now cache is {3=3, 1=1}, 1 is LRU
		cache.Set(1, 1) // 3 become LRU {1=1, 3=3}

		cache.Set(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 1=1}
		_, ok = cache.Get(3)
		ut.AssertEqual(t, false, ok)

		result, _ = cache.Get(1)
		ut.AssertEqual(t, 1, result)

		result, _ = cache.Get(4)
		ut.AssertEqual(t, 4, result)

		ut.AssertEqual(t, len(cache.cacheMap), 2)
		ut.AssertEqual(t, cache.linkedList.Len(), 2)
	})
}

func Test_LRUCache_Delete_Clear(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		var cache Cacher[int, int] = NewLRUCache[int, int](3)

		cache.Set(1, 1)
		cache.Set(2, 2)
		cache.Set(3, 3)
		cache.Set(4, 4) // evicts key 1
		ut.AssertEqual(t, 3, cache.Len())

		cache.Delete(3)
		ut.AssertEqual(t, 2, cache.Len())

		cache.Clear()
		ut.AssertEqual(t, 0, cache.Len())
	})
}
