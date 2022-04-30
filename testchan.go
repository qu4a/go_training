package main

import "fmt"

func odd(chanel chan int) {
	chanel <- 1
	chanel <- 3

}
func even(chanel chan int) {
	chanel <- 2
	chanel <- 4
}
func main() {
	chanelA := make(chan int)
	chanelB := make(chan int)
	go odd(chanelA)
	go even(chanelB)
	fmt.Println(<-chanelA)
	fmt.Println(<-chanelB)
	fmt.Println(<-chanelA)
	fmt.Println(<-chanelB)
}
