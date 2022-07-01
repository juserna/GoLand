package main

import "fmt"

func main() {
	fmt.Println("Calling worker")
	done := make(chan bool)

	go worker(done)

	fmt.Println(<-done)

}

func worker(done chan bool) {
	fmt.Println("inside the worker")

	done <- true

}

func bufferedChannel() {
	c := make(chan string, 2)

	c <- "one"
	c <- "two"
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func firstExample() {
	c := make(chan string)

	go func() {
		c <- "Hello from Channel"
	}()

	res := <-c
	fmt.Println(res)
}
