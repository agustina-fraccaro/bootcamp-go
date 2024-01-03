package main

import "fmt"

func main() {
	var palabra string
	var cantLetras int = 0
	fmt.Scan(&palabra)

	for i := 0; i < len(palabra); i++ {
		cantLetras++
		fmt.Printf("%q", palabra[i])
	}

	fmt.Println("Cantidad de letras:", cantLetras)

}
