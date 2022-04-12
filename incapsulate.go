package main

import (
	"errors" //позволяет создатавать значения ошибок
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error { //добавлем возвращаемый тип error
	if year < 1 {
		return errors.New("invalid year") //Если uод недействителен, возвращается признак ошибки
	}
	d.Year = year //в противном случае присваеваем значение поля
	return nil    //возвращаем знание nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}
func (d *Date) SetDay(day int) error {

	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}
	err := date.SetYear(-1) //сохраняет любую ошибку
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year)
	/*
		date := Date{}
		err := date.SetDay(27)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(date.Day)
		date := Date{}
		err := date.SetMonth(5)
		if err != nil {
		log.Fatal(err)
		}
		fmt.Println(date.Month)
	*/
}
