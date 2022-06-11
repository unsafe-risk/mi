package arena

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/unsafe-risk/mi/mimalloc"
)

// This code is implemented by lemon-mint.
// I brought the code from unsafe-risk/umem.
// I edit a lot of the code to make it work with mimalloc's arena.

type Arena struct {
	allocatedPointers []unsafe.Pointer
}

func NewFinalizer() *Arena {
	a := &Arena{}
	runtime.SetFinalizer(a, arenaFinalizer)
	return a
}

func New() *Arena {
	a := &Arena{}
	return a
}

func arenaFinalizer(a *Arena) {
	a.Free()
}

func (r *Arena) Free() {
	for _, p := range r.allocatedPointers {
		mimalloc.Free(p)
	}
}

func (r *Arena) allocate(size int) uintptr {
	p := mimalloc.Malloc(size)
	r.allocatedPointers = append(r.allocatedPointers, p)
	return uintptr(p)
}

func (r *Arena) NewBytes(size int) []byte {
	sh := reflect.SliceHeader{
		Data: r.allocate(size),
		Len:  int(size),
		Cap:  int(size),
	}
	return *(*[]byte)(unsafe.Pointer(&sh))
}

func (r *Arena) NewString(b []byte) string {
	s := r.NewBytes(len(b))
	copy(s, b)
	return *(*string)(unsafe.Pointer(&s))
}

func (r *Arena) HeapString(b string) string {
	s := r.NewBytes(len(b))
	copy(s, b)
	return *(*string)(unsafe.Pointer(&s))
}

func (r *Arena) Allocate(size int) unsafe.Pointer {
	return unsafe.Pointer(r.allocate(size))
}

func NewOf[T any](r *Arena) *T {
	var zero T
	p := (*T)(r.Allocate(int(unsafe.Sizeof(zero))))
	*p = zero
	return p
}

func NewOfUninitialized[T any](r *Arena) *T {
	var zero T
	return (*T)(r.Allocate(int(unsafe.Sizeof(zero))))
}

func NewSliceOfUninitialized[T any](r *Arena, len int) []T {
	var zero T
	p := r.Allocate(int(unsafe.Sizeof(zero) * uintptr(len)))
	sh := reflect.SliceHeader{
		Data: uintptr(p),
		Len:  len,
		Cap:  len,
	}
	v := *(*[]T)(unsafe.Pointer(&sh))
	return v
}

func NewSliceOf[T any](r *Arena, len int) []T {
	var zero T
	v := NewSliceOfUninitialized[T](r, len)
	for i := 0; i < len; i++ {
		v[i] = zero
	}
	return v
}
