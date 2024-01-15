package loader

import (
	"app/internal"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (tickets map[int]internal.TicketAttributes, err error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	tickets = make(map[int]internal.TicketAttributes)
	for {
		record, er := r.Read()
		if er != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			return
		}

		// serialize the record
		id, er := strconv.Atoi(record[0])
		if er != nil {
			err = fmt.Errorf("error converting id to int: %v", err)
			return
		}
		ticket := internal.TicketAttributes{
			Name:    record[1].(string),
			Email:   record[2].(string),
			Country: record[3].(string),
			Hour:    record[4].(string),
			Price:   record[5].(int),
		}

		// add the ticket to the map
		tickets[id] = ticket
	}

	return
}
