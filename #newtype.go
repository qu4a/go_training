package main

import "fmt"

//определяем два новых типа, оба используют базу float64
type Litets float64
type Gallons float64

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}
func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	// короткие объявдления
	carFuel := Gallons(1.2)
	busFuel := Litets(4.5)
	carFuel += ToGallons(Litets(40.0))
	busFuel = +ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)

}
