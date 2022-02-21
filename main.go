package main

import	"fmt"
		const confrenceTickets int = 50
		var confrenceName = "Go Confrence" 
		var remainingTickets uint = 50
		var bookings = make([]UserData, 0)

		type UserData struct {
			firstName string
			lastName string
			email string
			numberOfTickets uint 

		}

func main(){

		greetUsers()

		fmt.Printf("confrenceTickets is %T, remainingTickets is %T, confrenceName is %T\n ", confrenceTickets, remainingTickets, confrenceName)

	for {
					firstName, lastName, email, userTickets := getUserInput()
					isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

					if isValidName && isValidEmail && isValidTicketNumber {
 
					bookTicket(userTickets,firstName, lastName, email)
					
					firstNames := getFirstNames()
					fmt.Printf("These are the first names bookings %v \n", firstNames)

					if remainingTickets == 0 {
						// end program
						fmt.Println("Our confrence is booked out. Come back next year.")
						break
					}
					} else {
						if !isValidName{
							fmt.Println("first name or last name is not correct")
						}
						if !isValidEmail{
							fmt.Println("email address you entred doesnot contain @ sign")
						}
						if !isValidTicketNumber{
							fmt.Println("Number of tickets is invalid")
						}
					}
				}
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available for you.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get Your tickets here to attend")
}
func getFirstNames() []string {
					firstNames := []string{}
					for _, booking := range bookings {
						firstNames = append(firstNames, booking.firstName)
					}
					return firstNames
}

func getUserInput()(string, string, string, uint){
						var firstName string
					var lastName string
					var email string

					var userTickets uint
					// ask user for their name
					fmt.Println("Enter your first name:")
					fmt.Scan(&firstName)

					fmt.Println("Enter your last name:")
					fmt.Scan(&lastName)

					fmt.Println("Enter your email:")
					fmt.Scan(&email)

					fmt.Println("Enter your number of Tickets:")
					fmt.Scan(&userTickets)
					return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string ){
					remainingTickets = remainingTickets - userTickets

					// create a map for a user //Key value names ("firstName","lastName"... ) could be any names
					var userData = UserData {
									firstName: firstName,
									lastName: lastName,
									email: email,
									numberOfTickets: userTickets,
					}
					bookings = append(bookings, userData)
					fmt.Printf("List of bookings %v\n", bookings)
					fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
					fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)
}