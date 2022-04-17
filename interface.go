package main

import (
	"fmt"
	"mypkg"
)

func main() {
	var value mypkg.MyInterface //объявление переменной и использованием типа Interface
	value = mypkg.MyType(5.7)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.4)
	fmt.Println(value.MethodWithReturnValue())

}
