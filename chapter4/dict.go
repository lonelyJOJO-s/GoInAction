package main

import (
	"fmt"
)



func main() {
	dict := map[string]string{
		"red" : "1",
		"blue" : "2",
		"yellow": "3", 
	}
	for key, value := range dict {
		key = "xxx"
		fmt.Println(key, value)
	}
	
	for key, value := range dict {
		fmt.Println(key, value)
	}
}


