package main

import (
	"fmt"
	"math/rand"
)

/*
- Close channel, values returned in channel output
- Range channel
- len/cap channel
- Channels directions
- Channel zero value
*/

func main() {
	//rangeChannelsExample()
}

func rangeChannelsExample() {
	c := make(chan int, 3)

	go rangeChannels(c)

	for val := range c {
		fmt.Print(val)
		fmt.Print(" ")
	}

}

func rangeChannels(c chan int) {
	c <- 1
	fmt.Printf("len %d y cap %d \n", len(c), cap(c))
	fmt.Println()
	c <- 2
	fmt.Printf("len %d y cap %d \n", len(c), cap(c))
	fmt.Println()
	c <- 3
	fmt.Printf("len %d y cap %d \n", len(c), cap(c))
	fmt.Println()

	close(c)
}

func closeChannel() {
	randChannel := make(chan int)

	var count int
	go rands(randChannel)
	for {
		i, ok := <-randChannel

		count++
		if !ok {
			break
		}
		fmt.Printf("Index %d , value %d \n", count, i)
	}

	fmt.Println("out of the loop")
}

func rands(c chan int) {
	for i := 0; i < 10; i++ {
		c <- rand.Int()
	}

	close(c)
}
