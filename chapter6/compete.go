package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var (
// 	counter int
// 	wg      sync.WaitGroup
// )

// func compete() {
// 	wg.Add(2)

// 	go access()
// 	go access()

// 	wg.Wait()
// 	fmt.Print("finished!", counter)

// }

// /*
// 出现混乱的原因是否为运行过程中调度器自动地进行了goroutine的切换
// */
// func access() {

// 	defer wg.Done()

// 	for count := 0; count < 2; count++ {
// 		value := counter
// 		fmt.Println(value)
// 		runtime.Gosched() // 退出当前线程
// 		fmt.Println(value)
// 		value++
// 		counter = value
// 	}
// }
