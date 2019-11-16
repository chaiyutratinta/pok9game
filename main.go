package main

import (
	"fmt"
	"os"
	"pok9game/domain"
	"pok9game/utils"
)

func main() {
	allCard := [52]string{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9", "a10", "aj", "aq", "ak", "b1", "b2", "b3", "b4", "b5", "b6", "b7", "b8", "b9", "b10", "bj", "bq", "bk", "c1", "c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "c10", "cj", "cq", "ck", "d1", "d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "d10", "dj", "dq", "dk"}
	const playerAmount int = 5
	const lead int = 2
	const amountPerRound int = 10
	file, err := os.Create("./game.log")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	gameInit := domain.GameStart(playerAmount, lead, amountPerRound)

	round := 1

	endGame := false
	for !endGame {
		fmt.Printf("\nround %d\n", round)
		shuffledCards := utils.ShuffleCard(allCard)
		file.WriteString(fmt.Sprintf("\nround %d\n", round))
		utils.DealOutCards(&gameInit, &shuffledCards)
		domain.GamePlay(&gameInit, &shuffledCards, &endGame)
		domain.GetReward(&gameInit, amountPerRound, file)

		checkSum := 0
		for _, player := range gameInit {
			checkSum += player.Money
		}
		
		str := fmt.Sprintf("sum check >>> %d\n", checkSum)
		fmt.Print(str)
		file.WriteString(str)

		round++
	}
}
