package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4 // Number of goroutines to use
	taskLoad = 10 // Amount of work to process
)

// wg is used to wait for the program to finish
var wgBuf sync.WaitGroup

// init is called to initialize the package by the
// Go runtime prior to any other code being executed.
func init() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())
}

func main() {
	// Create a buffered channel to manage the task load
	tasks := make(chan string, taskLoad)


	// Launch goRoutines to handle the work
	wgBuf.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks,gr)
	}

	// Add a bunch of work to get done
	for post := 1; post <= taskLoad; post ++{
		tasks <- fmt.Sprintf("Task : %d",post)
	}
}


func worker(tasks chan string, worker int) {
	// Report that we just returned
	defer wgBuf.Done()

	for  {
		task , ok := <- tasks
		if !ok {
			// This means the channel is empty and closed
			fmt.Printf("Worker: %d : shutting Down \n",worker)
			return
		}

		// Display we are starting the work
		fmt.Printf("Worker: %d : started %s \n",worker,task)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)

		//Display we finished the work
		fmt.Printf("Worker %d : Completed %s \n",worker,task)

	}
}
