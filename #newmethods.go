package main

import "fmt"

type MyType string // определяется новый тип

func (m MyType) SayHi() { //m - параметр получателя, MyType-определяется метод
	fmt.Println(m)
}
func main() {
	value := MyType("a mytype value") //создаем значение MyType
	value.SayHi()                     //Вызывем SayHi для этого значения

}
