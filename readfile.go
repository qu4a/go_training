package main //чтение из файла (выводит строки)

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt") //открываем файл (функция возращает ошибку err)
	if err != nil {
		log.Fatal(err)
	}
	scaner := bufio.NewScanner(file) // для file создает новое значение scaner
	for scaner.Scan() {              //читаем строку из файла
		fmt.Println(scaner.Text()) //выводим строки
	}
	err = file.Close() //закрываем файл
	if err != nil {    // если при закрытия файла произошла ошибка сообщить о ней и завершить работу
		log.Fatal(err)
	}
	if scaner.Err() != nil { //Если при сканировании файла произошла ошибка, сообщить о ней и завершить работу
		log.Fatal(scaner.Err())
	}
}
