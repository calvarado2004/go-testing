package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main function
func main() {

	//print a welcome message
	intro()

	//crete a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(doneChan)

	// block until the doneChan gets a value
	<-doneChan

	// close the channel
	close(doneChan)

	// print a goodbye message
	fmt.Println("Goodbye")

}

// readUserInput reads user input and runs the program
func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

// checkNumbers checks if the user input is a valid number
func checkNumbers(scanner *bufio.Scanner) (string, bool) {

	// read user input
	scanner.Scan()

	// check if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// convert the user input to an integer
	n, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "Please enter a valid integer number", false
	}

	// check if the number is prime
	_, msg := isPrime(n)

	return msg, false
}

// intro prints a welcome message
func intro() {
	fmt.Println("Welcome to the prime number checker")
	fmt.Println("Enter a number to check if it is prime")
	fmt.Println("Enter 'q' to quit")
	prompt()
}

// prompt prints a prompt to the user
func prompt() {
	fmt.Print("Enter a number: ")
}

// isPrime returns a boolean value and a string
func isPrime(n int) (bool, string) {
	// do not include 0 and 1 by definition

	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by definition", n)
	}

	if n < 0 {
		return false, fmt.Sprintf("%d negatives are not prime", n)
	}

	// use the modules operator repetitively to check if the number is prime
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime because %d is a factor of it", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)

}
