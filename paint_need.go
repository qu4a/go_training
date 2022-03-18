package main

import "fmt"

func main() {
	paintNeed(2.5, 3.0)
	paintNeed(10.0, 3.0)
	paintNeed(5.7, 4.8)
}
func paintNeed(a float64, b float64) {
	area := a * b
	fmt.Printf(" %.3f need\n", area)
}
