package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int) //создаем карту с участниками голосования имена и счетсчик голосов
	for _, line := range lines {   //включаем счетчик вхождений
		counts[line]++
	}
	for name, count := range counts { //цикл для мапы с выводом
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
