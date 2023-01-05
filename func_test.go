// Package snickerboa Test
package snickerboa

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	if CheckPalindrome("AmanaplanacanalPanama") != true {
		t.Errorf("got false on testin string A man..")
	}
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
		t.Errorf("got false wanted 17 got %d", v)
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
		t.Errorf("got false wanted 4 witdh 99 9:s and 8 after got %s", str)
	}
}

func TestFactorial(t *testing.T) {
	v := Factorial(5)
	if v != 120 {
		t.Errorf("got false wanted 120 got %d", v)
	}
}

func TestFactorialBig(t *testing.T) {
	v := FactorialBig(big.NewInt(50))
	if fmt.Sprintf("%d", v) != "30414093201713378043612608166064768844377641568960512000000000000" {
		t.Errorf("got false got %d", v)
	}
}

// SortString input "ÖÄÅöäåabcABC" output "ABCabcÄÅÖäåö"
func TestSortString(t *testing.T) {
	got := SortString("ÖÄÅöäåabcABC")
	want := "ABCabcÄÅÖäåö"
	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}

func TestStrip(t *testing.T) {
	got := Strip("A man, a plan, a canal: Panama åäöÅÄÖ")
	want := "AmanaplanacanalPanama"
	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}

func TestStripSwe(t *testing.T) {
	got := StripSwe("A man, a plan, a canal: Panama åäöÅÄÖ")
	want := "AmanaplanacanalPanamaåäöÅÄÖ"
	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}

func TestMap(t *testing.T) {
	s := []int{2, 4, 8, 11}
	ds := Map(s, func(i int) string { return strconv.Itoa(2 * i) })
	got := ds[0]
	want := "2"
	if got != want {
		t.Errorf("got %v wanted %s\n", got, want)
	}
}

func TestFilter(t *testing.T) {
	s := []int{2, 4, 8, 11}
	evens := Filter(s, func(i int) bool { return i%2 != 0 })
	got := evens[0]
	want := 11
	if got != want {
		t.Errorf("got %d wanted %d\n", got, want)
	}
}

func TestReduce(t *testing.T) {
	s := []int{2, 4, 8, 11}
	got := Reduce(s, 1, func(a, b int) int { return a * b })
	want := 704
	if got != want {
		t.Errorf("got %d wanted %d\n", got, want)
	}
}

func TestFind(t *testing.T) {
	people := []string{"Kent Beck", "Martin Fowler", "Chris James"}
	got, found := Find(people, func(s string) bool {
		return strings.Contains(s, "Chris")
	})
	want := true
	if found != want {
		t.Errorf("got %v wanted %v\n", got, want)
	}
}

func TestIf(t *testing.T) {
	gender := "male"
	sex := If(gender == "male", "boy", "girl")
	want := "boy"
	if sex != want {
		t.Errorf("got %s wanted %s\n", sex, want)
	}
}

func TestIf2(t *testing.T) {
	a, b := "a", "b"
	got := If(strings.Compare(a, b) == 0, a, b)
	want := "b"
	if got != want {
		t.Errorf("got %s wanted %s\n", got, want)
	}
}
