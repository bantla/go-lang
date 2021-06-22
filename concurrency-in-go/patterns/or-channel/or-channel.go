// Combine one or more done channels into a single done channel
// that closes if any of its component channels close
package main

import (
	"fmt"
	"time"
)

func or(doneChannels ...<-chan interface{}) <-chan interface{} {
	length := len(doneChannels)

	switch length {
	// Check if the list of arguments is empty
	case 0:
		return nil

	// Returns the first channel in the list of arguments
	// if the list of arguments has only one element
	case 1:
		return doneChannels[0]
	default:
		// The "orDone" channel check if the child goroutine is done
		orDone := make(chan interface{})

		// Make a goroutine to receive the done value from all channels in the list of arguments
		go func() {
			defer close(orDone)

			if length == 2 {
				// Receive the done value of the first channel, the second channel
				select {
				case <-doneChannels[0]:
				case <-doneChannels[1]:
				}
			} else {
				// Receive the done value of the first channel, the second channel
				// and the third chanel onwards
				select {
				case <-doneChannels[0]:
				case <-doneChannels[1]:

				// Use recurvive to receive the done value of the third chanel onwards
				// in the list of arguments
				case <-or(doneChannels[2:]...):
				}
			}

		}()

		return orDone
	}
}

// Create a done channel that closes after "duration" seconds
func createDone(duration time.Duration) <-chan interface{} {
	done := make(chan interface{})

	go func() {
		defer close(done)

		time.Sleep(duration * time.Second)
	}()

	return done
}

func main() {
	start := time.Now()

	orChannel := or(
		createDone(4),
		createDone(5),
		createDone(6),
		createDone(3),
		createDone(7),
	)

	fmt.Printf(
		"Do not terminate the main goroutine when call 'or' function above %v\n",
		time.Since(start),
	)

	if orChannel != nil {
		<-orChannel
	}

	fmt.Printf("Done after %v", time.Since(start))
}

func orWithMoreGoroutine(doneChannels ...<-chan interface{}) <-chan interface{} {
	switch len(doneChannels) {
	case 0:
		return nil
	case 1:
		return doneChannels[0]
	default:
		orDone := make(chan interface{})

		go func () {
			defer close(orDone)

			select {
			case <-doneChannels[0]:
			case <-orWithMoreGoroutine(doneChannels[1:]...):
			}
		}()

		return orDone
	}
}
