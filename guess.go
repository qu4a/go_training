package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	second := time.Now().Unix() //data i vremya v celom chisle
	//fmt.Println(second)
	rand.Seed(second) //func generator sluchainih chisel
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)
	success := false //peremenaya dlya kollichestva popitok

	reader := bufio.NewReader(os.Stdin) //befereziruem vvod s klaviaturi
	for guesses := 0; 10 > guesses; guesses++ {

		fmt.Println("You have", 10-guesses, "guesses left") //cycle c4itaet colli4estvo oshibok

		fmt.Print("Make a giess:")
		input, err := reader.ReadString('\n') // 4itat dannie polzovatelya do najatiya Enter
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // udalenie simvola novoi stroki
		guess, err := strconv.Atoi(input) // preobrasovanie stroki v celoe 4islo
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(guess)
		//fmt.Print(reflect.TypeOf(guess))
		if guess < target {
			fmt.Println("Low")
		} else if guess > target {
			fmt.Println("High")
		} else {
			fmt.Println("Good job")
			success = true
			break
		}

	}
	if !success {
		fmt.Println("It was:", target) //esli igrok ne ugadal
	}
}
