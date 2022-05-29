package arena_test

import (
	"testing"

	"github.com/unsafe-risk/mi/arena"
)

type Person struct {
	Name string
	Age  int
	Addr string
	Zip  int
}

const MAX = 1000000

func BenchmarkPerson(b *testing.B) {
	a := arena.New()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX; j++ {
			p := arena.NewOf[Person](a)
			p.Name = "John"
			p.Age = 32
			p.Addr = "London"
			p.Zip = 1111
		}
	}
	a.Free()
}

//go:noinline
func StdNewPerson() *Person {
	p := new(Person)
	return p
}

func BenchmarkAllocateStdNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < MAX; j++ {
			p := StdNewPerson()
			p.Name = "John"
			p.Age = 32
			p.Addr = "London"
			p.Zip = 1111
		}
	}
}
