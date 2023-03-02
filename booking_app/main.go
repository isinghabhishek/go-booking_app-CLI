// 1). go mod init <project_name>
// 2). package main
// 3). import "fmt"

package main

import (
	"booking_app/helper"
	"fmt"
	"sync"
	"time"
)

// intialising a variable
var conferencesName = "Go Conferences"
var conferencesTickets int = 50
var remainingTickets uint = 50

// Arrays
// var variable_name [size]variable_type
// Array Declaration => var bookings [50]string
// slice
// var bookings = []string{}

// list of maps
// var bookings = make([]map[string]string, 0)

// list of struct
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// waitgroup
var wg = sync.WaitGroup{}

func main() {

	// information can be passed into function as parameters
	// parameters are also called arguments
	greetUser()

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferencesTickets, conferencesName, remainingTickets)

	// looping
	// for {
	// calling getUserInput Function
	firstName, lastName, email, userTickets := getUserInput()

	// validateUserInput function called
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// isValidCity := city == "Noida" || city == "Bangaluru"

	if isValidName && isValidEmail && isValidTicketNumber {

		// calling bookTicket Function
		bookTickets(userTickets, firstName, lastName, email)

		// calling the sendTicket function
		// "go" keywoard --> a gorouting is a lightweight thred managed by the Go runtime

		//Add: Sets the number of gorouting to wait for
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		// call function printFirstName
		firstNames := getFirstName()
		fmt.Printf("The first names of booking are: %v\n", firstNames)

		// if statement
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conferences is booked out, Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First Name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Printf("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Printf("Number of Tickets you entered is inValid")
		}
	}
	// Wait: Blocks until the WaitGroup counter is 0.
	wg.Wait()
	//}
}

// Switch statement
// city := "Delhi"

// switch city {
// case "Gurugram":
// 	// execute code for booking Gurugram conference tickets
// case "Noida":
// 	// execute code for booking Noida conference tickets
// case "Mumbai":
// 	// execute code for booking Delhi conference tickets
// case "Bengaluru", "Hyderabad":
// 	// execute code for booking Bengaluru, Hyderabad conference tickets
// default:
// 	fmt.Printf("No valid city selected")
// }

// same parameter name or other what we like must be relevent
func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferencesName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferencesTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// parameter that use in funtion[i/p parameters   o/p parameters]
func getFirstName() []string {
	firstNames := []string{}
	// for each loop
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	// Scan() function is used to take input from users

	// &(pointer) => A pointer is a variable that points to the memory address of another variable.
	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)
	// Scan function can now assign the user's value to the userName variable
	// Because it has a pointer to its memory addresss
	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email id")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	// as tickets are booking we need to decrease remainingTickets count
	remainingTickets = remainingTickets - uint(userTickets)

	// var mySlice []string
	// var mymap map[string]string

	// create a map for a user
	// var userData = make(map[string]string)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// don't these all after "struct"

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings[0] = firstName + " " + lastName
	// slice of sting
	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The First value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferencesName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
	// Done: Decrements the waitgroup counter by 1 So this is called by the goroutine to indicate that it's finished
	wg.Done()
}
