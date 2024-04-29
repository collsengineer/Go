/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143?
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// Number from which prime factors come from.
	n := 600851475143

	// Prime factors sequence.
	primeFactorsSeq := PrimeFactors(n)

	// Largest prime factor.
	maxFactor := Max(primeFactorsSeq)

	fmt.Printf("Max Prime factor: %d\n", maxFactor)
}

// PrimeFactors: Returns the prime factors of n
func PrimeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
			i--
		}
	}
	if n > 0 {
		factors = append(factors, n)
	}
	return factors
}

// Max: Returns the max element within a slice
func Max(factors []int) int {
	max := factors[0]

	for _, factor := range factors {
		if factor > max {
			max = factor
		}
	}
	return max
}
