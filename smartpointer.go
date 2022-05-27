package mi

import (
	"runtime"
	"unsafe"

	"github.com/unsafe-risk/mi/mimalloc"
)

type SmartPointer[T any] struct {
	Ref *T
}

func NewSmartPointer[T any](ptr unsafe.Pointer) *SmartPointer[T] {
	sp := &SmartPointer[T]{Ref: (*T)(ptr)}
	runtime.SetFinalizer(sp, func(sp *SmartPointer[T]) {
		mimalloc.Free((unsafe.Pointer)(sp.Ref))
	})
	return sp
}
