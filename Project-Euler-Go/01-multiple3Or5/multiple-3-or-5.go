/*
 If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import (
	"fmt"
)

func main() {
	// Modify n to increment or decrease list of natural numbers
	var n int
	fmt.Print("n: ")
	fmt.Scanln(&n)
	fmt.Println("Sum of multiples of 3 or 5:", Multiples(n))
}

// Multiples: Returns the multiples of 3 or 5 below the given number
func Multiples(num int) int {
	var sumMultiples int = 0

	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sumMultiples += i
		}
	}
	return sumMultiples
}
