package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Concurrency is about multiple things happening at same time in random order
Go provides a built-in support for concurrency
Concurrency is composition of independent execution computations, which may or may not run in parallel
*/

/*
Parallelism is ability to execute multiple computations simultaneously
*/

/*
Operating System :-
The Job of Operation system is to give fair chances for all processes access to CPU, memory and other resources.
Process :-
An Instances of a running program is called a process.
Process provides environment for program to execute.
-> OS allocates memory.
-> Code - machine instructions
-> Data - Global Data
-> Heap - Dynamic memory allocation
-> Stack -  Local Variables of function
Threads :-
Threads are smallest unit of execution that CPU accepts.
Process has atleast one thread main thread
Process can have multiple threads
Threads share same address space
Threads run independent of each other
OS scheduler makes scheduling decisions at thread level, not process level
Threads can run concurrently or in parallel
*/


/*
Communicating Sequential Processes(CSP)
TOny Hoare(1978)
Each Process is built for sequential execution
Data is communicated between processes. NO shared memory
Scale by adding more of the same.
*/

/*
GORoutines
WE can think Goroutine as User Space Threads managed by go runtime
GOroutines extremely lightweight. Goroutines starts with 2KB of stack, which grows and shrinks as required.
Low CPU overhead - three instructions per funstions call.
Can create hundreds of thousands of goroutines in same address space.
Channels are used for communication of data between goroutine. sharing of memory can be avoided.
*/


func main() {
	//goRoutineBasic()
	//goRoutineWaitGroup()
	//goRoutineClosures()
	//goRoutineClosurestwo()
}

/*
The GO scheduler is part of the GO Runtime. It is known as M:N scheduler
GO scheduler runs is user space
GO scheduler uses OS threads to schedule goroutine for execution
GO routine runs in the context of os threads
GO runtime create number of worker threads, equal to GOMaxprocs
Go MaxProcs - default value is number of processors on machine
Go scheduler distribute runnable goroutine over multiple worker OS threads
As of GO 1.14, Go scheduler implements asynchronous preemption
This prevents long running Goroutines from hogging onto CPU, that could block other Go routines
The asynchronous preemption is triggered based on a time condition, when a goroutine is running for more than 10ms, GO will try to preempt it.
*/


/*
Goroutine- closures
*/
func goRoutineClosures() {
	var wg sync.WaitGroup
	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("Value of i: %v",i)
		}()
		fmt.Println("Return from function")
		return
	}
	incr(&wg)
	wg.Wait()
	fmt.Println("Done")
}

func goRoutineClosurestwo()  {
	var wg sync.WaitGroup
	for i:=0;i<=3;i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

/*
Goroutine WaitGroup
*/
func goRoutineWaitGroup() {
	var data int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()
	wg.Wait()
	fmt.Printf("The Value of data is %v",+data)
}

/*
Goroutine call function basics
*/

func fun(s string){
	for i := 0;i < 3;i++{
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}


func goRoutineBasic() {
	fun("Direct calls")

	//go routine function calls
	go fun("go-routine-1")
	//go routine with anonymous function
	go func() {
		fun("go-routine-2")
	}()
	//go routine with function value call
	fv := fun
	go fv("go-routine-3")

	// Wait for goroutine to end
	fmt.Println("wait for goroutines..")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done....")
}