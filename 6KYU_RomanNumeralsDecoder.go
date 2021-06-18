/**
Alexander Chelpanov
Date: 18.06.2021

Create a function that takes a Roman numeral as its argument and returns its value as a numeric decimal integer.
You don't need to validate the form of the Roman numeral.
Modern Roman numerals are written by expressing each decimal digit of the number to be encoded separately,
starting with the leftmost digit and skipping any 0s. So 1990 is rendered "MCMXC" (1000 = M, 900 = CM, 90 = XC)
and 2008 is rendered "MMVIII" (2000 = MM, 8 = VIII). The Roman numeral for 1666, "MDCLXVI",
uses each letter in descending order.

Example:
solution('XXI'); // should return 21
*/

package main

import (
	"fmt"
	"strings"
)

var romanNumeralDict = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func main() {
	fmt.Println(Decode("XXI"))
	fmt.Println(Decode("MMVIII"))
	fmt.Println(Decode("VI"))
	fmt.Println(Decode("IV"))
}

func Decode(roman string) int {
	arr := strings.Split(roman, "")
	num := 0
	for i := 0; i < len(arr); i++ {
		if i < len(arr)-1 {
			str := arr[i] + arr[i+1]
			if romanNumeralDict[str] != 0 {
				num += romanNumeralDict[str]
				i++
			} else {
				num += romanNumeralDict[arr[i]]
			}
		} else {
			num += romanNumeralDict[arr[i]]
		}
	}
	return num
}
