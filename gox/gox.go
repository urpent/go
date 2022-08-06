package gox

// Ptr return the pointer of the value
func Ptr[T any](value T) *T {
	return &value
}
