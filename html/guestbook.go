package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) { //код сообщений об ошибках перемещается в эту функцию
	if err != nil {
		log.Fatal(err)
	}
}
func executeTemplate(text string, data interface{}) { //data значения передаются методу Execute шаблона
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	/*
		placeholder := []byte("signature list goes here")
		_, err := writer.Write(placeholder)
		check(err)
	*/
	html, err := template.ParseFiles("view.html") //сожержимое view.html используется для создания новго значения Template
	check(err)
	err = html.Execute(writer, nil) //содердимое шаблона записывается в ResponseWriter
	check(err)

func main() {
	http.HandleFunc("/guesbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err) //ошибка никогда не равно nil, по этому не вызываем check
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.5)
}
