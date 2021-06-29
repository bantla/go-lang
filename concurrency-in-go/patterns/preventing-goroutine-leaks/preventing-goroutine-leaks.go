// If a goroutine is responsible for creating a goroutine,
// it is also responsible for ensuring it can stop the goroutine.
package main

import "fmt"

func produceNumbers(done <-chan interface{}) <-chan int {
	ch := make(chan int)

	// The createChan function create goroutine
	go func() {
		defer fmt.Println("Release goroutine")
		defer close(ch)

		for i := 0; ; i++ {
			select {
			case ch <- i:

			// So the createChan function is also responsible for ensuring it can stop the goroutine
			case <-done:
				return
			}
		}
	}()

	return ch
}

func main() {
	done := make(chan interface{})
	ch := produceNumbers(done)

	for v := range ch {
		if v == 3 {
			// Terminate the int stream
			close(done)
		}

		// 4 can be printed, we do not guarantee the main goroutine or
		// the produceNumbers goroutine will be run next. A subsequent goroutine is run at random
		fmt.Println(v)
	}
}
