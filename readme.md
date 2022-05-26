# mimalloc-go

this is go binding of `mimalloc` with cgo.

# installing

`go get -u github.com/unsafe-risk/mi`

# how to use

## MAlloc and pointer

```go
package main

import (
	"fmt"

	"github.com/unsafe-risk/mi"
)

func main() {
	a := mi.MAllocOf[int64]()
	*a = 99
	fmt.Println(*a)
	mi.FreeOf(a)
}
```

`MAllocOf` is make a dynamic pointer of type `T`.

`FreeOf` is free the memory of the pointer.

## CAlloc and CArray

```go
package main

import (
	"fmt"

	"github.com/unsafe-risk/mi"
)

func main() {
	a := mi.CAllocOf[int64](6)
	a.Set(0, 1)
	a.Set(1, 4)
	a.Set(2, 9)
	a.Set(3, 16)
	a.Set(4, 25)
	a.Set(5, 36)
	for i := 0; i < a.Length(); i++ {
		fmt.Println(a.At(i))
	}
	mi.FreeCArray(a)
}
```

`CAllocOf` is make a `CArray` of type `T` with length.

`FreeCArray` is free the memory of the pointer.
