package main

import "fmt"

func main() {
	testFunc("hello", 5)
}

func testFunc(word string, a int) {
	for i := 0; i < a; i++ {
		fmt.Println(word)
	}
}
