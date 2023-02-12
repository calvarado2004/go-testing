package main

import "testing"

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
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotMessage := isPrime(tt.testNumber)
			if gotResult != tt.expectedResult {
				t.Errorf("isPrime() got = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotMessage != tt.expectedMsg {
				t.Errorf("isPrime() got1 = %v, want %v", gotMessage, tt.expectedMsg)
			}
		})
	}

}
