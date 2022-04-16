package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string // поле изменяется на неэкспортируемое
	Date
}

func (e *Event) Title() string { //get метод
	return e.title
}

func (e *Event) SetTitle(title string) error { //set-метод
	if utf8.RuneCountInString(title) > 30 { //подсчет колличества символов в названии
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
