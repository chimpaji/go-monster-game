package actions

import (
	"math/rand"
	"time"
)

// to generate random number we need source(a reference that are difference everytime it got called)
//and  generator(generator from the "source")
var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource) //randGenerator.Intn(upperLimitInt) will give us random numberthru Intn method
//create constant in seperate file wiht same package, this trick will cause less bugs
//can export these 2 value to main package to show value, but we might accidentally change these val, better return in in fnc
var currenctMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

 
func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG
	 
	if(isSpecialAttack){
		minAttackValue = PLAYER_ATTACK_SPEACIAL_MIN_DMG
		maxAttackValue = PLAYER_ATTACK_SPEACIAL_MAX_DMG
	}
	dmgValue := generateRandomNumber(minAttackValue, maxAttackValue)
	currenctMonsterHealth -= dmgValue
	return dmgValue
}

func HealPlayer() int {
	minHealValue := PLAYER_HEAL_MIN
	maxHealValue := PLAYER_HEAL_MAX

	healValue := generateRandomNumber(minHealValue, maxHealValue)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth
	if(healthDiff > healValue){
		currentPlayerHealth += healValue
	}else {
		currentPlayerHealth = PLAYER_HEALTH
	}
	return healValue
}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	dmgValue := generateRandomNumber(minAttackValue, maxAttackValue)
	currentPlayerHealth -= dmgValue
	return dmgValue
}

func GetHealthAmounts () (int, int) {
	return currentPlayerHealth, currenctMonsterHealth
}

func generateRandomNumber(min int , max int) int {
	return randGenerator.Intn(max-min) + min
}