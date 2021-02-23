package main

import (
	"fmt"
	"runtime"
	"sync"
	atomic2 "sync/atomic"
	"time"
)

func main() {
	//mutex()
	atomic()
}

func mutex() {
	runtime.GOMAXPROCS(4)
	var balance int
	var wg sync.WaitGroup
	var mut sync.Mutex
	deposit := func(amount int){
		mut.Lock()
		balance += amount
		mut.Unlock()
	}
	withdrawal := func(amount int) {
		mut.Lock()
		balance -= amount
		defer mut.Unlock()
	}

	//make 100 deposits of $1
	//and 100 withdrawal of $1 concurrently
	// run the program and check result.
	// TODO: fix the issue for consistent output
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)

}


func atomic(){
	runtime.GOMAXPROCS(4)
	var counter uint64
	var wg sync.WaitGroup

	for i := 0;i < 50; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0;c < 1000; c++{
				atomic2.AddUint64(&counter,1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter: ",counter)
}

var sharedRsc = make(map[string]interface{})
func cond(){
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	c.L.Unlock()
	wg.Wait()
}



func once(){
	var wg sync.WaitGroup
	var once sync.Once

	load := func() {
		fmt.Println("Run only once initialization function")
	}

	for i := 0;i<10;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(load)
		}()
	}
	once.Do(load)
}

func pool(){

}