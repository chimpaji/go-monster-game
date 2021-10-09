package interaction

import "fmt"

//will handle log, intereation with user

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