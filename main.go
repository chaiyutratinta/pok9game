package main

import (
	"fmt"
	"math/rand"
	"time"
	"pok9game/domain"
	"pok9game/model"
	"pok9game/utils"
)

func main() {
  playerAmount := 5
  lead := 2
  amountPerRound := 10
	gameInit := make([]model.Player, playerAmount)
	rand.Seed(time.Now().UnixNano())
	for round := 0; round < 2; round++ {
		for player := 0; player < playerAmount; player++ {
      modelPlayer := &model.Player{}
			utils.DrawCard(modelPlayer, gameInit[player].Card)
			modelPlayer.Name = fmt.Sprintf("player%d", player+1)
			modelPlayer.Money = 100
      if player == lead {
        modelPlayer.Lead = true
      }
			gameInit[player] = *modelPlayer
		}
	}

  domain.GamePlay(&gameInit)
  domain.GetReward(&gameInit, amountPerRound)
  fmt.Println(gameInit)
}