package tickets

import (
	"fmt"
	"testing"
	"time"
)

var ticketList = []Ticket{
	{
		ID:         1,
		Name:       "Cliente1",
		Email:      "cliente1@mercadolibre.com",
		Destiny:    "China",
		FlightTime: time.Date(2023, time.August, 15, 14, 30, 0, 0, time.UTC),
		Price:      100,
	},
	{
		ID:         2,
		Name:       "Cliente2",
		Email:      "cliente2@mercadolibre.com",
		Destiny:    "Brazil",
		FlightTime: time.Date(2023, time.August, 15, 14, 30, 0, 0, time.UTC),
		Price:      200,
	},
	{
		ID:         3,
		Name:       "Cliente3",
		Email:      "cliente3@mercadolibre.com",
		Destiny:    "Brazil",
		FlightTime: time.Date(2023, time.August, 15, 14, 30, 0, 0, time.UTC),
		Price:      5000,
	},
}

func TestGetTotalTickets1(t *testing.T) {

	// Caso de prueba 1: Verificar la cantidad de personas que viajan a "Brazil"
	t.Run("TestGetTotalTickes_Brazil", func(t *testing.T) {
		want := 2
		got := GetTotalTickets("Brazil", ticketList)
		if got != want {
			t.Errorf("GetTotalTicketas(Brazil) = %v, want %v", got, want)
		}
		fmt.Printf("Sucess GetTotalTicketas(Brazil) = %v, want %v", got, want)
	})

	// Caso de prueba 1: Verificar la cantidad de personas que viajan a "Mexico"
	t.Run("TestGetTotalTickes_Brazil", func(t *testing.T) {
		want := 0
		got := GetTotalTickets("Mexico", ticketList)
		if got != want {
			t.Errorf("GetTotalTicketas(Mexico) = %v, want %v", got, want)
		}
		fmt.Printf("Sucess GetTotalTicketas(Mexico) = %v, want %v", got, want)
	})
}

func TestCountPeopleByTimeRange(t *testing.T) {
	// Caso de prueba 1: Verificar la cantidad de personas que viajan a la tarde (13 - 19hs)
	t.Run("TestCountPeopleByTimeRange_Tarde", func(t *testing.T) {
		want := 3
		got := CountPeopleByTimeRange(ticketList, 13, 19)
		if got != want {
			t.Errorf("CountPeopleByTimeRange(Tarde) = %v, want %v", got, want)
		}
		fmt.Printf("Sucess CountPeopleByTimeRange(Tarde) = %v, want %v", got, want)
	})

	// Caso de prueba 2: Verificar la cantidad de personas que viajan a la tarde (13 - 19hs)
	t.Run("TestCountPeopleByTimeRange_Noche", func(t *testing.T) {
		want := 0
		got := CountPeopleByTimeRange(ticketList, 20, 23)
		if got != want {
			t.Errorf("CountPeopleByTimeRange(Noche) = %v, want %v", got, want)
		}
		fmt.Printf("Sucess CountPeopleByTimeRange(Noche) = %v, want %v", got, want)
	})
}
