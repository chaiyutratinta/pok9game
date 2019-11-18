package main

import (
	"fmt"
	"os"
	"pok9game/domain"
	"pok9game/utils"
)

func main() {
	allCard := [52]string{"♣1", "♣2", "♣3", "♣4", "♣5", "♣6", "♣7", "♣8", "♣9", "♣10", "♣j", "♣q", "♣k", "♦1", "♦2", "♦3", "♦4", "♦5", "♦6", "♦7", "♦8", "♦9", "♦10", "♦j", "♦q", "♦k", "♥1", "♥2", "♥3", "♥4", "♥5", "♥6", "♥7", "♥8", "♥9", "♥10", "♥j", "♥q", "♥k", "♠1", "♠2", "♠3", "♠4", "♠5", "♠6", "♠7", "♠8", "♠9", "♠10", "♠j", "♠q", "♠k"}
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
