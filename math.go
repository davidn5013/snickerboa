// Golang package with func for any use
// math part
package snickerboa

import (
	"github.com/fxtlabs/primes"
)

// sumofPrimesUntil sums all primenumber upp to n
// Example 10 = 17
func SumofPrimesUntil(n int) (sumPri uint64) {
	p := primes.Sieve(n)
	for i := 0; i < len(p); i++ {
		sumPri = sumPri + uint64(p[i])
	}
	// fmt.Printf("%v\n", p)
	return sumPri
}

// checkPalindrome check if a string can be read back to front
// thanks for the steal (https://www.eventslooped.com/posts/interview-question-in-go-2/)
func CheckPalindrome(testString string) bool {
	isPalindrome := true
	sLength := len(testString)

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

// pythTripProd A Pythagorean triplet is a set of three natural numbers
// For example, 32 + 42 = 9 + 16 = 25 = 52
// https://projecteuler.net/problem=9
func pythTripProd(n int) (a, b, c, prod int) {
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
