package main

import (
	"app/internal/application"
	"fmt"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")
	path := os.Getenv("FILE_PATH")
	app := application.NewDefaultHTTP(":8080", token, path)

	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
