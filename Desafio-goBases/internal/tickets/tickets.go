package tickets

import (
	"time"
)

type Ticket struct {
	ID         int
	Name       string
	Email      string
	Destiny    string
	FlightTime time.Time
	Price      int
}

func GetTotalTickets(destination string, ticketList []Ticket) int {
	count := 0
	for _, ticket := range ticketList {
		if ticket.Destiny == destination {
			count++
		}
	}
	return count
}

func CountPeopleByTimeRange(ticketList []Ticket, startHour, endHour int) int {
	count := 0
	for _, ticket := range ticketList {
		flightHour := ticket.FlightTime.Hour()
		if flightHour >= startHour && flightHour <= endHour {
			count++
		}
	}
	return count
}

func AverageDestination(destination string, ticketList []Ticket) float64 {
	count := GetTotalTickets(destination, ticketList)
	var average = float64(count) / float64(len(ticketList)) * 100
	return average
}
