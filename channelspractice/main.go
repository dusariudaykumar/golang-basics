package main

import "fmt"

// pipelines

func main() {
	nums := []int{2, 4, 3, 5, 1}

	stageOne := sliceToChannel(nums)

	stageTwo := sq(stageOne)

	for i := range stageTwo {
		fmt.Println(i)
	}
}

func sliceToChannel(num []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range num {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}
