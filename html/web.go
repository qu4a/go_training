package main

import (
	"log"
	"net/http"
)

func viewHandler(write http.ResponseWriter, request *http.Request) {
	/*
	   write http.ResponseWriter - значение для обновление ответа, которое будет отправлено браузеру
	   request *http.Request - значение, представляющее запрос от браузера
	*/
	message := []byte("Hello, web!") //добавляем строку Hello,web в ответ преварительно преобразовав в срез байтов
	_, err := write.Write(message)   //строка добавлется в ответ
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	http.HandleFunc("/hello", viewHandler)
	/*
	   /hello - если был получен запрос на URL адресс
	   viewHandler - вызывается функция, которая генерирует ответ
	*/
	err := http.ListenAndServe("localhost:8080", nil) //обрабатывает запрос браузера и отвечает на него
	log.Fatal(err)
}
