package main

import (
	"errors"
	"fmt"
)

var (
	ErrSalaryLessThanMinimum = errors.New("Error: the salary entered does not reach the taxable minimum")
)

func CheckSalary(salary int) (string, error) {
	if salary <= 10000 {
		return "", fmt.Errorf("%w. Error: the minimum taxable amount is 150,000 and the salary entered is: %d", ErrSalaryLessThanMinimum, salary)
	} else {
		return "Must pay tax", nil
	}
}

func main() {
	var salary int = 10

	msg, err := CheckSalary(salary)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(msg)
}
