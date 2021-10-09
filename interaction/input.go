package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(speacialAttackIsAvailable bool) string {
	
	//run true loop until user get the right choice! and if "return"! to breakout the loop
	for {
		//getPlayerInput 
		playerChoice, _ := getPlayerInput() 

		//check if playerChoice is in the avail rank
		if playerChoice == "1"{
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && speacialAttackIsAvailable {
			return "SPECIAL_ATTACK"
		} else {
			fmt.Println("Fetching the user input failed, please try again!")
		}
	}
	
}


func getPlayerInput()(string,error){
	fmt.Print("Your choice:")
	userInput,err := reader.ReadString('\n')
	if err != nil {
		return "",err
	}
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	return userInput, nil
}