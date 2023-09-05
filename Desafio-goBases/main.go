package main

import (
	"Bootcamp/Desafio-goBases/internal/tickets"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("tickets.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo CSV:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}
	var ticketList []tickets.Ticket

	for _, record := range records {
		if len(record) == 6 {
			ticket := tickets.Ticket{}
			ticket.ID = parseInt(record[0])
			ticket.Name = record[1]
			ticket.Email = record[2]
			ticket.Destiny = record[3]

			// Parsear el tiempo en formato "15:04"
			flightTime, err := time.Parse("15:04", record[4])
			if err != nil {
				fmt.Println("Error al parsear la hora de vuelo:", err)
				continue
			}
			ticket.FlightTime = flightTime
			ticket.Price = parseInt(record[5])

			ticketList = append(ticketList, ticket)
		}
	}

	// Calcular cuántas personas viajan a un destino específico
	destination := "China"
	total := tickets.GetTotalTickets(destination, ticketList)
	fmt.Printf("Número de personas viajando a %s: %d\n", destination, total)

	// Calcular la cantidad de personas que vuela en cada horario

	madrugada := tickets.CountPeopleByTimeRange(ticketList, 0, 6)
	mañana := tickets.CountPeopleByTimeRange(ticketList, 7, 12)
	tarde := tickets.CountPeopleByTimeRange(ticketList, 13, 19)
	noche := tickets.CountPeopleByTimeRange(ticketList, 20, 23)

	fmt.Printf("Personas en madrugada: %d\n", madrugada)
	fmt.Printf("Personas en mañana: %d\n", mañana)
	fmt.Printf("Personas en tarde: %d\n", tarde)
	fmt.Printf("Personas en noche: %d\n", noche)

	// Calcular el promedio de personas que viajan a un pais
	average := tickets.AverageDestination(destination, ticketList)
	fmt.Printf("Promedio de personas que viaja a %s: %.2f%%\n", destination, average)
}

func parseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error al convertir a entero:", err)
		return 0
	}
	return val
}
