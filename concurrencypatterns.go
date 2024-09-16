package main

import (
	"fmt"
	"sync"
	"time"
)

// Go empty interface can take any data type because all datatypes in go implement zero interface
func PrintAnyType(i interface{}) {
	fmt.Println(i)
}

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
Eaxmple: for-select-case done channel pattern to prevent go routine leak
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

func testWaitGroupConfined(wg2 *sync.WaitGroup, resultIdx *int, data int) {
	defer wg2.Done()
	mutex.Lock()
	*resultIdx = data * 2
	mutex.Unlock()
}

func main() {
	mySlice := []int{1, 2, 3, 4}
	done := make(chan bool)

	//Example: Multi-stage synchronisation using go channel & done channal
	firstStageChannel := preparedataStage(mySlice)
	secondStageChannel := squaredataStage(firstStageChannel, done)

	for result := range secondStageChannel {
		PrintAnyType(result)
	}

	close(done)

	// Another example of done channel
	done2 := make(chan bool)
	go func(done2 <-chan bool) {
		for {
			select {
			case <-done2:
			default:
				PrintAnyType("printing until done...")
			}
		}
	}(done2)

	time.Sleep(1 * time.Second)

	close(done2)

	//Example: WaitGroup & mutex - wg.Add, wg.Done, wg.Wait
	// Here we have created a worker pool with 5 workers/thread working parallely on each input data
	var wg sync.WaitGroup
	input := []int{1, 2, 3, 4, 5}
	result := []int{}
	for _, data := range input {
		wg.Add(1)
		go testWaitGroup(&wg, &result, data)
	}

	wg.Wait()
	fmt.Println(result)

	//Example: Confinement
	//Instead of locking the entire datastructure or array, we can confine the lock to the minimum space needed to be locked
	//So, lock only the index box for each index, not the entire array!

	result2 := make([]int, len(input))
	var wg2 sync.WaitGroup
	for i, data := range input {
		wg2.Add(1)
		go testWaitGroupConfined(&wg2, &result2[i], data)
	}
	wg2.Wait()
	PrintAnyType(result2)

	//Sync Pool to efficiently reuse objects instead of burdening GC to clean

	//Worker pool to create workers instead of too mayn go routines
	int workerPool = 10;
	int jobcount = 1000;


}
