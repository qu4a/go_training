package main // переработка с учетом каналов

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	size int
}

func responseSize(url string, chanel chan Page) { //Канал, передоваемый responSize, будет использоваться Page, a не int
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
	//fmt.Println(len(body)) //размер среза в байтах равен размеру страницы
	chanel <- Page{URL: url, size: len(body)}
}

func main() {
	pages := make(chan Page) // создаем канал для значений int
	urls := []string{"https://golang.org/",
		"https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, pages) // вызывает responsSize для каждого url
	}
	for i := 0; i < len(urls); i++ { //получение данных из канала по каждой оправки, выполненой responSize
		page := <-pages
		fmt.Println(page.URL, page.size)
	}
	//go responseSize("https://golang.org/", size) //используем go команды для поддержание в программе параллелизма и конкурентности
	//go responseSize("https://golang.org/doc", size)
	//time.Sleep(5 * time.Second)
	//fmt.Println(<-size)
	//fmt.Println(<-size)
}
