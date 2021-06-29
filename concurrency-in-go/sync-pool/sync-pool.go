// Using sync.Pool structure we can pool objects and reuse them.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Person dummy struct.
type Person struct {
	name, address string
	age           int
}

var pool = sync.Pool{
	// New creates an object when the pool has nothing available to return.
	New: func() interface{} {
		fmt.Println("New an instance in the pool")
		return new(Person)
	},
}

func main() {
	for i := 0; i < 10; i++ {
		func(index int) {
			p := pool.Get().(*Person)
			defer pool.Put(p)

			time.Sleep(time.Second)

			fmt.Println(i, p)
		}(i)
	}

	fmt.Println("Only one new person is created")
}
