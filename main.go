package main

import (
	"fmt"
	weather "sk8app/APIs"
	dbb "sk8app/DB"
	"strings"
)

func main() {
	answer := ""
	fmt.Println("Do you wanna signup, login, or delete your account? (s, l, d)")
	fmt.Scanf("%s", &answer)
	begin(answer)

	//Figure out how to stop code from moving on if login, delete, or signup fails

	choice := ""

	fmt.Println("\nWelcome to your Sk8 journal, Want me to recommend you a trick a park or tell you the weather? (Trick, Park, Weather):")
	fmt.Scanf("%s", &choice)
	fmt.Println("\n")
	parkortrick(choice)
}

func parkortrick(s string) {
	if strings.ToLower(s) == "trick" {
		dbb.GetTricks()

	} else if strings.ToLower(s) == "park" {
		dbb.GetParks()
	} else if strings.ToLower(s) == "weather" {
		weather.Main()
	} else {
		fmt.Println("\nYo dude, looks like your typing needs some work, make sure you type the specified options correctly.")
	}
}

func begin(answer string) {
	fName := ""
	lName := ""
	Email := ""
	uName := ""
	pWord := ""
	if strings.ToLower(answer) == "s" {
		fmt.Println("Name: ")
		fmt.Scanf("%s", &fName)
		fmt.Println("Last Name: ")
		fmt.Scanf("%s", &lName)
		fmt.Println("Email: ")
		fmt.Scanf("%s", &Email)
		fmt.Println("Username: ")
		fmt.Scanf("%s", &uName)
		fmt.Println("Password: ")
		fmt.Scanf("%s", &pWord)

		//db.CreateUser(fName, lName, Email, uName) will be replaced with grpc call
	} else if strings.ToLower(answer) == "l" {
		fmt.Println("Enter username & password to login: ")
		fmt.Println("Username: ")
		fmt.Scanf("%s", &uName)
		fmt.Println("Password: ")
		fmt.Scanf("%s", &pWord)
		//db.SignIn(uName) will be replaced with grpc call
	} else if strings.ToLower(answer) == "d" {
		fmt.Println("Which user do you want to delete? (Type Username & Password for the User you want deleted)")
		fmt.Println("Username: ")
		fmt.Scanf("%s", &uName)
		fmt.Println("Password: ")
		fmt.Scanf("%s", &pWord)
		//db.DeleteUser(uName) will be replaced with grpc call
	} else {
		fmt.Println("Make sure you typed s, l, or d", "\n")
	}
}
