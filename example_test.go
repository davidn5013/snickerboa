package snickerboa_test

import (
	"fmt"

	boa "github.com/davidn5013/snickerboa"
)

func ExampleFileToStr() {
	s := boa.FileToStr("test.txt")
	fmt.Println(s)
	// Output: TEST FILE
}

func ExampleFileToFields() {
	s := boa.FileToFields("test.txt")
	fmt.Println(s[0])
	// Output: TEST
}
