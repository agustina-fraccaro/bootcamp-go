package main

import "fmt"

func main() {
	var edad int
	var empleado bool
	var antiguedad int
	var sueldo float64

	fmt.Scan(&edad, &empleado, &antiguedad, &sueldo)
	fmt.Println(edad, empleado, antiguedad, sueldo)
	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Se te otorga préstamo sin intereses")
		} else {
			fmt.Println("Se te otorga préstamo con intereses")
		}
	} else {
		fmt.Println("No se te otorga préstamo")
	}

}
