package main

import "fmt"

func main() {
	slice := []string{"a", "c", "d"}
	fmt.Println(slice, len((slice)))
	slice = append(slice, "e")
	fmt.Println(slice, len((slice)))
}
