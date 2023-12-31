package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = bookTickets(userTickets, firstName, lastName, email)
			
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entere doesn't contain an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
		}
	}

	wg.Wait()
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) uint {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("These are all our bookings: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	return remainingTickets
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("How many tickets you need")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("##########")
	wg.Done()
}
