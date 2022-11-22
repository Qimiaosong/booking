package main

import (
	"fmt"
	"strings"
)

func main(){
	const conferenceTickets int = 50 
	var remainingTickets int = 50 
	conferenceName := "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGrt your tickets here to attend\n",conferenceName,conferenceTickets,remainingTickets)

	for{
		var firstName string 
		var lastName string 
		var email string
		var userTickets int

		fmt.Println("Enter your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		//validate user input 
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber{
			remainingTickets = remainingTickets - userTickets 
			bookings = append(bookings,firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
			fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

			firstNames := []string{}
			for _,booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names %v\n",firstNames)
			if remainingTickets == 0{
				fmt.Println("Our conference is booked out.Come back next year")
				break	
			}
		}else{
			if !isValidName{
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
	city := "London"
	switch city {
	case "New York":
		fmt.Println("Welcome to NewYork")
	case "Singapore","HongKong":
		fmt.Println("Welcome to HongKong")
	case "London","Berlin":
		fmt.Println("Welcome to LB")
	case "Mexico City":
		fmt.Println("Welcome to Mexico City")
	default:
		fmt.Println("No valid city selected")
	}
}