# snickerboa 

(or for short boa)

Golang package with functions for any use.

Just a beginner doing his best to create a open source package for Go

## Usage

Use function like this

```Go
package main

import (
	"fmt"

	boa "github.com/davidn5013/snickerboa"
)

func main() {
	boa.OneArgChkOrExit("Missing argument")
	n := boa.Arg1ToIntOrExit()
	fmt.Println(boa.SumofPrimesUntil(n))
}
```

```Text
> go mod tidy
go: finding module for package github.com/davidn5013/snickerboa
go: downloading github.com/davidn5013/snickerboa...
```

## FunctionsProgress

### Math
- AddBigStr([]string) string Add big number in strings
- CheckPalindrome(string) check of a string is a palindrome use strconv.Itoa to checknumbers [project euler 4](https://projecteuler.net/problem=4)
- Factorial() (uint64) factorial using uint64 with overflow check, Panic on overflow
  Note Rename to MostFactorial
- FactorialBig() (\*big.Int) factorial using bignumber
- PythTripProd(int) A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, [project euler 9](https://projecteuler.net/problem=9)
- SumofPrimesUntil(int) my solution to [project euler 10](https://projecteuler.net/problem=10)

### Tools
- Arg1ToIntOrExit() (int) converts args[1] to int
- ClearScreen() Clear the screen using ANSI code not os
- FileToFields() ([]string) Read a whole to one slice of strings using space and newline as delimit
- FileToStr() (string) Read a whole file to one string
- Make2D() ([][]T) create a [][]matrix of any type
- OneArgChkOrExit(string) check for one argument if not exit
- PressEnter() Show press return and waits for return
- PrintMemUsage() (int) Report memory being used in MB and garage collection cycles completed
- Quiet() (func()) revert all output from program for testing av main, return a function revert output back
- SetTimer() (func time.Duration) Time exaction
- SortString() (string) sort character in string

### Types

- stack of []interface{} methods: NewStack, Push, Pop, Peek, IsEmpty, Size

```Go
/*
 * Usage Exampel
 */

func main() {
	s := snickerboa.NewStack()
	fmt.Println(s.IsEmpty()) // true
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Size()) // 3
	item, err := s.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(item) // 3
	}
	item, err = s.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(item) // 3
	}
	fmt.Println(s.Size())    // 2
    fmt.Println(s.IsEmpty()) // false
}
```

## Contribut

- If you want to contribut it fine, beginner friendly
- Spellchecking is very much appreciated

And again bare with me I'am a beginner 
And conduct yourself as a proper person and all is good

/ David
