package main

import (
	"fmt"
)

type TaxError struct {
	msg string
}

func (e *TaxError) Error() string {
	return fmt.Sprintln(e.msg)
}

func CheckSalary(salary int) (string, error) {
	if salary < 15000 {
		return "", &TaxError{"Error: the salary entered does not reach the taxable minimum"}
	} else {
		return "Must pay tax", nil
	}
}

func main() {
	var salary int
	fmt.Scan(&salary)

	msg, err := CheckSalary(salary)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(msg)
}
