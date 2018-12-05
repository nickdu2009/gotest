package main

import (
	"github.com/nickdu2009/gotest/copy/player"
	"log"
)

func main() {
	originPlayer := player.GetOriginPlayer()

	copyPlayer := player.GetPlayerWithCopy()

	copyPlayer.GoodAt = "copy"
	clonePlayer := player.GetPlayerWithClone()

	clonePlayer.GoodAt = "clone"

	log.Printf("originPlayer: \n %v", originPlayer)
	log.Printf("copyPlayer: \n %v", copyPlayer)
	log.Printf("clonePlayer: \n %v", clonePlayer)
}
