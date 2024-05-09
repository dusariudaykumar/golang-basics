package main

import (
	"fmt"
	"sync"
)

func main() {
	// A channel is a pipe between goroutines to synchronize execution and communicate by sending/receiving data.

	/*
		-  Unbuffered Channels:
			Unbuffered channels are the simplest form of channels.
			When you create an unbuffered channel, it has a capacity of zero.
			This means that every send operation on the channel blocks until another goroutine is ready to receive the value.
		    Likewise, every receive operation blocks until another goroutine is ready to send a value.

	*/
	fmt.Println("Channels in golang")

	myChannel := make(chan int) // Creating an unbuffered channel

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		// pushing value into channel
		myChannel <- 10

		close(myChannel)
	}(myChannel, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		// receiving value from channel
		value := <-myChannel

		fmt.Println("Value -", value) //output: 10
	}(myChannel, wg)

	wg.Wait()
	x := <-myChannel
	fmt.Println(x) // output: 0 because the channel is closed due to that its sends 0

	// to check if channel is open or not
	value, isChannelOpened := <-myChannel
	fmt.Printf("Is channel Open: %v\n value from the channel: %v\n", isChannelOpened, value)
	// BufferedChannel()
	goroutines()

}

func BufferedChannel() {
	/*
		Buffered channels, on the other hand, have a specified capacity greater than zero.
		This means that they can hold a certain number of values before blocking send operations.
		Buffered channels decouple the sender and receiver, allowing for asynchronous communication.
	*/

	ch := make(chan int, 2) // Creating a buffered channel with a capacity of 2

	ch <- 1 // Sending the value 1 to the channel
	ch <- 2 // Sending the value 2 to the channel

	x := <-ch      // Receiving the value from the channel
	fmt.Println(x) // Output: 1

	y := <-ch      // Receiving the value from the channel
	fmt.Println(y) // Output: 2
}

func goroutines() {

	myChan := make(chan string)
	anotherChan := make(chan string)

	go func() {
		myChan <- "data"
	}()

	go func() {
		anotherChan <- "Hello World!"
	}()

	select {
	case data := <-myChan:
		fmt.Println("DATA: ", data)
	case value := <-anotherChan:
		fmt.Println("Msg: ", value)
	}

}
