package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	numberGoRoutines = 4
	taskLoad         = 10
)

func main() {

	wg.Add(numberGoRoutines)
	tasks := make(chan string, taskLoad)

	for i := 0; i < numberGoRoutines; i++ {
		go work(tasks, i+1)
	}

	for i := 0; i < taskLoad; i++ {
		tasks <- fmt.Sprintf("Task: %d", i+1)
	}

	close(tasks)

	wg.Wait()
}

func work(tasks chan string, id int) {

	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Println("finished!")
			return
		}
		fmt.Println(id, "is working with ", task)

		sleep := rand.Int31n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Println(id, "complete ", task)
	}

}
