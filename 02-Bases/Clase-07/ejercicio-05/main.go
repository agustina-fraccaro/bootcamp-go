package main

import "fmt"

func main() {
	salary, err := CalculateSalary(90, 10000)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(salary)
}

func CalculateSalary(workHours int, hourCost float64) (float64, error) {
	if workHours < 80 || workHours < 0 {
		return 0, fmt.Errorf("Error: the worker cannot have worked less than 80 hours per month")
	}

	salary := float64(workHours) * hourCost

	if salary >= 150000 {
		salary -= (salary * 0.1)
	}

	return salary, nil
}
