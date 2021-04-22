package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wgUnbuf sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}


func main() {
	// Create an UnBuffered Channel
	court := make(chan int)

	// Add a count of two , one for each goroutine.
	wgUnbuf.Add(2)


	// Launch two Players
	go player("Nadal",court)
	go player("Udit", court)

	// Start the set
	court <- 1

	// Wait fot the game to finish
	wgUnbuf.Wait()
}

// player simulates a person playing the game of tennis
func player(name string, court chan int) {
	// schedule the call to done to tell main we are done
	defer wgUnbuf.Done()

	for  {
		// Wait for the ball to be hit back to us
		ball , ok := <- court
		fmt.Printf("ok %v",ok)
		if !ok {
			// If the channel was closed we won
			fmt.Printf("Player %s won \n",name)
			return
		}

		// pick a random number and see if we miss the ball
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed \n",name)
			// Close the channel to signal we lost
			close(court)
			return
		}

		// Display and then increment the hit count by one
		fmt.Printf("Player %s hit %d \n",name,ball)
		ball++

		//Hit the ball to the opposing player
		court <- ball
	}
}
