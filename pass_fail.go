package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var status string
	fmt.Print("Endter grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(input)
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Print(reflect.TypeOf(grade))

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Print("A grade ", grade, " is ", status)
}
