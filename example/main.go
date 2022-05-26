package main

import (
	"fmt"

	"github.com/unsafe-risk/mi"
)

func main() {
	m := mi.MAllocOf[int64]()
	*m = 99
	fmt.Println(*m)
	mi.FreeOf(m)

	fmt.Println("----------")

	c := mi.CAllocOf[int64](6)
	c.Set(0, 1)
	c.Set(1, 4)
	c.Set(2, 9)
	c.Set(3, 16)
	c.Set(4, 25)
	c.Set(5, 36)
	for i := 0; i < c.Length(); i++ {
		fmt.Println(c.At(i))
	}
	mi.FreeCArray(c)
}
