package main

import (
	"errors"
	"fmt"
)

func CheckSalary(salary int) (string, error) {
	if salary <= 10000 {
		return "", errors.New("Error: the salary entered does not reach the taxable minimum")
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
