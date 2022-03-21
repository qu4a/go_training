package main

import "fmt"

func creatPoint() *float64 { //бъявляем, что функция возвращает указатель на float64
	var myFloat = 99.5
	return &myFloat //возвращает указатель данного типа
}

func main() {
	var myFloatPointer *float64 = creatPoint() // возращает указатель в переменную

	fmt.Println(*myFloatPointer)
	fmt.Print(&myFloatPointer)

	var myInt int = 42
	myIntPointer := &myInt
	fmt.Print(*myIntPointer)

}
