package domain

import (
	"fmt"
	"pok9game/model"
	"pok9game/utils"
)

//GetReward for collected reward
func GetReward(game *[]model.Player, amountPerRound int) {
	getGame := *game
	playerAmount := len(getGame)
	totalReward := amountPerRound * playerAmount
	leadIdx := 0

	for p := 0; p < playerAmount; p++ {
		if getGame[p].Lead {
			leadIdx = p
		}
	}

	for player := 0; player < playerAmount; player++ {
		if !getGame[player].Lead {
			fmt.Printf("%d vs %d\n",getGame[leadIdx].Total, getGame[player].Total)
			money := &getGame[player].Money
			if getGame[leadIdx].Total < getGame[player].Total {
					totalReward = totalReward - amountPerRound
					leadMoney := &getGame[leadIdx].Money
					*leadMoney = *leadMoney - amountPerRound
					*money = *money + (amountPerRound *2)
			} else if getGame[leadIdx].Total == getGame[player].Total {
					totalReward = totalReward - amountPerRound
					*money = *money + amountPerRound
			}
		}
	}

	leadMoney := &getGame[leadIdx].Money
	*leadMoney = *leadMoney + totalReward
}

//GamePlay for ...
func GamePlay(game *[]model.Player) {
	getGame := *game
	gameLen := len(getGame)
	for i := 0; i < gameLen; i++ {
		money := &getGame[i].Money
    total := &getGame[i].Total
    *money = *money - 10
    point := utils.GetPoint(getGame[i].Card)
    if point < 5 {
      utils.DrawCard(&getGame[i], getGame[i].Card )
      point = utils.GetPoint(getGame[i].Card)
      *total = point
    }
    *total = point
	}
}