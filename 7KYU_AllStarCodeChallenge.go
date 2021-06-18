/**
Alexander Chelpanov
Date: 18.06.2021

Write a function, called sumPPG, that takes two NBA player objects/struct/Hash/Dict/Class and sums their PPG

Examples:
iverson := NBAPlayer{ Team: "76ers", Ppg: 11.2 }
jordan  := NBAPlayer{ Team: "bulls", Ppg: 20.2 }
SumPpg(iverson,jordan) // => 31.4
*/
package main

import "fmt"

func main() {
	iverson := NBAPlayer{Team: "76ers", Ppg: 11.2}
	jordan := NBAPlayer{Team: "bulls", Ppg: 20.2}
	fmt.Println(SumPpg(iverson, jordan))
}

type NBAPlayer struct {
	Team string
	Ppg  float64
}

func SumPpg(playerOne, playerTwo NBAPlayer) float64 {
	return playerOne.Ppg + playerTwo.Ppg
}
