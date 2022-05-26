package mi

import (
	"unsafe"

	"github.com/unsafe-risk/mi/mimalloc"
)

func Free(ptr unsafe.Pointer) {
	mimalloc.Free(ptr)
}

func FreeOf[T any](ptr *T) {
	mimalloc.Free(unsafe.Pointer(ptr))
}

func FreeCArray[T any](ptr CArray[T]) {
	mimalloc.Free(ptr.ptr)
}
