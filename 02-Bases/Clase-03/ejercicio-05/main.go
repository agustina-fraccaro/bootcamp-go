package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

const (
	cantDog               = 10
	cantCat               = 5
	cantHamster   float64 = 0.25
	cantTarantula float64 = 0.15
)

func main() {
	animalDog, err := Animal(dog)
	animalCat, err := Animal(cat)
	animalHamster, err := Animal(hamster)
	animalTarantula, err := Animal(tarantula)

	if err != nil {
		fmt.Println(err)
		return
	}

	var amount float64 = 0
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(10)
	amount += animalTarantula(10)

	fmt.Println(amount)
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, fmt.Errorf("animal %s not found", animal)
	}
}

func animalDog(cant int) float64 {
	return float64(cantDog * cant)
}

func animalCat(cant int) float64 {
	return float64(cantCat * cant)
}

func animalHamster(cant int) float64 {
	return cantHamster * float64(cant)
}

func animalTarantula(cant int) float64 {
	return cantTarantula * float64(cant)
}
