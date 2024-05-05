package main

import (
	"fmt"
	"sync"
)

// refernece blog https://kamnagarg-10157.medium.com/understanding-mutex-in-go-5f41199085b9

var counter = 0
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup

	expectedCounter := 1000

	for i := 0; i < expectedCounter; i++ {
		wg.Add(1)
		// go increment(&wg)
		go incrementWithMutex(&wg)
	}

	wg.Wait()
	fmt.Println("Expected Counter:", expectedCounter)
	fmt.Println("Actual Counter:", counter)

	if expectedCounter != counter {
		fmt.Println("Race condition detected!")
	} else {
		fmt.Println("No race condition detected.")
	}

	/*
		output : Expected Counter: 1000
				Actual Counter: 972
				Race condition detected!

		Here multiple goroutines are simultaneously reading and updating the value of the counter without any proper synchronization.
	*/
	/*
		We need to use the sync.Mutex type to prevent multiple goroutines from accessing counter at the same time:

		after using incrementWithMutex()
		output: Expected Counter: 1000
				Actual Counter: 1000
				No race condition detected.

	*/
}

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()

//		counter++
//	}

func incrementWithMutex(wg *sync.WaitGroup) {
	/*
	   Note : Whenever you call the Lock method, you must ensure that Unlock is eventually called,
	        otherwise any goroutine trying to acquire the same lock will be blocked forever.
	*/
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

/*
	Mutexes play a crucial role in concurrent programming to ensure proper synchronization and
	prevent race conditions. By using the sync.Mutex type in Go,
	developers can protect shared resources and control access to critical sections of code.
*/
