package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func main() {

// 	wg.Add(2)

// 	court := make(chan int)

// 	go player("james", court)
// 	go player("lisa", court)

// 	// 发球
// 	court <- 1

// 	wg.Wait()

// }

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// func player(name string, court chan int) {
// 	defer wg.Done()

// 	for {
// 		ball, ok := <-court //注意语法
// 		if !ok {            // 如果通道被关闭
// 			fmt.Println(name + "赢了")
// 			return
// 		}
// 		n := rand.Intn(100)
// 		if n%13 == 0 {
// 			fmt.Println(name + "输了")
// 			close(court)
// 			return
// 		}
// 		fmt.Println(name+"击球成功 ", ball)
// 		ball++
// 		court <- ball

// 	}
// }
