/**

Alexander Chelpanov
Date: 18.06.2021

Your task is to write function factorial.
https://en.wikipedia.org/wiki/Factorial
*/
package main

import "fmt"

func main() {
	fmt.Println(Factorial(4))
}

func Factorial(n int) int {
	num := 1
	for i := 1; i < n+1; i++ {
		num *= i
	}
	return num
}
