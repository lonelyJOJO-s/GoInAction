package main

import "fmt"

func TrimSpace(s []byte) []byte {
	fmt.Print(s)
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	fmt.Print(s)
	fmt.Print(b)
	return b
}

// copy the head info.
func testSlice(a []int) []int {

	a = a[:len(a)-1]
	fmt.Print(a)
	return a
}

func twice(a []int) {
	for i := range a {
		a[i] *= 2
	}
}

// func main() {
// 	a := []int{1, 2, 3}
// 	testSlice(a)
// 	fmt.Print(a)
// 	var Add = func(a, b int) int {
// 		return a + b
// 	}
// 	x := Add(1, 2)
// 	fmt.Println(x)
// 	fmt.Println(a)
// 	fmt.Println(a)
// 	twice(a)
// 	fmt.Println(a)
// }
