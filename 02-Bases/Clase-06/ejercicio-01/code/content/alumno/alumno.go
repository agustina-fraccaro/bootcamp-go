package alumno

import "fmt"

type Alumno struct {
	Name            string
	LastName        string
	DNI             int
	DateOfAdmission string
}

func (a Alumno) Detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha de ingreso: %s\n", a.Name, a.LastName, a.DNI, a.DateOfAdmission)
}
