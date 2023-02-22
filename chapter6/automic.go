package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"sync/atomic"
// )

// var (
// 	ctr int64
// 	wgp sync.WaitGroup
// )

// func main() {
// 	wgp.Add(2)

// 	go acc()
// 	go acc()

// 	wgp.Wait()
// 	fmt.Print("finished!", ctr)

// }

// /*
// 出现混乱的原因是否为运行过程中调度器自动地进行了goroutine的切换
// */
// func acc() {

// 	defer wgp.Done()

// 	for count := 0; count < 2; count++ {
// 		atomic.AddInt64(&ctr, 1)
// 		// 同时还有LoadInt64, StoreInt64等等原子操作
// 	}

// 	runtime.Gosched()
// }
