package main

import (
	"fmt"
	"strings"
)

var (
	answersTocheck = []string{
		"n",
		"no",
	}
	num1      int    = 0
	num2      int    = 0
	value     string = " "
	valueStop bool   = true
	answer    string = ""
)

func checkInArray(array []string, strToCheck string) bool { //This func checking for no answers
	for _, i := range array {
		if strToCheck == i {
			return true
		}
	}
	return false
}

func higherNumSwitcher(num1 int, num2 int) (int, int) { // For corectly deviding we need to find higest num
	if num1 < num2 {
		answerNum1 := num2
		return answerNum1, num1
	}
	return num1, num2
}

func checkInputOperation() string { // Here we are checking inputs of operatinon
	var flagStop = true
	var value string
	fmt.Println("Enter the operation ('+ , - , /, *'): ")

	for flagStop == true {
		fmt.Scan(&value)
		if value == "+" {
			return "+"
		} else if value == "-" {
			return "-"
		} else if value == "/" {
			return "/"
		} else if value == "*" {
			return "*"
		} else {
			fmt.Println("Pleas enter only this operation ('+ , - , /, *'): ")
		}
	}
	return ""
}

func calculate(num1 int, num2 int, valueToCalc string) int { // Func where we calculate inputed items
	num1, num2 = higherNumSwitcher(num1, num2)
	if valueToCalc == "*" {
		return num1 * num2
	} else if valueToCalc == "+" {
		return num1 + num2
	} else if valueToCalc == "-" {
		return num1 - num2
	} else {
		return num1 / num2
	}
}

func main() { // Entry point
	fmt.Println("Hello, this is calc")
	for valueStop == true {
		fmt.Println("Enter first number: ")
		fmt.Scan(&num1)
		value = checkInputOperation()
		fmt.Println("Enter second number: ")
		fmt.Scan(&num2)
		calcAnswer := calculate(num1, num2, value)
		fmt.Printf("The answer is: %v\n", calcAnswer)
		fmt.Print("Do you wanna continue? Y/N \n")
		fmt.Scan(&answer)
		if checkInArray(answersTocheck, strings.ToLower(answer)) {
			valueStop = false
		}
	}
}
