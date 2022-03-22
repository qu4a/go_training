package main

import "fmt"

func main() {
	amout := 30
	double(&amout) //вместо значения переменной передеется указатель на него
	fmt.Println(amout)

}

func double(value *int) { // вместо занчения int получает указатель
	*value *= 2.0 //обновляет значение на которое ссылается указатель

}
