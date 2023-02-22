package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var (
// 	counter int
// 	mutex   sync.Mutex
// 	wp      sync.WaitGroup
// )

// func main() {

// 	wp.Add(2)

// 	go access()
// 	go access()

// 	wp.Wait()
// 	fmt.Print(counter)
// }

// func access() {

// 	defer wp.Done()

// 	for count := 0; count < 2; count++ {
// 		mutex.Lock()
// 		{
// 			value := counter
// 			runtime.Gosched()
// 			value++
// 			counter = value
// 		}
// 		mutex.Unlock()
// 	}
// }
