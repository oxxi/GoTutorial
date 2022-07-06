package main

import (
	"fmt"
	"sync"
)

type Lockable[T any] struct {
	sync.Mutex
	Data T
}

func main() {
	var n Lockable[uint64]
	n.Lock()
	n.Data = 50
	n.Multiple(func(v *uint64) {
		n := *v
		*v = n * 3

	})
	n.Unlock()

	fmt.Printf("the values for uint64 is %v\n", n.Data)

	var f Lockable[float64]
	f.Data = 1.50
	f.Multiple(func(t *float64) {
		f := *t
		*t = f * 2.205
	})
	fmt.Printf("the values for float64 is %v\n", f.Data)

}

func (l *Lockable[T]) Multiple(f func(*T)) {
	f(&l.Data)
}
