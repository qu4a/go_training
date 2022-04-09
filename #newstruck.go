package main

import "fmt"

type subscriber struct { //создается структрура subscriber
	name   string
	rate   float64
	active bool
}

func printInfo(s subscriber) { // объявлется один параметр s с типом subscriber
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}
func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 4.99
	s.active = true
	return s

}
func main() {
	subscriber1 := defaultSubscriber("Anna")
	fmt.Println(subscriber1)
	printInfo(subscriber1)
	subscriber1.rate = 5.99
	printInfo(subscriber1)
	subscriber1.active = false
	printInfo(subscriber1)

}
