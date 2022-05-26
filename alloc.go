package mi

import (
	"unsafe"

	"github.com/unsafe-risk/mi/mimalloc"
)

func MAlloc(size int) unsafe.Pointer {
	return mimalloc.Malloc(size)
}

func MAllocOf[T any]() *T {
	var a T
	return (*T)(mimalloc.Malloc(int(unsafe.Sizeof(a))))
}

func CAlloc(count, size int) unsafe.Pointer {
	return mimalloc.Calloc(count, size)
}

func CAllocOf[T any](count int) CArray[T] {
	var a T
	size := int(unsafe.Sizeof(a))
	rs := CArray[T]{
		ptr:  mimalloc.Calloc(count, size),
		len:  count,
		size: size,
	}
	return rs
}
