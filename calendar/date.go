package calendar //икапсуляция - сокрыттие данные

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

//set-меотды для присваивания значений поля структур или переменной (сеттеры) (защита структур от неверных значений) (также должны работать с указателем*)

func (d *Date) SetYear(year int) error { //добавлем возвращаемый тип error
	if year < 1 {
		return errors.New("invalid year") //Если uод недействителен, возвращается признак ошибки
	}
	d.year = year //в противном случае присваеваем значение поля
	return nil    //возвращаем знание nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {

	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

//get-методы (геттеры) получения значения поля структуры или переменной (также должны искользоватся с указателями*)
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}
