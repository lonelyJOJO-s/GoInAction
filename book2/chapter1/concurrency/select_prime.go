package main

import "fmt"

// 生成自然数
func generateNetural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 质数过滤器
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := generateNetural()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Println(i+1, " ", prime)
		ch = PrimeFilter(ch, prime)
	}

}
