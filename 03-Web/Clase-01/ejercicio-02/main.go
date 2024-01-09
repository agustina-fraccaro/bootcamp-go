package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GreetingRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	http.HandleFunc("/greetings", greetingsHandler)

	fmt.Println("Servidor escuchando en http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var request GreetingRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
		return
	}

	greeting := fmt.Sprintf("Hello %s %s", request.FirstName, request.LastName)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, greeting)

}
