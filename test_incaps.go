package main // с инкапсуляцией

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(20)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-170)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
