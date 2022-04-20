package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(filename string) (*os.File, error) { //Открывает файл и возвращает указатель на него, а также  начение обнаруженной ошибки.
	fmt.Println("Opening", filename)
	return os.Open(filename)
}
func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file) //добовление ключевого слова defer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
func main() {
	numbers, err := GetFloats(os.Args[1]) //аргумент командной строки используется как аргумент
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f\n", sum)
}
