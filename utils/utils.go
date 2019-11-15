package utils

import (
	"strings"
	"strconv"
	"pok9game/model"
	"math/rand"
	// "fmt"
)

//GetPoint is calculate point on hand of player
func GetPoint(cardOnhand []string) int{
	sum := 0
	for i := 0; i < len(cardOnhand); i++ {
		removeFirstAlphabet := strings.Split(cardOnhand[i], "")
		joinNum := strings.Join(removeFirstAlphabet[1:], "")
		getNum, err := strconv.Atoi(joinNum)

		if err != nil {
			sum += 0
		}

		sum += getNum
	}

	sumStr := strconv.Itoa(sum)
	getFinaly := strings.Split(sumStr, "")
	almostFinaly := strings.Join(getFinaly[len(getFinaly) - 1:len(getFinaly)], "")
	reallyFinaly, err := strconv.Atoi(almostFinaly)
	if err != nil {
		// fmt.Println(err)
		return 0
	}
	return reallyFinaly
}

//DrawCard for ...
func DrawCard(player *model.Player, cardOnhand []string) {
  allCard := []string{"a1","a2","a3","a4","a5","a6","a7","a8","a9","a10","aj","aq","ak","b1","b2","b3","b4","b5","b6","b7","b8","b9","b10","bj","bq","bk","c1","c2","c3","c4","c5","c6","c7","c8","c9","c10","cj","cq","ck","d1","d2","d3","d4","d5","d6","d7","d8","d9","d10","dj","dq","dk",}
  player.Card = append(cardOnhand, allCard[rand.Intn(52)])
}