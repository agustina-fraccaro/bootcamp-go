package main

import (
	"content/person"
)

func main() {
	person := person.Person{ID: 1, Name: "Agustina", DateOfBirth: "01/01/2000"}
	employee := person.Employee{ID: 1, Position: "Developer", Person: person}

	employee.PrintEmployee()
}
