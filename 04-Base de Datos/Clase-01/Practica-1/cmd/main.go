package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// application
	// - config
	cfg := &application.ConfigDefault{
		Database: mysql.Config{
			User:      "user1",
			Passwd:    "secret_password",
			Net:       "tcp",
			Addr:      "127.0.0.1:3306",
			DBName:    "my_db",
			ParseTime: true,
		},
		Address: "127.0.0.1:8081",
	}
	app := application.NewApplicationDefault(cfg)
	// - run
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
}
