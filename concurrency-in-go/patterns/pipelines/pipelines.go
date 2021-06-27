// A pipeline is a series of stages connected by channels.
// A stages is just a function that:
// - Receive values from upstream via inbound channels.
// - Perform some function on that data, usually producing new values.
// - Send values downstream via outbound channels.
// Each stage has any number of inbound and outbound channels, except the first
// and last stages, which have only outbound or inbound channels, respectively:
// - The first stage is called the source or producer.
// - The last stage is called the sink or consumer.
package main

import (
	"fmt"
)

// add function is a stage of pipeline
func add(done <-chan interface{}, numChan <-chan int, additive int) <-chan int {
	ch := make(chan int)

	go func() {
		defer fmt.Println("The goroutine child of the add function is released")
		defer close(ch)

		// You can use "for range" to add a number to emitted value "numChan".
		// We will learn the pattern or-channel in the next samples.
		/* for v := range numChan {
			select {
			case <-done:
				return
			case ch <- v + additive:
			}
		} */

		// Here uses "for loop" to add a number to emitted value "numChan".
		for {
			select {
			case <-done:
				return
			case v, ok := <-numChan:
				if !ok {
					return
				}

				select {
				case <-done:
					return
				case ch <- v + additive:
				}
			}
		}
	}()

	return ch
}

// multiply function is a stage of pipeline
func multiply(done <-chan interface{}, numChan <-chan int, multiplier int) <-chan int {
	ch := make(chan int)

	go func() {
		defer fmt.Println("The goroutine child of the multiply function is released")
		defer close(ch)

		// You can use "for range" to multiply a number to emitted value "numChan"
		/* for v := range numChan {
			select {
			case <-done:
				return
			case ch <- v * multiply:
			}
		} */

		/**
		Using the or done channel and the for range to replace the for loop code at line 81
			Example:
				for v := range orDone(done, numChan) {
					ch <- v * multiplier
				}
		*/

		// Here uses "for loop" to multiply a number to emitted value "numChan"
		for {
			select {
			case <-done:
				return
			case v, ok := <-numChan:
				if !ok {
					return
				}

				select {
				case <-done:
					return
				case ch <- v * multiplier:
				}
			}
		}
	}()

	return ch
}

func main() {
	additive := 1
	multiplier := 10

	done := make(chan interface{})
	numChan := generate(done)

	ch := take(done, add(done, multiply(done, numChan, multiplier), additive), 17)

	// Consumer
	for v := range ch {
		fmt.Println(v)
	}
}

// =============================================================================
// =====================    GENERATOR CHANNELS    ==============================
// =============================================================================

// generate returns a chan that emit int value until the done channel is done.
// generate function is a stage of pipeline as producer stage.
func generate(done <-chan interface{}) <-chan int {
	numChan := make(chan int)

	go func() {
		defer close(numChan)

		for i := 0; ; i++ {
			select {
			case <-done:
				return
			case numChan <- i:
			}
		}
	}()

	return numChan
}

// take function is a stage of pipeline.
func take(
	done <-chan interface{},
	numChan <-chan int,
	num int,
) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case ch <- <-numChan:
			}
		}
	}()

	return ch
}
