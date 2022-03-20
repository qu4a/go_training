package main

import (
	"fmt"
)

func main() {
	var myInt int
	fmt.Println(&myInt)   //адресс переменной myInt
	var myIntPointer *int // объявление переменной содерж указ на int
	myIntPointer = &myInt // указатель присвается переменной
	fmt.Println(myIntPointer)

	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)

	//короткое объвление указателя
	var myBool bool
	fmt.Println(&myBool)
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
}
