package main

import (
	"errors"
	"fmt"
)

type ErrorSalaryLessThanMinimum struct {
	msg string
}

func (e ErrorSalaryLessThanMinimum) Error() string {
	return e.msg
}

func CheckSalary(salary int) (string, error) {
	if salary <= 10000 {
		return "", ErrorSalaryLessThanMinimum{"Error: the salary entered does not reach the taxable minimum"}
	} else {
		return "Must pay tax", nil
	}
}

func main() {
	var salary int = 1000000

	msg, err := CheckSalary(salary)

	errTarget := ErrorSalaryLessThanMinimum{"Error: the salary entered does not reach the taxable minimum"}

	if errors.Is(err, errTarget) {
		fmt.Print(err)
		return
	}

	fmt.Println(msg)
}
