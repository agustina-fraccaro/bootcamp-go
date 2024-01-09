package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/reader"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/storage"
)

func main() {
	ticketSlice := storage.NewTicketSlice()
	err := reader.ReadTickets("tickets.csv", ticketSlice)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalTickets, err := ticketSlice.GetTotalTickets("China")

	if err != nil {
		return
	}

	fmt.Println(totalTickets)

}
