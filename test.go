package main

import (
	"fmt"
	"headfirst/magazine"
)

func main() {
	var employye magazine.Employye
	employye.Name = "Jon Carr"
	employye.Salary = 60000
	fmt.Println(employye.Name)
	fmt.Println(employye.Salary)
	var address magazine.Address
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"
	fmt.Println(address)
	var s magazine.Subscriber
	s.HomeAddress.City = "68111"
	fmt.Println(s.HomeAddress.City)

}
