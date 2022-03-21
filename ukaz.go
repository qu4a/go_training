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

	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 8          //новое значение присваевается пересменной на которую ссылается указатель myInt
	fmt.Println(*myIntPointer) // выводится значение пересенной на которую ссылается указатель
	fmt.Println(myInt)         //значение переменной на которое ссылается указатель
}
