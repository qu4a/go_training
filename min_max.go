package main // поиск максимального значения min и max

import (
	"fmt"
	"math"
)

func max(numbers ...float64) float64 { //получаем любое колличество аргументов в float64
	max := math.Inf(-1) //начинает с очень низкого значения
	for _, number := range numbers {
		if number > max {
			max = number // находит наибольшее значение среди аргументов

		}
	}
	return max

}
func main() {
	fmt.Println(max(20, 34, -10, 20, -8))
}
