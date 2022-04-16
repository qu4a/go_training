package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Mom birhday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
	fmt.Println(event.Title())

}
