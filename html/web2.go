package main

import (
	"log"
	"net/http"
)

func write(write http.ResponseWriter, message string) {
	_, err := write.Write([]byte(message)) //строка перробразует в серз байтов и пишется в ответ
	if err != nil {
		log.Fatal(err)
	}
}
func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!") //строка пишется в ответ
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!") //строка пишется в ответ
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!") //строка пишется в ответ
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
