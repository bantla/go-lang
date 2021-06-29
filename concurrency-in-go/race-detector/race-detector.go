// A data race occurs when two goroutines access the same variable concurrently
// and at least one of the accesses is a write.
// https://golang.org/ref/mem
package main

import (
	"fmt"
	"os"
	"time"
)

// Watchdog ...
type Watchdog struct{ last int64 }

// Start ...
func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

// KeepAlive ...
func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access.
}

func main() {
	w := &Watchdog{8}

	w.Start()
	w.KeepAlive()

	time.Sleep(10 * time.Second)
}

// func main() {
// 	c := make(chan bool)
// 	m := make(map[string]string)
// 	go func() {
// 		m["1"] = "a" // First conflicting access.
// 		c <- true
// 	}()
// 	m["2"] = "b" // Second conflicting access.
// 	<-c
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}
// }
