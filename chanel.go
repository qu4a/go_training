package main

import "fmt"

func greeting(myChanel chan string) {
	myChanel <- "HI"
}

func main() {
	myChanel := make(chan string) //короткое объявление канала
	go greeting(myChanel)
	receiceValue := <-myChanel
	fmt.Println(receiceValue)
}
