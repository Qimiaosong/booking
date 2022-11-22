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

		remainingTickets = remainingTickets - userTickets 
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
		fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

		firstNames := []string{}
		for _,booking := range bookings{
			//将字符串转换为slice切片
			var names = strings.Fields(booking)
			fmt.Printf("llllllllllllllllll%v",names)
			firstNames = append(firstNames,names[0])
		}
		fmt.Printf("The first names %v\n", firstNames)

	}
}