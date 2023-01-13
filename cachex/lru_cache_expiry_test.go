package cachex

import (
	"testing"
	"time"

	"github.com/urpent/go/ut"
)

func Test_LRUCacheExpiry(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		cache := NewLRUCacheExpiry[int, int](2, Options{MaxTimeToLive: 1 * time.Minute})

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

		cache.Set(4, 40)
		result, _ = cache.Get(4)
		ut.AssertEqual(t, 40, result)

		ut.AssertEqual(t, len(cache.cacheMap), 2)
		ut.AssertEqual(t, cache.linkedList.Len(), 2)
	})
}

func Test_LRUCacheExpiry_Expired(t *testing.T) {
	t.Run("test cache expiry", func(t *testing.T) {
		cache := NewLRUCacheExpiry[int, int](3, Options{MaxTimeToLive: 30 * time.Millisecond})

		cache.Set(1, 1)
		cache.Set(2, 2)
		result, _ := cache.Get(1)
		ut.AssertEqual(t, 1, result) // ok, key 1 not expired yet, return 1

		time.Sleep(30 * time.Millisecond)
		_, ok := cache.Get(2)
		ut.AssertEqual(t, false, ok) // 2 is expired and should not be found

		cache.Set(2, 20)
		result, ok = cache.Get(2)
		ut.AssertEqual(t, 20, result)
		ut.AssertEqual(t, true, ok)
	})
}
