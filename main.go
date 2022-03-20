package main

import "fmt"

func double(number float64) float64 {
	return number * 2
}

func main() {
	fmt.Println(double(4))
	tst := double(7)
	fmt.Println(tst)
}
