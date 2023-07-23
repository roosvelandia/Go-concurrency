package main

import "fmt"

func main() {
	// channel with buffer
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	// it will exceed the capacity of the channel (3)
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
