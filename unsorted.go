// Golang package with func for any use
// unsorted functions or datastructs part
package snickerboa

import (
	"fmt"
	"os"
	"strconv"
)

// oneArgChkOrExit check for 1 argument and
// exit print s and exit if none
func oneArgChkOrExit(s string) {
	if len(os.Args) < 2 {
		fmt.Println(s)
		os.Exit(1)
	}
}

// arg1ToIntOrExit convert args[1] to int
// exit with prints args[1] is not a number on fail
func arg1ToIntOrExit() int {
	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s is not a number", os.Args[1])
		os.Exit(1)
	}
	return v
}
