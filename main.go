package main

import (
	"errors"
	"fmt"
	"github.com/thestuckster/gopherfacts/pkg/clients"
	"os"
	"sync"
	"time"
)

func main() {
	time.Sleep(58 * time.Second)

	fmt.Println("starting gopherfacts")
	token := os.Getenv("TOKEN")
	character := "Billbert"
	minerCharacter := "AMiner"
	client := clients.NewClient(&token)

	var wg sync.WaitGroup
	wg.Add(3)

	go farmChickens(character, client, &wg)
	go farmCopper(minerCharacter, client, &wg)

	wg.Wait()
}

func farmChickens(name string, client *clients.GopherFactClient, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("farming chickens")
	_, err := client.EasyClient.MoveToChickens(name)
	if err != nil {
		panic(err)
	}

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
			} else {
				panic(err)
			}
		}

		if err == nil {
			fmt.Printf("FIGHT: turn %d: Got %d xp from fight\n", turns, fightData.Fight.Xp)
			coolDown := fightData.Cooldown.TotalSeconds
			fmt.Printf("Fighting done. Cooling down for %d seconds\n", coolDown)
			time.Sleep(time.Duration(coolDown) * time.Second)
			turns++
		}
	}

}

func farmCopper(name string, client *clients.GopherFactClient, wg *sync.WaitGroup) {
	wg.Done()

	fmt.Println("farming copper")
	turns := 0
	for {
		gatherData, err := client.EasyClient.MineCopper(name)
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
			} else {
				panic(err)
			}
		}

		if err == nil {
			message := fmt.Sprintf("MINING: turn %d: Got %d xp from gather\n and looted:\n", turns, gatherData.Details.XpGained)
			for _, item := range gatherData.Details.Items {
				message += fmt.Sprintf("    %d %s\n", item.Quantity, item.Code)
			}

			println(message)
			turns++
		}
	}
}

func dumpInventoryIntoBank(name string, client *clients.GopherFactClient) {

	charData, err := client.CharacterClient.GetCharacterInfo(name)
	if err != nil {
		panic(err)
	}

	inventory := charData.Inventory
	for _, item := range inventory {
		fmt.Printf("Dumping all %s into bank\n", item.Code)
		if item.Code != "" {
			_, err := client.EasyClient.DepositIntoBank(name, item.Code, item.Quantity)
			if err != nil {
				panic(err)
			}
		}
	}
}
