package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	responseSize("https://golang.org/")
	responseSize("https://golang.org/doc")
}

func responseSize(url string) {
	fmt.Println("Getting:", url)
	response, err := http.Get(url) //вызываем http.Get с URL адресом нужной страницы
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()                //освобождает сетевое подключение после выхода их main
	body, err := ioutil.ReadAll(response.Body) //читает все данные из ответа
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body)) //размер среза в байтах равен размеру страницы
}
