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
	crafter := "smellyfeet"

	client := clients.NewClient(&token)

	var wg sync.WaitGroup
	wg.Add(3)

	go farmChickens(character, client, &wg)
	go miningFarm(minerCharacter, client, &wg)
	go smeltCopper(crafter, client, &wg)

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

func miningFarm(name string, client *clients.GopherFactClient, wg *sync.WaitGroup) {
	defer wg.Done()

	miningLogger := logger.With().Str("method", "miningFarm").Str("character", name).Logger()
	miningLogger.Info().Msg("Mining Operation commencing....")

	turns := 0
	for {
		miningLogger.Debug().Msg("Fetching character info to determine best mining resource")
		characterData, err := client.CharacterClient.GetCharacterInfo(name)
		if err != nil {
			miningLogger.Error().Err(err).Msg("Fetching character info failed")
		}

		var gatherData *clients.GatherData
		miningLevel := characterData.MiningLevel
		miningLogger.Debug().Msgf("Character mining level is %d", miningLevel)

		if miningLevel < 11 {
			miningLogger.Info().Msg("Mining Copper")
			gatherData, err = client.EasyClient.MineCopper(name)
		} else {
			miningLogger.Info().Msg("Mining Iron")
			gatherData, err = client.EasyClient.MineIron(name)
		}

		if err != nil {
			var ex *clients.CharacterInventoryFullException
			if errors.As(err, &ex) {
				miningLogger.Debug().Msg("Inventory full")
				dumpInventoryIntoBank(name, client)
			} else {
				miningLogger.Debug().Err(err).Msg("Mining operation  failed")
				panic(err)
			}
		}

		if err == nil {

			message := fmt.Sprintf("MINING: turn %d: Got %d xp from gather\n and looted:%v", turns, gatherData.Details.XpGained, gatherData.Details.Items)
			miningLogger.Info().Msg(message)
			turns++
		}
	}
}

func smeltCopper(name string, client *clients.GopherFactClient, wg *sync.WaitGroup) {
	defer wg.Done()

	smeltLogger := logger.With().Str("method", "smeltCopper").Str("character", name).Logger()
	smeltLogger.Info().Msg("Smelting Copper")

	turns := 0
	for {
		_, err := client.EasyClient.WithdrawFromBank(name, "copper_ore", 100)
		if err != nil {
			smeltLogger.Error().Err(err).Msg("Moving to bank failed")
			break
		}

		_, err = client.EasyClient.MoveToForge(name)
		if err != nil {
			smeltLogger.Error().Err(err).Msg("Moving to forge failed")
			break
		}

		craftData, err := client.CharacterClient.Craft(name, "copper", 12)
		if err != nil {
			smeltLogger.Error().Err(err).Msg("Craft failed")
			break
		}

		details := craftData.Details
		item := details.Items[0] //since its just copper for now we assume only 1 item returned
		message := fmt.Sprintf("Turn %d: Crafted %d %s and got %d XP", turns, item.Quantity, item.Code, craftData.Details.XpGained)
		smeltLogger.Info().Msg(message)

		coolDown := craftData.Cooldown.RemainingSeconds
		time.Sleep(time.Duration(coolDown) * time.Second)

		dumpInventoryIntoBank(name, client)

		turns++
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
