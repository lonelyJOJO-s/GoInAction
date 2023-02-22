package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func main() {

// 	wg.Add(1)
// 	baton := make(chan int)
// 	go Runner(baton)

// 	baton <- 1
// 	wg.Wait()
// 	fmt.Println("比赛结束!")

// }

// func Runner(baton chan int) {
// 	var newRunner int
// 	runner := <-baton
// 	fmt.Printf("runner %d start running with abton\n", runner)

// 	if runner != 4 {
// 		newRunner = runner + 1
// 		go Runner(baton)
// 	}

// 	time.Sleep(time.Millisecond * 100)

// 	if runner == 4 {
// 		fmt.Printf("runner %d finished\n", runner)
// 		wg.Done()
// 		return
// 	}

// 	fmt.Printf("pass baton from runner%d to runner%d", runner, newRunner)
// 	baton <- newRunner

// }
