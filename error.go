package main

import (
	"fmt"
)

func main() {
	err := fmt.Errorf("a height of %0.2f is invalid", -2.33333) //vozvrahaet zna4eni oshibki
	fmt.Println(err)                                            // vivod sombcheniya ob oshibki 1
	fmt.Println(err.Error())                                    //vivod sombcheniya ob oshibki 2
}
