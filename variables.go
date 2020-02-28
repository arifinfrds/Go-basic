// MARK: - Part 3: Variables
// Reference: https://golangbot.com/variables/

package main

import "fmt"

func main() {
	printMyBio()
	printOtherBio()
	printDate()
}

func printDate() {
	day := "Friday"
	month := "February"
	year := 2020

	fmt.Println(day, month, year)
}

func printOtherBio() {
	var firstName = "Go"
	var lastName = "Lang"
	var age = 2

	fmt.Println("Hello, my first name is", firstName, "and my last name is", lastName, ".", "I am", age, "years old.")
}

func printMyBio() {
	var (
		firstName = "Arifin"
		lastName  = "Firdaus"
		age       = 23
	)
	fmt.Println("Hello, my first name is", firstName, "and my last name is", lastName, ".", "I am", age, "years old.")
}

func printAge() {
	var myAge int = 23
	fmt.Println("Hello everyone, my age is", myAge)

	var andyAge int = 17
	fmt.Println("Andy's age is", andyAge)

	var sisterAge = 20
	fmt.Println("My sister's age is", sisterAge)

	var name = "Arifin"
	fmt.Println("Hello all. My name is", name, ".", "My age is", myAge)
}
