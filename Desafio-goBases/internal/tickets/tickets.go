package tickets

import "time"

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

// ejemplo 3
//func AverageDestination(destination string, ticketList []Ticket) int {
//	count := GetTotalTickets(destination, ticketList)
//}
