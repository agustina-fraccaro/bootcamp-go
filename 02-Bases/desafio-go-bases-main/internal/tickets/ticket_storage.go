package tickets

type StorageTicket interface {
	AddTicket(id int, name string, email string, destination string, country string, flightTime string, price int)
	GetTotalTickets(destination string) (int, error)
	AverageDestination(destination string) (int, error)
}
