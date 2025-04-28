package main

import (
	"fmt"
	"github.com/google/uuid"
)

var totalTickets = 100

func main() {

	for totalTickets > 0 {
		var ticket = 0
		fmt.Println("Welcome to the ticketBooking system!")
		fmt.Println("Available tickets: ", totalTickets)

		fmt.Println("Please enter ano of tickets: ")
		fmt.Scan(&ticket)

		if totalTickets > ticket && totalTickets > 0 && ticket > 0 {
			totalTickets = totalTickets - ticket
		} else {
			fmt.Println("Please enter valid number of ticket !!!!")
			continue
		}

		uuid := uuid.New()
		fmt.Println("Here is the code: ", uuid)

		fmt.Println("Currently Available tickets: ", totalTickets)
	}
}
