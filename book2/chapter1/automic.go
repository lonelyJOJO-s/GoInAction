// 模拟最简单的原子操作
package main

import (
	"sync"
	"sync/atomic"
)

var Total struct {
	sync.Mutex
	value int
}

func worker1(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 100; i++ {
		Total.Lock()
		Total.value++
		Total.Unlock()
	}

}

var total uint32

// better way
func worker2(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 100; i++ {
		atomic.AddUint32(&total, 1)
	}

}

// 单例

type signleton struct{}

var (
	once     sync.Once
	instance *signleton
)

func Instance() *signleton {
	once.Do(func() {
		instance = &signleton{}
	})
	return instance
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go worker1(&wg)
// 	go worker1(&wg)

// 	wg.Wait()
// 	fmt.Print(Total.value)
// }
