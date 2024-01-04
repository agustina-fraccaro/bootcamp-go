package person

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Position:", e.Position)
	fmt.Println("Person ID:", e.Person.ID)
	fmt.Println("Person Name:", e.Person.Name)
	fmt.Println("Person DateOfBirth:", e.Person.DateOfBirth)
}
