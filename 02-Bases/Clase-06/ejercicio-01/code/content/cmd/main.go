package main

import (
	"content/alumno"
)

func main() {
	alumno := alumno.Alumno{Name: "Agustina", LastName: "Gonzalez", DNI: 12345678, DateOfAdmission: "01/01/2000"}

	alumno.Detalle()
}
