//will handle log, intereation with user
package interaction

import (
	"fmt"
	"os"
)

type RoundData struct {
	Action 				string
	PlayerAttackDmg 	int
	PlayerHealValue 	int
	MonsterAttackDmg 	int
	PlayerHealth 		int
	MonsterHealth 		int
}

func PrintGreeting () {
	//Capticla function, will be import in other package
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Loadign to the game")
	fmt.Println("Good luck....")
	
}

func ShowAvailableActions(speacialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("--------------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if(speacialAttackIsAvailable){
		fmt.Println("(3) Special Attack")
	}

}

func DeclareWinner(winner string) {
	fmt.Println("------------------------------------------------")
	fmt.Println("GAME OVER")
	fmt.Println("------------------------------------------------")
	fmt.Printf("%v won!",winner)
}



func (roundData *RoundData) PrintRoudStatistics() {
	if(roundData.Action=="ATTACK"){
		fmt.Printf("Player attacked monster for %v damage.\n",roundData.PlayerAttackDmg)
	} else if(roundData.Action=="SPECIAL_ATTACK"){
		fmt.Printf("Player special attacked monster for %v damage.\n",roundData.PlayerAttackDmg)
		} else if(roundData.Action=="HEAL"){
		fmt.Printf("Player heal for %v.\n",roundData.PlayerHealValue)
	}
	fmt.Printf("Monster attacked player for %v damage.\n",roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v\n",roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n",roundData.MonsterHealth)
}

//write a RoundData array to a file
func WriteLogFile (rounds *[]RoundData) {
	//we can pass pointer in this fnc, or just the object itself
	file,err :=os.Create("gamelog.txt")
	if err != nil {
		println("Error creating file")
	}

	for index,value := range *rounds {
		//each value is a roundData object, should convert into a map of string both key,value; bc some value in roundData is an integer
		logEntry := map[string]string{
			"Round": fmt.Sprint(index+1),
			"Action":value.Action,
			"Player Attack Damage": fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value": fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health": fmt.Sprint(value.PlayerHealth),
			"Monster Health": fmt.Sprint(value.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry) // convert logEntry map into string
		_,err = file.WriteString(logLine)
		if err != nil {
			fmt.Println("Writing into log file failed. Exiting")
			continue // then we move on to next round of loop  **dont break the loop!
		}
	}



	// file.WriteString(content)

	file.Close()
	fmt.Println("Wrote data to log!")

	

}