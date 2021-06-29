// Fan-out: Multiple functions can read from the same channel until that channel is closed.
// Fan-in: A function can read from multiple inputs and proceed until all are closed
// by multiplexing the input channels onto a single channel that's closed when all the inputs are closed.
// Using fan-out/in when:
// - Do not care about the order of the values of downstream.
// - Stage consumes many time.
package main

import (
	"fmt"
	"sync"
	"time"
)

// MiddleStage represents the stages of the pipeline excluding the first and last stages.
type MiddleStage = func(done, in <-chan interface{}) <-chan interface{}

func main() {
	start := time.Now()

	done := make(chan interface{})

	// Without fan-in/out, uncomment line 21 and comment line 23 to compare time consumed
	// ch := sleep(done, take(done, gen(done), 5))

	ch := fanIn(done, fanOut(done, take(done, gen(done), 5), sleep)...)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Printf("Time consumed is %v", time.Since(start))
}

func fanIn(done <-chan interface{}, ins ...<-chan interface{}) <-chan interface{} {
	if len(ins) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	multiplex := make(chan interface{})

	// Start an output goroutine for each input channel in "ins".
	// The output goroutine copies values from the "in" channel to the "multiplex" channel
	// until the "in" channel is closed, then calls wg.Done.
	output := func(in <-chan interface{}) {
		defer wg.Done()

		/**
		Using the or done channel and the for range to replace the for loop code at line 57
			Example:
				for v := range orDone(done, in) {
					multiplex <- v
				}
		*/

		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-done:
					return
				case multiplex <- v:
				}
			}
		}
	}

	wg.Add(len(ins))
	for _, in := range ins {
		go output(in)
	}

	// Start a goroutine to close multiplex once all the "output" goroutines are done.
	go func() {
		defer close(multiplex)

		wg.Wait()
	}()

	return multiplex
}

func fanOut(done, in <-chan interface{}, heavy MiddleStage) []<-chan interface{} {
	fanOutNum := 3
	outs := make([]<-chan interface{}, fanOutNum)

	// Distribute the heavy work across fanOutNum goroutines that read from in.
	for i := 0; i < fanOutNum; i++ {
		outs[i] = heavy(done, in)

		// Do not need to transfer channel
		// out := make(chan interface{})
		// outs[i] = out

		// go func() {
		// 	defer close(out)

		// 	ch := heavy(done, in)

		// 	for {
		// 		select {
		// 		case <-done:
		// 			return
		// 		case v, ok := <-ch:
		// 			if !ok {
		// 				return
		// 			}

		// 			select {
		// 			case <-done:
		// 				return
		// 			case out <- v:
		// 			}
		// 		}
		// 	}
		// }()
	}

	return outs
}

func sleep(done, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for v := range in {
			time.Sleep(time.Second)

			select {
			// case <-done:
			// 	return

			case out <- v:
			}
		}
	}()

	return out
}

func take(done, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return

			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}

	}()

	return out
}

func gen(done <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for i := 0; ; i++ {
			select {
			case <-done:
				return

			case out <- i:
			}
		}
	}()

	return out
}
