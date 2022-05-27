package mi

import "unsafe"

var pointerSize = int(unsafe.Sizeof((*byte)(nil)))
