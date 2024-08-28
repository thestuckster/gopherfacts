package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/thestuckster/gopherfacts/pkg/clients"
	"os"
	"sync"
	"time"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

func main() {

	logger.Info().Msg("Starting Gopherfacts...")

	wait := 60
	logger.Debug().Msgf("Waiting for %d seconds for all cooldowns to finish", wait)
	time.Sleep(time.Duration(wait) * time.Second)

	token := os.Getenv("TOKEN")
	character := "Billbert"
	minerCharacter := "AMiner"
	//crafter := "smellyfeet"

	client := clients.NewClient(&token)

	var wg sync.WaitGroup
	wg.Add(3)

	go farmChickens(character, client, &wg)
	go farmCopper(minerCharacter, client, &wg)

	wg.Wait()
}

func farmChickens(name string, client *clients.GopherFactClient, wg *sync.WaitGroup) {
	defer wg.Done()

	chickenLogger := logger.With().Str("method", "farmChickens").Str("character", name).Logger()

	chickenLogger.Info().Msg("Farming Chickens")
	_, err := client.EasyClient.MoveToChickens(name)
	if err != nil {
		chickenLogger.Error().Err(err).Msg("Farming Chickens failed")
		panic(err)
	}

	turns := 0
	for {
		fightData, err := client.CharacterClient.Fight(name)
		if err != nil {
			var ex *clients.CharacterInventoryFullException
			if errors.As(err, &ex) {
				chickenLogger.Debug().Msg("Inventory full")
				dumpInventoryIntoBank(name, client)
				_, err := client.EasyClient.MoveToChickens(name)
				if err != nil {
					chickenLogger.Error().Err(err).Msg("Moving back to chickens after bank dump failed")
					panic(err)
				}
			} else {
				chickenLogger.Error().Err(err).Msg("Fighting chickens failed")
				panic(err)
			}
		}

		if err == nil {
			chickenLogger.Info().Msgf("turn %d: Got %d xp from fight\n", turns, fightData.Fight.Xp)
			coolDown := fightData.Cooldown.TotalSeconds
			chickenLogger.Info().Msgf("Fighting done. Cooling down for %d seconds\n", coolDown)
			time.Sleep(time.Duration(coolDown) * time.Second)
			turns++
		}
	}

}

func farmCopper(name string, client *clients.GopherFactClient, wg *sync.WaitGroup) {
	defer wg.Done()

	copperLogger := logger.With().Str("method", "farmCopper").Str("character", name).Logger()

	copperLogger.Info().Msg("Farming Copper")
	turns := 0
	for {
		gatherData, err := client.EasyClient.MineCopper(name)
		if err != nil {
			var ex *clients.CharacterInventoryFullException
			if errors.As(err, &ex) {
				copperLogger.Debug().Msg("Inventory full")
				dumpInventoryIntoBank(name, client)
				_, err := client.EasyClient.MoveToChickens(name)
				if err != nil {
					copperLogger.Debug().Err(err).Msg("Moving back to copper after bank dump failed")
					panic(err)
				}
			} else {
				copperLogger.Debug().Err(err).Msg("Farming Copper failed")
				panic(err)
			}
		}

		if err == nil {

			message := fmt.Sprintf("MINING: turn %d: Got %d xp from gather\n and looted:%v", turns, gatherData.Details.XpGained, gatherData.Details.Items)
			for _, item := range gatherData.Details.Items {
				message += fmt.Sprintf("    %d %s\n", item.Quantity, item.Code)
			}

			copperLogger.Info().Msg(message)
			turns++
		}
	}
}

func dumpInventoryIntoBank(name string, client *clients.GopherFactClient) {

	dumpLogger := logger.With().Str("method", "dumpInventoryIntoBank").Str("character", name).Logger()

	charData, err := client.CharacterClient.GetCharacterInfo(name)
	if err != nil {
		dumpLogger.Error().Err(err).Msg("Character info failed")
		panic(err)
	}

	inventory := charData.Inventory
	for _, item := range inventory {
		//TODO: Item.Code seems to be logging as an empty string... but the deposit code is working... weird.
		dumpLogger.Info().Msgf("Dumping all %s into bank", item.Code)
		if item.Code != "" {
			_, err := client.EasyClient.DepositIntoBank(name, item.Code, item.Quantity)
			if err != nil {
				dumpLogger.Error().Err(err).Msg("Deposit into bank failed")
				panic(err)
			}
		}
	}
}
