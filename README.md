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

- Hello() print Hello from snickerboa for testing
- LastEdit() giv last edit date
- SumofPrimesUntil(int) my solution to [project euler 10](https://projecteuler.net/problem=10)
- CheckPalindrome(string) check of a string is a palindrome use strconv.Itoa to checknumbers [project euler 4](https://projecteuler.net/problem=4)
- PythTripProd(int) A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, [project euler 9](https://projecteuler.net/problem=9)
- AddBigStr([]string) string Add big number in strings
- OneArgChkOrExit(string) check for one argument if not exit
- Arg1ToIntOrExit() (int) converts args[1] to int

- ## Contribut

- If you want to contribut it fine, beginner friendly
- Spellchecking is very much appreciated

And again bare with me I'am a beginner 
And conduct yourself as a proper person and all is good

/ David
