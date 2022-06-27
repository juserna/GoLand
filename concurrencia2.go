package main

import (
	"fmt"
	"runtime"
	"time"
)

//Goroutines en un servidor web
//simulando async http

/*
GO routines advantages:

- They are lightweight
- Ability to scale seamlessly
- They are virtual threads. (One OS thread can multiplex many goroutines)
- Less memory required (2K)
- Goroutines can communitace each other using channels
*/

//Secuential execution = 7.074346549s
//concurrent execution = 5.00182222s
func main() {
	cpus := runtime.NumCPU() // 4
	runtime.GOMAXPROCS(cpus)

	numGR := runtime.NumGoroutine()

	fmt.Println(numGR)

	now := time.Now()
	//1st Go routine
	go func() {
		for i := 10; i < 20; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i)
		}
	}()

	//2nd Go routine
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i)
		}
	}()

	go longAndHeavyTask()

	elapse := time.Since(now)
	fmt.Println(elapse.String())

}

func longAndHeavyTask() {
	time.Sleep(5 * time.Second)
}

func main2() {
	now := time.Now()
	for i := 10; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

	longAndHeavyTask()

	elapse := time.Since(now)

	fmt.Println(elapse.String())
}
