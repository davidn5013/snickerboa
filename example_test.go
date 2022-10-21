package snickerboa_test

import (
	"fmt"

	boa "github.com/davidn5013/snickerboa"
)

func ExampleFileToStr() {
	// Reading a test.txt containing
	// TEST FILE
	s := boa.FileToStr("test.txt")
	fmt.Println(s)
	// Output: TEST FILE
}

func ExampleFileToFields() {
	// Reading a test.txt containing
	// TEST FILE
	s := boa.FileToFields("test.txt")
	fmt.Println(s[0])
	// Output: TEST
}
