package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Edad Benjamin:", employees["Benjamin"])

	// saber cuantos empleados son mayores a 21 años
	var cantMayores21 int = 0
	for _, value := range employees {
		if value > 21 {
			cantMayores21++
		}
	}
	fmt.Println("Cantidad de empleados mayores de 21 años:", cantMayores21)

	// agregar un empleado llamado Federico que tiene 25 años
	employees["Federico"] = 25

	// eliminar el empleado Pedro
	delete(employees, "Pedro")

	fmt.Println(employees)

}
