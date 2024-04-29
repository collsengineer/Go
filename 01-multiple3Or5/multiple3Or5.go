/*
 If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {
	// Modify n to increment or decrease the list of natural numbers
	var n int
	fmt.Print("n: ")
	fmt.Scanln(&n)
	fmt.Printf("Sum of multiples of 3 or 5 below %d: %d\n", n, sumMultiples(n))
}

// Multiples: Returns the multiples of 3 or 5 below n
func sumMultiples(n int) int {
	var sum int

	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}
