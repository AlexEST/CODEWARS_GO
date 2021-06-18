/**

Alexander Chelpanov
Date: 18.06.2021

Write a program to determine if a string contains only unique characters. Return true if it does and false otherwise.
The string may contain any of the 128 ASCII characters. Characters are case-sensitive, e.g. 'a' and 'A'
are considered different characters.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(HasUniqueChar("  nAa"))
	fmt.Println(HasUniqueChar("abcde"))
}

func HasUniqueChar(str string) bool {
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		if strings.Count(str, string(arr[i])) > 1 {
			return false
		}
	}
	return true
}
