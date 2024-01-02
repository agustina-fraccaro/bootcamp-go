package main

import "fmt"

func main() {
	impuesto := calcularImpuesto(50001)
	fmt.Println(impuesto)
}

func calcularImpuesto(salario float64) float64 {
	var impuesto float64 = 0
	if salario > 50000 {
		impuesto += (salario * 0.17)
	}

	if salario > 150000 {
		impuesto += (salario * 0.10)
	}

	return impuesto
}
