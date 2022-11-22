package main 

import(
	"fmt"
	"strings"
)

const conferenceTickets int = 50 

var remainingTickets int = 50 
var conferenceName = "Go conference"
var bookings = make([]User,0)

type User struct{
	firstName 		string 
	lastName  		string 
	email     		string 
	numberOfTickets int 
}

func main(){

	greetUsers()

	for{

		firstName,lastName,email,userTickets := getUserInput()
		isValidName,isValidEmail,isValidTicketNumber := validateUserInput(firstName,lastName,email,userTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
			bookTicket(userTickets,firstName,lastName,email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n",firstNames)

			if remainingTickets == 0{
				break
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
			continue
		}
	}
}

func printFirstNames() []string{
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

func validateUserInput(firstName string,lastName string,email string,userTickets int)(bool,bool,bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets 
	return isValidName,isValidEmail,isValidTicketNumber
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still avaiable.\nGet your tickets here to attend\n",conferenceName,conferenceTickets,remainingTickets)
}

func bookTicket(userTickets int,firstName string,lastName string,email string){
	remainingTickets = remainingTickets - userTickets

	//create user map 
	var user = User{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings,user)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

}
