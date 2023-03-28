package main

import "fmt"

var (
	a    string
	done bool
)

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	for i := 1; i < 100; i++ {
		go setup()
		for !done {
		}
		fmt.Println(a)
	}

}
