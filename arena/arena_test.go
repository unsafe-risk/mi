package arena_test

import (
	"runtime"
	"testing"

	"github.com/unsafe-risk/mi/arena"
)

type Person struct {
	Name string
	Age  int
	Addr string
	Zip  int
}

const MAX = 2000000

func BenchmarkMiArenaPerson(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			a := arena.New()
			for i := 0; i < MAX; i++ {
				p := arena.NewOf[Person](a)
				p.Name = "John"
				p.Age = 32
				p.Addr = "Istanbul"
				p.Zip = 397
			}
			a.Free()
		}
	})
}

//go:noinline
func StdNewPerson() *Person {
	p := new(Person)
	return p
}

func BenchmarkStdNew(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := 0; i < MAX; i++ {
				p := StdNewPerson()
				p.Name = "John"
				p.Age = 32
				p.Addr = "London"
				p.Zip = 1111
			}
			runtime.GC()
		}
	})
}
