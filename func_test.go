// Package snickerboa Test
package snickerboa

import (
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	if CheckPalindrome("9009") != true {
		t.Errorf("got false on testin 9009")
	}
	// for prober test you need to comment out code
	if CheckPalindrome("12345") != false {
		t.Errorf("failed test udd numer")
	}
}

// TODO Test for Arg1ToIntOrExit
// () int
// This did not work
// func TestArg1ToIntOrExit(t *testing.T) {
// 	os.Args = append(os.Args, "100")
// 	v := Arg1ToIntOrExit()
// 	if v != 100 {
// 		t.Errorf("got %d wanted 100", v)
// 	}
// }

// TODO Test OneArgChkOrExit()
// Test this how it runs os.Exit(1) in fail...
// args (s string)
// func TestOneArgChkOrExit(t *testing.T) {}

func TestPythTripProd(t *testing.T) {
	_, _, _, prod := PythTripProd(24)
	if prod != 480 {
		t.Errorf("got false on test 100")
	}
}

func TestSumofPrimesUntil(t *testing.T) {
	v := SumofPrimesUntil(10)
	if v != 17 {
		t.Errorf("got false wanted 17 gut %d", v)
	}
}

// TODO Make test for
// func FileToStr(f string) (s string) {
// func FileToFields(f string) []string {

func TestAddBigStr(t *testing.T) {
	str, _ := AddBigStr([]string{
		"19999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999",
		"29999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"})
	if str !=
		"49999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999998" {
		t.Errorf("got false wanted 4 witdh 99 9:s and 8 after gut %s", str)
	}
}
