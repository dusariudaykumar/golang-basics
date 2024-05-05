package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// go keyword used to initialize goroutine

	/*
		   In this example, first, the main goroutine starts.
		   Then it invokes the helloWorld() function, and the helloWorld goroutine starts.

		   After the helloWorld goroutine finishes its operation,
		   the main goroutine waits for 1 second and invokes the goodBye() function.

		   Without time.Sleep():
		   output: Good bye!

		   After adding time.Sleep(),
		   the helloworld goroutine is able to finish its execution before main exits:
		   output: Hello World
				   Good bye!
	*/

	// go helloWorld()

	// time.Sleep(1 * time.Second)
	// goodBye()

	websites := []string{
		"https://google.com",
		"https://linkedin.com",
		"https://github.com",
		"https://everbee.io",
		"https://app.everbee.io",
	}

	// other example using sync.WaitGroup
	var wg sync.WaitGroup

	for _, endpoint := range websites {
		go webSiteStatus(endpoint, &wg)
		wg.Add(1)
	}
	// Here it works same as above example, but it doesn't block the main for 2 seconds.
	/*
			wg.Add(int): This method indicates the number of goroutines to wait.
			In the above code, I have added inside for loop so that it can wait for all goroutines which is initialized inside loop.
		    Hence the internal counter wait becomes 5 (because lenght of my slice is 5).

			wg.Wait(): This method blocks the execution of code until the internal counter becomes 0.

			wg.Done(): This will reduce the internal counter value by 1.

			NOTE: If a WaitGroup is explicitly passed into functions, it should be added by a pointer.
	*/
	wg.Wait()
}

// func helloWorld() {
// 	fmt.Println("Hello World")
// }

// func goodBye() {
// 	fmt.Println("Good bye!")
// }

// Other examples

func webSiteStatus(endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Something went wrong!")
	} else {
		fmt.Printf("Status code: %d for %s\n", res.StatusCode, endpoint)
	}

	defer res.Body.Close()
}
