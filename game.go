package main

import (
	"fmt"

	interaction "github.com/chimpaji/go-monster-game/interaction"
)

var currentRound int = 0;

func main() {
	getUserAction()

}

func getUserAction() {
	//atakc, heal, spacial attack(avail every 3 rounds)
	//all with random number calc
	startGame()

	//loop as long as there's a winner, then we end the game
	winner := "" // "player" || "monster"

	for winner == "" {
		fmt.Printf("\nWinner:%v\n",winner)
		winner = executeRound()

	}

	endGame()
}

//update monster health

//monster attacks player

//update player health

//check winner, health > 0

func startGame() {
	//greeting
	interaction.PrintGreeting()
}

func endGame() {}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound % 3 == 0 //every 3 round, let user do special action
	
	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	fmt.Printf("User Choice:%v",userChoice)

	//each action will trigger function from action package
	if userChoice == "ATTACK"{

	} else if userChoice == "HEAL" {

	} else if userChoice == "SPEACIAL_ATTACK"{

	}
	//return string of winner
	return ""
}