package main

import (
	"fmt"
	"gopherfacts/pkg/clients"
)

func main() {
	token := "MyToken"
	character := "Billbert"
	client := clients.NewClient(&token)

	err := client.MyCharacterClient.Move(character, -1, 0)
	if err != nil {
		panic(err)
	}

	gatherData, err := client.MyCharacterClient.Gather(character)
	if err != nil {
		panic(err)
	}

	prettyPrintStruct(gatherData)
}


func prettyPrintStruct(data any) {
	fmt.Printf("%#v\n", data)
}