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
