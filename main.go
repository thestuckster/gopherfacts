package main

import (
	"fmt"
	"gopherfacts/pkg/clients"
)

func main() {
	token := "MyToken"
	character := "Billbert"
	client := clients.NewClient(&token)

	//TODO: throwing 422 for some reason. need to debug
	fightData, err := client.MyCharacterClient.Fight(character)
	if err != nil {
		panic(err)
	}

	fmt.Println(fightData)
}
