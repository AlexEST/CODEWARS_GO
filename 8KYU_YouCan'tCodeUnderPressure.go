/**
Alexander Chelpanov
Date: 18.06.2021

Code as fast as you can! You need to double the integer and return it.
*/
package main

import "fmt"

func main() {
	fmt.Println(DoubleInteger(1))
	fmt.Println(DoubleInteger(5))
}

func DoubleInteger(i int) int {
	return i * 2
}
