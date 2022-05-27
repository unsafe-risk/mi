package mi

import (
	"sync/atomic"
	"unsafe"

	"github.com/unsafe-risk/mi/mimalloc"
)

func Alloced() int64 {
	return atomic.LoadInt64(&alloced)
}

func MAlloc(size int) unsafe.Pointer {
	return mimalloc.Malloc(size)
}

func MAllocOf[T any]() *T {
	var a T
	size := unsafe.Sizeof(a)
	atomic.AddInt64(&alloced, int64(size))
	return (*T)(mimalloc.Malloc(int(size)))
}

func MAllocSmart[T any]() *SmartPointer[T] {
	var a T
	size := unsafe.Sizeof(a)
	atomic.AddInt64(&alloced, int64(size))
	return NewSmartPointer[T](mimalloc.Malloc(int(size)))
}

func CAlloc(count, size int) unsafe.Pointer {
	atomic.AddInt64(&alloced, int64(count*size))
	return mimalloc.Calloc(count, size)
}

func CAllocOf[T any](count int) CArray[T] {
	var a T
	size := int(unsafe.Sizeof(a))
	atomic.AddInt64(&alloced, int64(count*size))
	rs := CArray[T]{
		ptr:  mimalloc.Calloc(count, size),
		len:  count,
		size: size,
	}
	return rs
}
