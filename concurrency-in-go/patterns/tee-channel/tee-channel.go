// The tea channel get value upstream of an inbound channel
// then send this value downstream via outbound channels.
package main

import (
	"fmt"
	"sync"
)

func tee(done, in <-chan interface{}, teeNum int) []<-chan interface{} {
	outs := make([]chan interface{}, teeNum)

	for i := 0; i < teeNum; i++ {
		outs[i] = make(chan interface{})
	}

	go func() {
		defer func() {
			for i := 0; i < teeNum; i++ {
				close(outs[i])
			}
		}()

		// Should use the or done channel here.
		for v := range in {
			var wg sync.WaitGroup

			wg.Add(teeNum)
			for i := 0; i < teeNum; i++ {
				go func(i int) {
					defer wg.Done()

					select {
					case <-done:
						return
					case outs[i] <- v:
					}
				}(i)
			}

			wg.Wait()
		}
	}()

	tees := make([]<-chan interface{}, teeNum)

	for i, v := range outs {
		tees[i] = v
	}

	return tees
}

func main() {
	var wg sync.WaitGroup

	teeNum := 5
	done := make(chan interface{})
	chs := tee(done, gen(done), teeNum)

	wg.Add(teeNum)

	// Do not use for rang here because inside the loop uses the function literal goroutine.
	for i := 0; i < teeNum; i++ {
		go func(i int) {
			defer wg.Done()

			for v := range chs[i] {
				fmt.Printf("Channel[%v] %v\n", i, v)
			}
		}(i)
	}

	wg.Wait()
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

func simpleTee(
	done <-chan interface{},
	in <-chan interface{},
) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)

		// Should use the or done channel here.
		for val := range in {
			var out1, out2 = out1, out2

			for i := 0; i < 2; i++ {
				select {
				case <-done:

				case out1 <- val:
					// Once we’ve written to a channel, we set its shadowed copy to nil
					// so that further writes will block and the other channel may continue.
					// We set out1 = nil, so the "case out1/nil <- val:" will not run because of error.
					out1 = nil

				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}
