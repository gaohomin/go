package main

import (
	"apis/src/hms/practice/4_package/greeting"
	"apis/src/hms/practice/4_package/keyboard"
	"fmt"
	"log"
)

func main() {
	// call package method
	greeting.Hello()
	greeting.Hi()

	/* call method - pass_fail*/
	fmt.Print("Enter a grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
	/* call method, fahrenheit to celsius*/
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degrees Celsius \n", celsius)
}
