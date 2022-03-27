package main

import "fmt"

func main() {
	var sum float64 = 0.0
	number := [3]float64{2.54, 3.56, 4.09}
	for _, value := range number { //перебор всех чисел в массиве
		sum += value
	}
	fmt.Printf("Sum is %0.1f\n", sum)

	count := float64(len((number))) //получить значение длины массивы в int и переобразовать в значение float
	fmt.Printf("Average is %0.1f\n", sum/count)
}
