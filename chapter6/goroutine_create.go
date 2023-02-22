package main

/*
并发显示大小写字母
*/

import (
	"fmt"
	"sync"
)

func test() {
	// runtime.GOMAXPROCS(1) // set max logic processor : 1
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}

	}()

	go func() {

		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("finished!")
}
