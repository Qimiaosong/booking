package main

import (
	"fmt"
	"sync"
	"time"
)


const conferenceTickets int = 50 

var remainingTickets int = 50 
var conferenceName = "Go conference"
var bookings = make([]UserData,0)

type UserData struct{
	firstName		string 
	lastName		string 
	email			string 
	numberOfTickets	int
}

var wg = sync.WaitGroup{}

func main(){

	greetUsers()


	firstName,lastName,email,userTickets := getUserInput()
	isValidName,isValidEmail,isValidTicketNumber := validateUserInput(firstName,lastName,email,userTickets)

	if isValidName && isValidEmail && isValidTicketNumber{
		bookTicket(userTickets,firstName,lastName,email)

		wg.Add(1)
		go sendTicket(userTickets,firstName,lastName,email)

		firstNames := printFirstNames()
		fmt.Printf("The first names of bookings are:%v\n",firstNames)

		if remainingTickets == 0{
			fmt.Println("Our conferences is booked out.Come back next year.")
		}
	} else {
		if !isValidName{
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail{
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber{
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.wait()
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still avaiable.\nGet your tickets here to attend\n",conferenceName,conferenceTickets,remainingTickets)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _,booking := range bookings{
		firstNames = append(firstNames,booking.firstName)
	}
	return firstNames
}

func getUserInput()(string,string,string,int){
	var firstName string
	var lastName string 
	var email string 
	var userTickets int 

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter Number of Tickets: ")
	fmt.Scanln(&userTickets)

	return firstName,lastName,email,userTickets
}

func bookTicket(userTickets int,firstName string,lastName string,email string){
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:	     firstName,
		lastName:		 lastName,
		email:			 email,
		numberOfTickets: userTickets
	}

	bookings = append(bookings,userData)
	fmt.Printf("List of bookings is %v\n",bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

}

func sendTicket(userTickets int,firstName string,lastName string,email string){
	time.Sleep(50*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n",ticket,email)
	fmt.Println("#########################")
	wg.Done()
}
