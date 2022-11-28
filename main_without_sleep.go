package main

import (
	"fmt"
	"sync"
	"time"
)

var result float64 = 1

func parallelCalculation() {
	totalIterations := 1_000_000
	startParallel := time.Now()
	var wg sync.WaitGroup
	// Tell the 'wg' WaitGroup how many threads/goroutines
	//   that are about to run concurrently.
	wg.Add(totalIterations)
	for i := 0; i < totalIterations; i++ {
		// Spawn a thread for each iteration in the loop.
		// Pass 'i' into the goroutine's function
		//   in order to make sure each goroutine
		//   uses a different value for 'i'.
		go func(i int) {
			// At the end of the goroutine, tell the WaitGroup
			//   that another thread has completed.
			defer wg.Done()
			var num float64 = float64(4.0 * (i + 1) * (i + 1))
			result *= num / (num - 1)
		}(i)
	}
	result *= 2
	// Wait for `wg.Done()` to be exectued the number of times
	//   specified in the `wg.Add()` call.
	// `wg.Done()` should be called the exact number of times
	//   that was specified in `wg.Add()`.
	// If the numbers do not match, `wg.Wait()` will either
	//   hang infinitely or throw a panic error.
	wg.Wait()
	fmt.Printf("Parallel implementation time %v\n", time.Since(startParallel))

}
func wallis(i int) {
	var num float64 = float64(4.0 * i * i)
	result *= num / (num - 1)
}
func sequentialCalculation() {
	totalIterations := 1000000
	startSequential := time.Now()
	for i := 1; i <= totalIterations; i++ {
		wallis(i)
	}
	result *= 2
	fmt.Printf("Sequential execution ended in %v\n", time.Since(startSequential))
}
func main() {
	sequentialCalculation()
	result = 1
	parallelCalculation()
}
