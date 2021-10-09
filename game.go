package main

import (
	"fmt"

	"github.com/chimpaji/go-monster-game/actions"
	interaction "github.com/chimpaji/go-monster-game/interaction"
)

var currentRound int = 0;
var gameRounds = []interaction.RoundData{} // for record round data in array to write log file when game done

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
		winner = executeRound()

	}
	//print out show winner
	endGame(winner)
}

func startGame() {
	//greeting
	interaction.PrintGreeting()
}

func endGame(winner string) {
	//log the winner
	interaction.DeclareWinner(winner)
	//write round log into a file
	interaction.WriteLogFile(&gameRounds)

}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound % 3 == 0 //every 3 round, let user do special action
	
	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	fmt.Printf("User Choice:%v\n",userChoice)

	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	//each action will trigger function from action package
	if userChoice == "ATTACK"{
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()	
	} else{
		//if user choose speacial attact
		playerAttackDmg = actions.AttackMonster(true)
	}

	//monster will always attack user 
	monsterAttackDmg = actions.AttackPlayer()
	//get lastest player, monster health
	playerHealth,monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData {
		Action:userChoice,
	PlayerAttackDmg:playerAttackDmg,
	PlayerHealValue:playerHealValue,
	MonsterAttackDmg:monsterAttackDmg,
	PlayerHealth:playerHealth,
	MonsterHealth:monsterHealth,
	}
	roundData.PrintRoudStatistics()
	gameRounds = append(gameRounds, roundData)

	//return string of winner
	if playerHealth <= 0{
		return "MONSTER"
	} else if monsterHealth <= 0{
		return "PLAYER"
	}
	//no winner yet
	return ""
}