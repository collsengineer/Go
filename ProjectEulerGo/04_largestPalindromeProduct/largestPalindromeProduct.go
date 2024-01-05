/*
A palindromic number reads the same both ways. The largest palindrome made
from the product of two-digit numbers is 909 = 91 * 99

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// if nDigitsNumbers == "3", it means 125 x 245 or 385 x 999, etc.
	nDigitsNumbers := "3"
	fmt.Println("Max palindrome product:", MaxPalindromeProductOf(nDigitsNumbers))
}

func Reverse(s string) (result string) {
	// Reverts the characters on a string.
	for _, char := range s {
		result = string(char) + result
	}
	return
}

func IsPalindrome(str string) interface{} {
	if str == Reverse(str) {
		return true
	}
	return false
}

func LargestElement(slice []int) int {
	// Finds the largest value of a Slice.
	var max int
	max = slice[0]
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}

func MaxPalindromeProductOf(nDigits string) int {
	// It checks if the product is palindromic, and returns the max value
	a := []int{}
	b := []int{}
	aTimesB := []int{}
	ABStr := []string{}
	ABPal := []int{}

	// Generating slices a and b with ints from 1 to 1000.
	if nDigits == "3" {

		for i := 1; i <= 999; i++ {
			a = append(a, i)
		}

		for i := 1; i <= 999; i++ {
			b = append(b, i)
		}
	}

	// Multiplying each a element by each b element
	for _, i := range a {
		for _, j := range b {
			product := i * j
			aTimesB = append(aTimesB, product)
		}
	}

	// Converting int elements to string
	for _, product := range aTimesB {
		toStr := strconv.Itoa(product)
		ABStr = append(ABStr, toStr)
	}

	// Check if product is palindrome -> add to ABPal
	for _, product := range ABStr {
		if IsPalindrome(product) == true {
			toInt, _ := strconv.Atoi(product)
			ABPal = append(ABPal, toInt)
		}
	}

	return LargestElement(ABPal)
}
