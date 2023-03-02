// 1). go mod init <project_name>
// 2). package main
// 3). import "fmt"

package main

import "fmt"

func main() {
	// intialising a variable
	conferencesName := "Go Conferences"
	const conferencesTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferencesTickets, conferencesName, remainingTickets)

	fmt.Printf("Welcome to our %v booking application\n", conferencesName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferencesTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Arrays
	// var variable_name [size]variable_type
	// Array Declaration => var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets int
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

	// as tickets are booking we need to decrease remainingTickets count
	remainingTickets = remainingTickets - uint(userTickets)
	bookings[0] = firstName + " " + lastName

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v/n", remainingTickets, conferencesName)
}
