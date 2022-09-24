// Package snickerboa is a for func for any use
package snickerboa

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/fxtlabs/primes"
)

//
//		I N F O
//

// LastEdit show last time snickerboa was change
// Manual change for now
func LastEdit() {
	fmt.Println("This hello was last edit 20220924 19:55:20")
}

// Hello print "Hello from snickerboa"
func Hello() {
	fmt.Println("Hello from snickerboa")
}

//
//		M A T H
//

// SumofPrimesUntil sums all primenumber upp to n
// Example 10 = 17
func SumofPrimesUntil(n int) (sumPri uint64) {
	p := primes.Sieve(n)
	for i := 0; i < len(p); i++ {
		sumPri = sumPri + uint64(p[i])
	}
	// fmt.Printf("%v\n", p)
	return sumPri
}

// CheckPalindrome check if a string can be read back to front
// thanks for the steal (https://www.eventslooped.com/posts/interview-question-in-go-2/)
func CheckPalindrome(testString string) bool {
	isPalindrome := true
	sLength := len(testString)

	// string of odd length can't be palindrome
	if sLength%2 != 0 {
		return false
	}

	for i := 0; i < sLength/2; i++ {
		currCharFwd := testString[i]           // character to compare going forward
		currCharBwd := testString[sLength-1-i] // character to compare going from the end

		if currCharFwd != currCharBwd {
			isPalindrome = false // we found a non-match, it's not a palindrome
			break                // no need to compare further
		}
	}

	return isPalindrome
}

// PythTripProd A Pythagorean triplet is a set of three natural numbers
// For example, 32 + 42 = 9 + 16 = 25 = 52
// https://projecteuler.net/problem=9
func PythTripProd(n int) (a, b, c, prod int) {
	for a = 1; a < (n / 3); a++ {
		for b = 1; b < (n / 2); b++ {
			c = n - b - a
			as, bs, cs := a*a, b*b, c*c
			if (as + bs) == cs {
				return a, b, c, a * b * c
			}
		}
	}
	return 0, 0, 0, 0
}

// AddBigStr take a string array of of big number sum them upp
// returns a string
func AddBigStr(arrStrNumber []string) (string, error) {
	x := new(big.Int)
	s := new(big.Int)

	for _, c := range arrStrNumber {
		if len(c) > 0 {
			_, err := fmt.Sscan(c, x)
			if err != nil {
				return "", err
			}
			s.Add(x, s)
		}
	}
	return s.String(), nil
}

//
//		s u p p o r t
//

// OneArgChkOrExit check for 1 argument and
// exit print s and exit if none
func OneArgChkOrExit(s string) {
	if len(os.Args) < 2 {
		fmt.Println(s)
		os.Exit(1)
	}
}

// Arg1ToIntOrExit convert args[1] to int
// exit with prints args[1] is not a number on fail
func Arg1ToIntOrExit() int {
	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s is not a number", os.Args[1])
		os.Exit(1)
	}
	return v
}

// FileToStr Read a whole file to one string
func FileToStr(f string) (s string) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(b)
}

// FileToFields Read a whole file to one slice of string
// using space and newline (strings.Fields delimit)
func FileToFields(f string) []string {
	return strings.Fields(FileToStr(f))
}
