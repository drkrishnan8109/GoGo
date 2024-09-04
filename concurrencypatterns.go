package main

import (
	"fmt"
)

func preparedataStage(data []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, c := range data {
			out <- c
		}
		close(out)
	}()
	return out
}

/*
for-select-done channel pattern to prevent go routine leak
where go routine can exit based on done channel closed by the parent or caller
*/
func squaredataStage(in <-chan int, done <-chan bool) <-chan int {
	out := make(chan int)
	go func(done <-chan bool) {
		for n := range in {
			select {
			case <-done:
			default:
				out <- n * n
			}
		}
		close(out)
	}(done)
	return out
}

func main() {
	mySlice := []int{1, 2, 3, 4}
	//firstStageChannel := make(chan int)
	//secondStageChannel := make(chan int)
	done := make(chan bool)

	firstStageChannel := preparedataStage(mySlice)
	secondStageChannel := squaredataStage(firstStageChannel, done)

	for result := range secondStageChannel {
		fmt.Println(result)
	}

	close(done)
}
