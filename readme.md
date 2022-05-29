# mimalloc-go

this is a go wrapper of `mimalloc` with cgo.

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

## CArray and Slice

```go
package main

import (
	"fmt"
	"sort"

	"github.com/unsafe-risk/mi"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	c := mi.CAllocOf[int](10)
	for i := 0; i < 10; i++ {
		c.Set(i, -i*i)
	}
	sort.Ints(c.ToSlice())
	fmt.Println(c.ToSlice())
	mi.FreeCArray(c)
}
```

`ToSlice()` method is return a fixed length slice of `CArray`.

It equals a slice of golang, but length is immutable.

## ARENA

```go
package main

import (
	"fmt"

	"github.com/unsafe-risk/mi/arena"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	a := arena.New()
	defer a.Free()
	p := arena.NewOf[Person](a)
	p.Name = "John"
	p.Age = 30
	fmt.Println(p)
}
```

`arena` is a memory pool implementation from unsafe-risk/umem.

The `arena` of this project is based on `mimalloc`.

And, this is test code. I cannot recommend to use `arena` for your project.

### arena bench

```bash
goos: darwin
goarch: amd64
pkg: github.com/unsafe-risk/mi/arena
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMiArenaPerson-16    	      33	  33580864 ns/op	      42 B/op	       1 allocs/op
BenchmarkStdNew-16           	      39	  30365484 ns/op	96002062 B/op	 2000022 allocs/op
```

```bash
goos: linux
goarch: amd64
pkg: github.com/unsafe-risk/mi/arena
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkMiArenaPerson-4   	      19	  59619689 ns/op	     195 B/op	       0 allocs/op
BenchmarkStdNew-4          	      20	  54869378 ns/op	96000574 B/op	 2000006 allocs/op
```

```bash
goos: linux
goarch: amd64
pkg: github.com/unsafe-risk/mi/arena
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkMiArenaPerson-16    	      32	  49734124 ns/op	     861 B/op	       2 allocs/op
BenchmarkStdNew-16           	      52	  20565911 ns/op	96001187 B/op	 2000011 allocs/op
```

```bash
goos: linux
goarch: arm64
pkg: github.com/unsafe-risk/mi/arena
BenchmarkMiArenaPerson-4   	      85	  11788425 ns/op	      10 B/op	       0 allocs/op
BenchmarkStdNew-4          	      26	  44319635 ns/op	96001133 B/op	 2000011 allocs/op
```
