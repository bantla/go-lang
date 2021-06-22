// Lexical confinement involves using lexical scope to expose only the correct
// data and concurrency primitives for multiple concurrent processes to use.
package main

import "fmt"

// Export the channel that is bounded by the produce function (Lexical confinement)
func produce() <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- fmt.Sprint(i)
		}
	}()

	return ch
}

func consumer(ch <-chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := produce()
	consumer(ch)
}
