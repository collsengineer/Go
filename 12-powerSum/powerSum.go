package main

import (
	"fmt"
	"math/big"
)

/*
   2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

   What's the sum of the digits of the number 2^1000?
*/

func main() {
	// 2^1000
	exp := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)

	total := sumOfDigits(exp)
	fmt.Printf("The sum of the digits of 2^1000 is: %d\n", total)
}

// sumOfDigits: Returns the sum of digits of x
func sumOfDigits(exp *big.Int) int {
	var total int

	expStr := exp.String()

	for _, charStr := range expStr {
		charInt := int(charStr - '0')
		total += charInt
	}
	return total
}
