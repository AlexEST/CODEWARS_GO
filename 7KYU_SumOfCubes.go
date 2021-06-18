/**
Alexander Chelpanov
Date: 18.06.2021

Write a function that takes a positive integer n, sums all the cubed values from 1 to n, and returns that sum.
Assume that the input n will always be a positive integer.

Examples:
SumCubes(2) == 9
// sum of the cubes of 1 and 2 is 1 + 8
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(SumCubes(10))
}

func SumCubes(n int) int {
	number := 0
	for i := 1; i < n+1; i++ {
		number += i * i * i
	}
	return int(number)
}
