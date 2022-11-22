package main 
import "fmt"

func main(){

	const conferenceTickets int = 50 
	var remainingTickets int = 50 
	conferenceName := "Go conference"

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v of tickets and %v are still available.\nGet your tickets here to attend\n",conferenceName,conferenceTickets,remainingTickets)

	var firstName string 
	var lastName string 
	var email string 
	var userTickets int 

	//asking for user input 
	fmt.Println("Enter your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scanln(&userTickets)

	remainingTickets = conferenceTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets of remaining for %v\n", remainingTickets, conferenceName)
}
