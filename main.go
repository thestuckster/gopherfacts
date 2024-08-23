package main

import (
	"fmt"
	"gopherfacts/pkg/clients"
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
		fightData, err := client.MyCharacterClient.Fight(name)
		if err != nil {
			panic(err)
		}

		fmt.Printf("turn %d: Got %d xp from fight\n", turns, fightData.Fight.Xp)
		coolDown := fightData.Cooldown.TotalSeconds
		fmt.Printf("Cooling down for %d seconds\n", coolDown)
		time.Sleep(time.Duration(coolDown) * time.Second)
		turns++
	}

}
