package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("mytxt.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fileInfo.Size())
}
