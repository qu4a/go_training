package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct { //объявляем новую структуру т.к. метод Execute принимает одно значение
	SignatureCount int
	Signatures     []string
}

func getStrings(filename string) []string { //функция для чтения из файла
	var lines []string
	file, err := os.Open(filename) //открываем файл
	if os.IsNotExist(err) {        //если файл отсутствует, то вернуть ошибку вместо среза
		return nil
	}
	check(err)         //в случае любой другой ошибке, сообщить об ней и завершить работу
	defer file.Close() //полсе выхода из функции в любом случае закрыть файл
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

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

	signatures := getStrings("signatures.txt")
	//fmt.Printf("%#v\n", signatures)
	guestbook := Guestbook{ // новая структура Guestbook
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook) //содердимое шаблона записывается в ResponseWriter, а структура передается методу Execute значения Template
	check(err)
}
func newHandler(writer http.ResponseWriter, request *http.Request) { //функция для /new
	html, err := template.ParseFiles("new.html") //загружает содержимое /new как как текст шаблона
	check(err)
	err = html.Execute(writer, nil) //записать шаблон в ответ
	check(err)
}
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature") //получает значение поля формы signature
	/*
		_, err := writer.Write([]byte(signature)) //записывает значение поля в ответ
		check(err)
	*/
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature) //записывает текст в новой строке файла
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}
func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler) //обработчик запросов
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err) //ошибка никогда не равно nil, по этому не вызываем check
	//executeTemplate("Dot is: {{.}}!\n", "ABC")
	//executeTemplate("Dot is: {{.}}!\n", 123.5)
}
