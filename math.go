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
