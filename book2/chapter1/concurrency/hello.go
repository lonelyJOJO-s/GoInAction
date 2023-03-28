package main

// 使用mutex实现同步
// func main() {
// 	var mu sync.Mutex

// 	mu.Lock()

// 	go func() {
// 		fmt.Println("hello, world!")
// 		mu.Unlock()
// 	}()

// 	mu.Lock()

// }

// 使用通道实现同步

// func main() {
// 	c := make(chan bool)

// 	go func() {
// 		fmt.Print("hello, world")
// 		c <- true
// 	}()

// 	<-c
// }

// func main() {
// 	done := make(chan int, 10)

// 	for i := 0; i < cap(done); i++ {
// 		j := i
// 		go func() {
// 			fmt.Println(j)
// 			done <- 1
// 		}()
// 	}

// 	for i := 0; i < cap(done); i++ {
// 		<-done
// 	}
// }

// 对于这种需要完成n个goroutine再进行下一步的同步操作，用waitGroup

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			fmt.Print("hello, world")
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// }
