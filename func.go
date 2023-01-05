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
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fxtlabs/primes"
)

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
func CheckPalindrome(s string) bool {
	s = strings.ToLower(s)
	revS := ReverseStr(s)
	return strings.Compare(s, revS) == 0
}

// ReverseStr Turn "David" "divaD"
func ReverseStr(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
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
		defer func() {
			err := null.Close()
			if err != nil {
				fmt.Printf("Can't close file %s\n", err)
				return
			}
		}()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

// Strip removes all non alfa number from string.
// Removing everthing but a-z  A-Z  0-9
func Strip(s string) string {
	var result strings.Builder
	result.Grow(len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}

// StripSwe removes all non alfa number in swedish from string.
// Removing everthing but a-ö  A-Ö  0-9
func StripSwe(s string) string {
	var result strings.Builder
	result.Grow(len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'ö') ||
			('A' <= b && b <= 'Ö') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
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

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	bToMb := func(b uint64) uint64 {
		return b / 1024 / 1024
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

// PressEnter Wait for user to press return
func PressEnter() {
	fmt.Println("Press return")
	_, err := fmt.Scanln()
	if err != nil {
		fmt.Printf("Error reading input %s\n", err)
		return
	}
}

const clsAnsicode = "\x1bc"

// ClearScreen cls - Clears Terminal Screen using Ansi code Hex 1bc (444) 1b
func ClearScreen() {
	fmt.Print(clsAnsicode)
}

//
//		a r r a y  f u n c t i o n s
//

// ReverseSlice of any time
func ReverseSlice[T any](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// Make2D return matrix [][] of any type of size n,m using generics
func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

// Map runs a function on every item of any type in array
// Example multi with 2 and convert to string:
// s := []int{2, 4, 8, 11}
// ds := Map(s, func(i int) string {return strconv.Itoa(2*i)})
// Example Create a ny array of uppcase names:
// names := []string{"joe", "mike", "sue"}
// namesUpper := Map(names, strings.ToUpper)
func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Filter returns part of array that complies to a func filter
// Example return a list of even number:
// evens := Filter(s, func(i int) bool {return i % 2 == 0})
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reduce run a math function on array return the result
// Example:
// s := []int{1, 2, 3, 4, 5}
// product := Reduce(s, 1, func(a, b int) int {return a*b})
// sum := Reduce(s, 1, func(a, b int) int {return a+b})
// name := []string{"One","Two","Three"}
// longstring := Reduce(s, 1, func(a, b int) int {return a+b})
func Reduce[T, U any](s []T, init U, f func(U, T) U) U {
	r := init
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Find seek value in a array of any type using a function
// Example find name Chris
//
//	people := []Person{
//		Person{Name: "Kent Beck"},
//		Person{Name: "Martin Fowler"},
//		Person{Name: "Chris James"},
//	}
//	king, found := Find(people, func(p Person) bool {
//		return strings.Contains(p.Name, "Chris")
//	})
func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}

// If Go replacement for ternary using generic,
// Example:
// gender := "male"
// sex := If(gender == "male", "boy", "girl)
// expectancy .= If(gender == "male",74,80)
// Example2:
// samestring := If(string.Compare(a,b)==0,a,b)
func If[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}
