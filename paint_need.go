package main

import "fmt"

func paintNeed(a float64, b float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("Cann't be %0.1f <0", a)
	}
	if b < 0 {
		return 0, fmt.Errorf("Cann't be %0.1f <0", b)
	}
	area := a * b
	return area, nil

}
func main() {
	amount, err := paintNeed(2.0, -4.5)
	fmt.Println(err)
	fmt.Printf("1st need is %0.1f\n", amount)

}
