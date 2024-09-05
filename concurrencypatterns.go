package main

import (
	"fmt"
	"sync"
	"time"
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
Eaxmple: for-select-done channel pattern to prevent go routine leak
by making go routine can exit based on done channel closed by the parent or caller
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

var mutex sync.Mutex

func testWaitGroup(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()
	mutex.Lock()
	*result = append(*result, data*2)
	mutex.Unlock()
}

func main() {
	mySlice := []int{1, 2, 3, 4}
	done := make(chan bool)

	//Example: Multi-stage synchronisation using go channel & done channale
	firstStageChannel := preparedataStage(mySlice)
	secondStageChannel := squaredataStage(firstStageChannel, done)

	for result := range secondStageChannel {
		fmt.Println(result)
	}

	close(done)

	// Another example of done channel
	done2 := make(chan bool)
	go func(done2 <-chan bool) {
		for {
			select {
			case <-done2:
			default:
				fmt.Println("printing until done...")
			}
		}
	}(done2)

	time.Sleep(1 * time.Second)

	close(done2)

	//Example: WaitGroup & mutex - wg.Add, wg.Done, wg.Wait
	var wg sync.WaitGroup
	input := []int{1, 2, 3, 4, 5}
	result := []int{}
	for _, data := range input {
		wg.Add(1)
		go func() {
			testWaitGroup(&wg, &result, data)
		}()
	}

	wg.Wait()
	fmt.Println(result)

	//TODO: Confinement

}
