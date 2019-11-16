package domain

import (
	"fmt"
	"os"
	"pok9game/model"
	"pok9game/utils"
)

//GameStart for initial player
func GameStart(playerAmount, lead, amountPerRound int) []model.Player {
	gameInit := make([]model.Player, playerAmount)

	for player := 0; player < playerAmount; player++ {
		modelPlayer := &model.Player{}
		modelPlayer.Name = fmt.Sprintf("player%d", player+1)
		modelPlayer.Money = 100

		if player == lead {
			modelPlayer.Lead = true
		}

		gameInit[player] = *modelPlayer
	}

	return gameInit
}

//GetReward for collected reward
func GetReward(game *[]model.Player, amountPerRound int, files *os.File) {
	getGame := *game
	playerAmount := len(getGame)
	leadIdx := 0
	leadPoint := 0
	var leadMoney *int

	for idx, player := range getGame {
		if player.Lead {
			leadPoint = utils.GetPoint(player.Card)
			leadIdx = idx
			leadMoney = &getGame[idx].Money
		}
	}

	for idx := 0; idx < playerAmount; idx++ {
		playerMoney := &getGame[idx].Money
		playPoint := utils.GetPoint(getGame[idx].Card)

		if !getGame[idx].Lead {
			str := fmt.Sprintf("%d vs %d | cards: %v ", leadPoint, playPoint, getGame[idx].Card)
			fmt.Print(str)
			files.WriteString(str)

			if leadPoint < playPoint {
				*leadMoney = *leadMoney - amountPerRound
				*playerMoney = *playerMoney + amountPerRound
			} else if leadPoint > playPoint {
				*leadMoney = *leadMoney + amountPerRound
				*playerMoney = *playerMoney - amountPerRound
			}
			str2 := fmt.Sprintf("balance: %d\n", *playerMoney)
			fmt.Print(str2)
			files.WriteString(str2)
		}
		
		playerCard := &getGame[idx].Card
		*playerCard = []string{}
	}
	srt3 := fmt.Sprintf("lead   | cards: %v, balance: %d \n", getGame[leadIdx].Card, getGame[leadIdx].Money)
	fmt.Print(srt3)
	files.WriteString(srt3)
}

//GamePlay for ...
func GamePlay(game *[]model.Player, shuffledCards *[]string, endGame *bool) {
	getGame := *game

	for idx := 0; idx < len(getGame); idx++ {
		playerMoney := &getGame[idx].Money
		playerCard := &getGame[idx].Card
		
		if *playerMoney < 10 {
			*endGame = true
			break
		}
		
		point := utils.GetPoint(getGame[idx].Card)
		if point < 5 && len(*playerCard) < 3 {
			*playerCard = append(*playerCard, utils.DrawCard(shuffledCards))
			point = utils.GetPoint(*playerCard)
		}
	}
}
