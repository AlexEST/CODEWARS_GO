/**
Alexander Chelpanov
Date: 18.06.2021

Your goal is to return multiplication table for number that is always an integer from 1 to 10.

For example, a multiplication table (string) for number == 5 looks like below:

1 * 5 = 5
2 * 5 = 10
3 * 5 = 15
4 * 5 = 20
5 * 5 = 25
6 * 5 = 30
7 * 5 = 35
8 * 5 = 40
9 * 5 = 45
10 * 5 = 50
P. S. You can use \n in string to jump to the next line.

Note: newlines should be added between rows, but there should be no trailing newline at the end.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MultiTable(5))
}

func MultiTable(number int) string {
	str := ""
	for i := 1; i < 11; i++ {
		if i == 10 {
			str += strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(i*number)
		} else {
			str += strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(i*number) + "\n"
		}
	}
	return str
}
