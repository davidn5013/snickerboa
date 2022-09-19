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
