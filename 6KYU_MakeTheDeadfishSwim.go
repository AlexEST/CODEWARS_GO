/**
Alexander Chelpanov
Date: 18.06.2021

Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored.

Parse("iiisdoso") == []int{8, 64}
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Parse("idoiido"))
	fmt.Println(Parse("iiisdoso"))
	fmt.Println(Parse("codewars"))
}

func Parse(data string) []int {
	arr := strings.Split(data, "")
	retArr := []int{}
	num := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == "i" {
			num++
		} else if arr[i] == "d" {
			num--
		} else if arr[i] == "s" {
			num *= num
		} else if arr[i] == "o" {
			retArr = append(retArr, num)
		}
	}
	return retArr
}
