package main

import "fmt"

func main() {
	//severalNumbs(1)
	//severalNumbs(1, 2, 3, 4, 5)
	mix(1, true, "test")
	mix(12, false, "again", "mote string", "more-more string")

}
func mix(one int, flag bool, strings ...string) {
	fmt.Println(one, flag, strings)
}
