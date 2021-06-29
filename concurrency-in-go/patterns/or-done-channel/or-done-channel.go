// The or done channel receives a done channel and an inbound channel.
// The or done will be closed when the done channel is closed or the inbound channel is closed.
package main

import "fmt"

func orDone(done, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		// Using loop to retrieve the values upstream.
		for {
			select {
			// Release the goroutine when receiving a of done channel.
			case <-done:
				return

			// Check values upstream of the inbound channel.
			case v, ok := <-in:
				// Release the goroutine when the inbound channel is closed.
				if !ok {
					return
				}

				select {
				// Release the goroutine when receiving a of done channel.
				case <-done:
					return

				// Send values upstream of the inbound channel to the downstream - the outbound channel.
				case out <- v:
				}
			}
		}
	}()

	return out
}

func main() {
	done := make(chan interface{})

	// Consume values upstream from gen
	for v := range orDone(done, gen(done)) {
		fmt.Println(v)
	}
}

func gen(done <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for i := 0; i < 10; i++ {
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()

	return out
}
