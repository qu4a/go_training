package main

import (
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.getFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range number { //перебор всех чисел в массиве
		sum += value

	}
	fmt.Printf("Sum is %0.1f\n", sum)

	count := float64(len((number))) //получить значение длины массивы в int и переобразовать в значение float
	fmt.Printf("Average is %0.1f\n", sum/count)
}
