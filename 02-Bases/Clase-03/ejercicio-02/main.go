package main

import "fmt"

func main() {
	promedio := calcularPromedio(10, 9, 9, 9)

	fmt.Println(promedio)
}

func calcularPromedio(notas ...int) int {
	var suma int = 0
	var cantNotas int = 0
	for _, value := range notas {
		if value > 0 {
			suma += value
			cantNotas++
		}
	}

	return suma / cantNotas
}
