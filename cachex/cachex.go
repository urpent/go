package cachex

type Cacher[K, V any] interface {
	Set(key K, value V)
	Get(key K) (value V, ok bool)
	Delete(key K)
	Clear()
	Len() int
}
