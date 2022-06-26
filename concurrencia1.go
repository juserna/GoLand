package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("number: %d \n", i)
	}

	fmt.Println("----------------")
	i := 0
	for ; i < 10; i++ {
		go ping(i)
	}
	time.Sleep(2 * time.Second)

}

func ping(i int) {
	fmt.Println(fmt.Sprintf("number: %d", i))
}
