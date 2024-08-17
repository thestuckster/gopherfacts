package main

import (
	"fmt"
	"gopherfacts/pkg/clients"
)

func main() {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6InRoZXN0dWNrc3R0ZXIiLCJwYXNzd29yZF9jaGFuZ2VkIjoiIn0.g3qYaJOY40MEjUu2234JeXYcY-1HeboS7LmGQcHhnuk"
	character := "Billbert"
	client := clients.NewClient(&token)

	//TODO: throwing 422 for some reason. need to debug
	fightData, err := client.MyCharacterClient.Fight(character)
	if err != nil {
		panic(err)
	}

	fmt.Println(fightData)
}
