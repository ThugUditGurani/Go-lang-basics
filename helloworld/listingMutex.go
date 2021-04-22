package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines
	counter int
	// wg is used to wait for the program to finish
	wgSync sync.WaitGroup
	// mutex is used to define a critical section of code
	mutex sync.Mutex
)

func main() {
	// Add a count of two , one for each goroutine
	wgSync.Add(2)

	// Create two goroutines
	go incCounter(1)
	go incCounter(2)

	//wait for the goroutines to finish
	wgSync.Wait()
	fmt.Printf("Final Counter : %d",counter)

}

func incCounter(id int)  {
	// Schedule the call to Done to tell main we are done
	defer wgSync.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			// Yield the thread and be placed back in queue
			runtime.Gosched()
			// Increment our local value of counter
			value++
			// Store the value back into counter
			counter = value
		}
		mutex.Unlock()
		// Release the lock and allow any waiting goroutine through
	}
}