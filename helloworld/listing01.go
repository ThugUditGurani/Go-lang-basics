package main

import (
	"fmt"
	"runtime"
	"sync"
)



func main() {

	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(4)


	// wg is used to wait for the program to finish.
	// Add a Count of two , one for each goRoutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start GoRoutines")

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to done to tell main we are done.
		defer wg.Done()
		for count := 0 ;  count < 3 ; count ++ {
			for char := 'a' ; char < 'a' + 26; char++ {
				fmt.Printf("%c ",char)
			}
		}
	}()

	// Decalre an anonymous function and create a goroutine
	go func() {
		//Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count  < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ",char)
			}
		}

	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\n Terminating Program")
}
