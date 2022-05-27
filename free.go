package mi

import (
	"sync/atomic"
	"unsafe"

	"github.com/unsafe-risk/mi/mimalloc"
)

func Free(ptr unsafe.Pointer) {
	mimalloc.Free(ptr)
}

func FreeOf[T any](ptr *T) {
	mimalloc.Free(unsafe.Pointer(ptr))
	atomic.AddInt64(&alloced, -int64(unsafe.Sizeof(*ptr)))
}

func FreeCArray[T any](ptr CArray[T]) {
	mimalloc.Free(ptr.ptr)
	atomic.AddInt64(&alloced, -int64(ptr.len*ptr.size))
}
