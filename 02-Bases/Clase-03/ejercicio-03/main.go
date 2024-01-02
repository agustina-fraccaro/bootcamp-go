package main

import "fmt"

func main() {
	salario := calcularSalario(5, "B")

	fmt.Println(salario)
}

func calcularSalario(cantHoras int, categoria string) float64 {
	var salario float64 = 0
	switch categoria {
	case "C":
		salario += float64((1000 * cantHoras))
	case "B":
		salario += float64(1500 * cantHoras)
		salario += (salario * 0.2)
	case "A":
		salario += float64(3000 * cantHoras)
		salario += (salario * 0.5)
	}

	return salario
}
