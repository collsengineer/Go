/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which:

            a² + b² = c²

for example:

            3² + 4² = 9 + 16 = 25 = 5²

There's exactly one Pythagorean triplet for which a + b + c = 1000.

Find the product abc.
*/

package main

import (
	"fmt"
)

func main() {
	d := 1000
	abc := PythagoreanTripletProduct(d)

	fmt.Printf("Product abc = %d\n", abc)
}

func PythagoreanTripletProduct(d int) int {
	// a + b + c = d
	var product int

	for a := 1; a < d; a++ {
		for b := a; b < d-a; b++ {
			c := d - a - b
			if (a*a)+(b*b) == c*c {
				product = a * b * c
				break
			}
		}
	}
	return product
}
