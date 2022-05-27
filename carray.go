package mi

import (
	"fmt"
	"reflect"
	"unsafe"
)

type CArray[T any] struct {
	ptr  unsafe.Pointer
	len  int
	size int
}

func (c *CArray[T]) Length() int {
	return c.len
}

func (c *CArray[T]) At(idx int) (T, error) {
	if idx >= c.len {
		var zero T
		return zero, fmt.Errorf("CArray.At: %w: %d", errOutOfRange, idx)
	}
	return *(*T)(unsafe.Pointer(uintptr(c.ptr) + uintptr(idx)*uintptr(c.size))), nil
}

func (c *CArray[T]) Set(idx int, value T) error {
	if idx >= c.len {
		return fmt.Errorf("CArray.Set: %w: %d", errOutOfRange, idx)
	}
	*(*T)(unsafe.Pointer(uintptr(c.ptr) + uintptr(idx)*uintptr(c.size))) = value
	return nil
}

func (c *CArray[T]) ToSlice() []T {
	result := []T{}
	(*reflect.SliceHeader)(unsafe.Pointer(&result)).Data = uintptr(c.ptr)
	(*reflect.SliceHeader)(unsafe.Pointer(&result)).Len = c.len
	(*reflect.SliceHeader)(unsafe.Pointer(&result)).Cap = c.len
	return result
}
