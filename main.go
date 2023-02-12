package main

import "fmt"

// main function
func main() {

	n := 123234

	isPrime(n)

	_, msg := isPrime(n)
	fmt.Println(msg)

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
