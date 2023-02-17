package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {

	//test tables
	primeTests := []struct {
		name           string
		testNumber     int
		expectedResult bool
		expectedMsg    string
	}{
		{"negative number", -1, false, "-1 negatives are not prime"},
		{"zero", 0, false, "0 is not prime by definition"},
		{"one", 1, false, "1 is not prime by definition"},
		{"two", 2, true, "2 is a prime number"},
		{"three", 3, true, "3 is a prime number"},
		{"four", 4, false, "4 is not prime because 2 is a factor of it"},
		{"five", 5, true, "5 is a prime number"},
		{"six", 6, false, "6 is not prime because 2 is a factor of it"},
		{"seven", 7, true, "7 is a prime number"},
	}

	// run the tests of the test table
	for _, tt := range primeTests {
		// run the test using testing Run
		t.Run(tt.name, func(t *testing.T) {
			// call the function to test
			gotResult, gotMessage := isPrime(tt.testNumber)
			// compare the results, expected result
			if gotResult != tt.expectedResult {
				t.Errorf("isPrime() got = %v, want %v", gotResult, tt.expectedResult)
			}
			// compare the results, expected message
			if gotMessage != tt.expectedMsg {
				t.Errorf("isPrime() got1 = %v, want %v", gotMessage, tt.expectedMsg)
			}
		})
	}

}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create  a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to the write pipe
	os.Stdout = w

	// call the function to test
	prompt()

	// close the write pipe
	_ = w.Close()

	// restore os.Stdout
	os.Stdout = oldOut

	// read the output from the read pipe
	out, _ := io.ReadAll(r)

	// compare the results
	if string(out) != "Enter a number: " {
		t.Errorf("prompt() got = %v, want %v", string(out), "Enter a number: ")
	}

}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create  a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to the write pipe
	os.Stdout = w

	// call the function to test
	intro()

	// close the write pipe
	_ = w.Close()

	// restore os.Stdout
	os.Stdout = oldOut

	// read the output from the read pipe
	out, _ := io.ReadAll(r)

	// compare the results
	if !strings.Contains(string(out), "Welcome to the prime number checker") {
		t.Errorf("intro() got = %v, want %v", string(out), "Welcome to the prime number checker")
	}
}
