package main

import (
	"fmt"
	weather "sk8app/APIs"
	dbb "sk8app/DB"
	"strings"
)

func parkortrick(s string) {
	if strings.ToLower(s) == "trick" {
		dbb.GetTricks()

	} else if strings.ToLower(s) == "park" {
		dbb.GetParks()
	} else {
		fmt.Println("\nYo dude, looks like your typing needs some work, make sure you type the specified options correctly.")
	}
}

func main() {
	choice := ""

	weather.Main()
	fmt.Println("\nWelcome to your Sk8 journal, Want me to recommend you a trick or a park? (Trick or Park):")
	fmt.Scanf("%s", &choice)
	fmt.Println("\n")
	parkortrick(choice)
}
