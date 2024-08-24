package main

import (
	"errors"
	"fmt"
	"github.com/thestuckster/gopherfacts/pkg/clients"
	"os"
	"time"
)

func main() {
	token := os.Getenv("TOKEN")
	character := "Billbert"
	client := clients.NewClient(&token)

	farmChickens(character, client)
}

func farmChickens(name string, client *clients.GopherFactClient) {

	turns := 0
	for {
		fightData, err := client.CharacterClient.Fight(name)
		if err != nil {
			var ex *clients.CharacterInventoryFullException
			if errors.As(err, &ex) {
				fmt.Println("%%%%%")
				fmt.Println(ex.Message)
				dumpInventoryIntoBank(name, client)
				_, err := client.EasyClient.MoveToChickens(name)
				if err != nil {
					panic(err)
				}
			}
		}

		fmt.Printf("turn %d: Got %d xp from fight\n", turns, fightData.Fight.Xp)
		coolDown := fightData.Cooldown.TotalSeconds
		fmt.Printf("Cooling down for %d seconds\n", coolDown)
		time.Sleep(time.Duration(coolDown) * time.Second)
		turns++
	}

}

func dumpInventoryIntoBank(name string, client *clients.GopherFactClient) {

	charData, err := client.CharacterClient.GetCharacterInfo(name)
	if err != nil {
		panic(err)
	}

	inventory := charData.Inventory
	for _, item := range inventory {
		_, err := client.EasyClient.DepositIntoBank(name, item.Code, item.Quantity)
		if err != nil {
			panic(err)
		}
	}
}
