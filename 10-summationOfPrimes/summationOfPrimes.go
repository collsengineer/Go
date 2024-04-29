/*
The sum of the primes below 10 is:
    2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million (2000000).
*/

package main

import (
	"fmt"
	"math"
)

// Modify start and end to set the primes's list range.

const (
	start = 2
	end   = 2000000
)

func main() {
	primesSlice := PrimesFromRange(start, end)
	sumOfPrimes := Sum(primesSlice)

	fmt.Printf("Prime's sum below %d = %d\n", end, sumOfPrimes)
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
			primes = append(primes, start)
		}
		start++
	}
	return primes
}

func Sum(primesSlice []int) int {
	var total int

	for _, val := range primesSlice {
		total += val
	}
	return total
}
