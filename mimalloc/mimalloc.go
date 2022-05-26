package mimalloc

/*
#cgo CFLAGS: -I./include
#include <./src/static.c>
*/
import "C"
import "unsafe"

func Malloc(size int) unsafe.Pointer {
	return C.mi_malloc(C.size_t(size))
}

func Zalloc(size int) unsafe.Pointer {
	return C.mi_zalloc(C.size_t(size))
}

func Calloc(count, size int) unsafe.Pointer {
	return C.mi_calloc(C.size_t(count), C.size_t(size))
}

func Realloc(ptr unsafe.Pointer, size int) unsafe.Pointer {
	return C.mi_realloc(ptr, C.size_t(size))
}

func Free(ptr unsafe.Pointer) {
	C.mi_free(ptr)
}
