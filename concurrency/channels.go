package main

import "fmt"

/*
Communicate data between Goroutines
Synchronise GO routines
Typed
Thread Safe
*/


func main() {
	//simpleChannels()
	//channelsRangeBuffered()
	//channelDirection()
	channelOwnership()
}

/*
Channel Ownership
*/
func channelOwnership() {
	owner := func() <- chan int{
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++{
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <- chan int){
		// read values from channel
		for v := range ch{
			fmt.Printf("Received: %d",v)
		}
		fmt.Println("Done receiving")
	}

	ch := owner()
	consumer(ch)

}

/*
Channel Direction
*/
func channelDirection() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go genMsg(ch1)
	go relayMsg(ch1,ch2)

	v := <-ch2
	fmt.Println(v)
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	m := <- ch1
	//send in on ch2
	ch2 <- m
}

func genMsg(ch1 chan<- string) {
	// Send message on ch1
	ch1 <- "message"
}


/*
Channels: Range and Buffered Channnels
*/

func channelsRangeBuffered() {
	ch := make(chan int,6)
	go func() {
		defer close(ch)
		for i := 0;i < 6; i ++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch{
		fmt.Println(v)
	}
}


/*
Simple Channels Communication
*/
func simpleChannels() {
	ch := make(chan int)
	go func(a,b int) {
		c := a + b
		ch <- c
	}(1,2)
	r := <-ch
	fmt.Printf("Computed Value %v",r)
}