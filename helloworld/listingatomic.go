package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown is a flag to alert running goroutines to shutdown
	shutdown int64
	// wg is used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {
	//Add a count of two , one for each goroutine.
	wg.Add(2)
	//Create two goroutines
	go doWork("A")
	go doWork("B")

	//Give the goroutines time to run
	time.Sleep(1 * time.Second)

	//Safely flag it is time to shutdown
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown,1)

	wg.Wait()
}

func doWork(name string)  {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work",name)
		time.Sleep(250 * time.Second)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s down \n", name)
			break
		}
	}
}
