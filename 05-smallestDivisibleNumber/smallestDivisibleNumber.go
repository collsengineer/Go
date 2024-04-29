/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10
without any remainder.

What is the smallest positive number that is evenly divisible by
all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
)

func main() {
	limit := 1000000000
	mySlice := []int{}
	evenlyDivisibleBy1To10 := EvenDivisible(mySlice, limit)
	fmt.Println(Min(evenlyDivisibleBy1To10))
}

func Min(slice []int) int {
	var min int
	min = slice[0]

	for _, val := range slice {
		if val < min {
			min = val
		}
	}
	return min
}

func EvenDivisible(slice []int, limit int) []int {
	slice = []int{}

	for value := 1; value <= limit; value++ {
		if value%1 == 0 &&
			value%2 == 0 &&
			value%3 == 0 &&
			value%4 == 0 &&
			value%5 == 0 &&
			value%6 == 0 &&
			value%7 == 0 &&
			value%8 == 0 &&
			value%9 == 0 &&
			value%10 == 0 &&
			value%11 == 0 &&
			value%12 == 0 &&
			value%13 == 0 &&
			value%14 == 0 &&
			value%15 == 0 &&
			value%16 == 0 &&
			value%17 == 0 &&
			value%18 == 0 &&
			value%19 == 0 &&
			value%20 == 0 {
			slice = append(slice, value)
		}
	}
	return slice
}
