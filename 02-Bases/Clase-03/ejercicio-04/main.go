package main

import "fmt"

const (
	minimumOpp = "minimum"
	maximumOpp = "maximum"
	averageOpp = "average"
)

func main() {
	operation, err := calculate(averageOpp)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(operation(1, 2, 3, 4, 5))
}

func calculate(operation string) (func(values ...int) int, error) {
	switch operation {
	case minimumOpp:
		return minimum, nil
	case maximumOpp:
		return maximum, nil
	case averageOpp:
		return average, nil
	default:
		return nil, fmt.Errorf("operation %s not supported", operation)
	}
}

func minimum(values ...int) int {
	var min int = values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func maximum(values ...int) int {
	var max int = values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func average(values ...int) int {
	var suma int = 0
	var cant int = 0
	for _, value := range values {
		suma += value
		cant++
	}
	return suma / cant
}
