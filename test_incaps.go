package main // с инкапсуляцией

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	location := geo.Landmark{}
	err := location.SetName("eqw")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(20)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-170)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
	fmt.Println(location.Name())
}
