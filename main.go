package main

import (
	"fmt"
	"strings"
)

func main(){
		confrenceName := "Go Confrence"
		const confrenceTickets int = 50
		var remainingTickets uint = 50
		bookings := []string{}

		greetUsers(confrenceName, confrenceTickets, remainingTickets)

		fmt.Printf("confrenceTickets is %T, remainingTickets is %T, confrenceName is %T\n ", confrenceTickets, remainingTickets, confrenceName)

	for {
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
					
					isValidName := len(firstName) >= 2 && len(lastName) >= 2
					isValidEmail := strings.Contains(email, "@")
					isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

					if isValidName && isValidEmail && isValidTicketNumber {
					remainingTickets = remainingTickets - userTickets
					bookings = append(bookings, firstName + " " + lastName)

					fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
					fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)
					
					printFirstNames(bookings)
					
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

func greetUsers(confName string, confTickets int, remTickets uint){
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available for you.\n", confTickets, remTickets)
	fmt.Println("Get Your tickets here to attend")
}
func printFirstNames(bookings []string) {
					firstNames := []string{}
					for _, booking := range bookings {
						var names = strings.Fields(booking)
						firstNames = append(firstNames, names[0])
					}
					fmt.Printf("These are the first names bookings %v \n", firstNames)
}
