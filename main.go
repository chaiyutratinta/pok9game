package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Player detail
type Player struct {
	name string
	money int
	lead bool
	card []string
}

func main() {
	playerAmount := 5
	allCard := []string{"a1","a2","a3","a4","a5","a6","a7","a8","a9","a10","aj","aq","ak","b1","b2","b3","b4","b5","b6","b7","b8","b9","b10","bj","bq","bk","c1","c2","c3","c4","c5","c6","c7","c8","c9","c10","cj","cq","ck","d1","d2","d3","d4","d5","d6","d7","d8","d9","d10","dj","dq","dk",}
	var game [5]Player
	rand.Seed(time.Now().UnixNano())
	for round := 0; round < 2; round++ {
		for player := 0; player < playerAmount; player++ {
			game[player] = Player{
				money: 100, 
				card: append(game[player].card, allCard[rand.Intn(52)]), 
				name: fmt.Sprintf("player%d", player+1)}
		}
	}

	fmt.Println(game)
}