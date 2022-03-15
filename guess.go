package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	second := time.Now().Unix() //data i vremya v celom chisle
	rand.Seed(second)           //generator sluchainih chisel
	target := rand.Intn(100) + 1
	fmt.Print(target)
}
