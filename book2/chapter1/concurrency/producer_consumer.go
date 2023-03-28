package main

import (
	"fmt"
)

func producer(factor int, out chan<- int) {

	for i := 0; ; i++ {
		out <- factor * i
	}

}

func comsumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

// func main() {
// 	store := make(chan int, 56)

// 	go producer(3, store)
// 	go producer(5, store)
// 	go comsumer(store)

// }
