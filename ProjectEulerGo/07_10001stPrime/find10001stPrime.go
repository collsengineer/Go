package main

import (
	"fmt"
	"math"
)

func main() {
	// start and end define the primes's range.
	start := 2
	end := 1000000
	primesList := PrimesFromRange(start, end)

	// Choose the prime number's desired position
	targetPosition := 10000
	GetByIndex(targetPosition, primesList)
}

func PrimesFromRange(start, end int) []int {
	primes := []int{}

	if start < 2 || end < 2 {
		fmt.Println("Numbers must be greater than 2")
	}

	for start <= end {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(start))); i++ {
			if start%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			// Uncomment below line to see all the primes listed.
			// fmt.Printf("%d ", start)
			primes = append(primes, start)
		}
		start++
	}
	// fmt.Println()
	return primes
}

func GetByIndex(position int, primes []int) {
	for index, prime := range primes {
		if index == position {
			fmt.Printf("Prime = %d\nPosition = %d\n", prime, index+1)
			break
		}
	}
}
