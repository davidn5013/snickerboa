// Package snickerboa is toolsbox for anything in Go - proj euler and more
//
// Created as an example of help module. Have some math function and some
// function for simple argument and file handling.
//
// https://github.com/davidn5013/snickerboa
package snickerboa

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fxtlabs/primes"
)

//
//		I N F O
//

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
// thanks for the steal
//
// https://www.eventslooped.com/posts/interview-question-in-go-2/
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
//
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

// AddBigStr take an array of strings with numbers and
// sum then up then returns a string or error if the
// strings can't be converted
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
// exit program if none or return args as string
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
	b, err := os.ReadFile(f)
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

// SetTimer Set a timmer and return a func
// that returns time.Duration from the timer
//
// func main() {
// stopTimer := SetTimer()
// ...
// fmt.Printf("Elapsed time:%v\n".stopTimer())
func SetTimer() func() time.Duration {
	t := time.Now()
	return func() time.Duration {
		return time.Since(t)
	}
}

// Quiet is func for test function and supressing output
// In test function set first line to
//
// defer quiet()()
func Quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

// FactorialBig factorial using bignumber
func FactorialBig(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, FactorialBig(n.Sub(x, n)))
}

// Factorial with overflow check panics
func Factorial(number int) uint64 {
	var result uint64 = 1
	if number < 0 {
	} else {
		for i := 1; i <= number; i++ {
			result = uint64(mul64p(int64(result), int64(i)))
		}
	}
	return result
}

// mul64p is the unchecked panicing version of Mul64
// used by Factorial()
func mul64p(a, b int64) int64 {
	r, ok := mul64(a, b)
	if !ok {
		panic("multiplication overflow")
	}
	return r
}

// mul64 performs * operation on two int64 operands
// returning a result and status
// Used by mul64p()
func mul64(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	c := a * b
	if (c < 0) == ((a < 0) != (b < 0)) {
		if c/b == a {
			return c, true
		}
	}
	return c, false
}

// SortString input "ÖÄÅöäåabcABC" output "ABCabcÄÅÖäåö"
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Make2D return matrix [][]anytime of size n,m using generics
func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}
