package reader

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/storage"
)

func ReadTickets(path string, tickets *storage.TicketsSlice) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	rd := csv.NewReader(f)

	for {
		record, err := rd.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}

		price, err := strconv.Atoi(record[6])
		if err != nil {
			return err
		}

		tickets.AddTicket(id, record[1], record[2], record[3], record[4], record[5], price)

	}

	return nil
}
