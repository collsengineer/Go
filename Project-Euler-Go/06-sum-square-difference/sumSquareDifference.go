/*
The sum of the squares of the first ten natural numbers is:

        1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is:

        (1 + 2 + ... + 10)^2 = 55^2 = 3025


Hence the difference between the sum of the squares of the first ten
natural numbers and the square of the sum is:

        3025 - 385 = 2640


Find the difference between the sum of the squares of the first
one hundred natural numbers and the square of the sum
*/

package main

import (
	"fmt"
)

func main() {
	n := 100
	resultSumSquareDiff := SumSquareDifferenceOf(n)
	fmt.Printf("Sum of Squares difference: %v\n", resultSumSquareDiff)
}

func SumSquareDifferenceOf(naturalNumbers int) int {
	var sumOfSquares int
	var squareOfSum int

	for val := 1; val <= naturalNumbers; val++ {
		valSquared := val * val
		sumOfSquares += valSquared
		squareOfSum += val
	}
	return ((squareOfSum * squareOfSum) - sumOfSquares)
}
