package clients

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const chickens = "0:1"
const cookingLocation = "1:1"
const weaponCraftingLocation = "2:1"
const gearCraftingLocation = "3:1"
const bankLocation = "4:1"
const geLocation = "5:1"
const jewelryLocation = "1:3"
const miningLocation = "1:5"

type EasyClient struct {
	token      *string
	charClient *CharacterClient
}

func (c *EasyClient) DepositIntoBank(characterName, itemCode string, amount int) (*DepositData, Error) {
	_, err := c.MoveToBank(characterName)
	if err != nil {
		return nil, err
	}

	depositData, err := c.charClient.DepositIntoBank(characterName, itemCode, amount)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(depositData.Cooldown.RemainingSeconds) * time.Second)
	return depositData, nil
}

func (c *EasyClient) MoveToChickens(characterName string) (*MoveData, Error) {
	return c.moveToLocation(characterName, chickens)
}

func (c *EasyClient) MoveToBank(characterName string) (*MoveData, Error) {
	return c.moveToLocation(characterName, bankLocation)
}

func (c *EasyClient) moveToLocation(characterName, coords string) (*MoveData, Error) {
	x, y := getCoords(coords)
	moveResp, err := c.charClient.Move(characterName, x, y)
	if err != nil {
		if ex, ok := err.(*CharacterAlreadyAtDestinationException); ok {
			// err is of type *CharacterAlreadyAtDestinationException
			fmt.Println(ex.Message)
			return nil, nil
		} else {
			// Handle other errors
			return nil, err
		}
	}

	coolDown := moveResp.Cooldown.RemainingSeconds
	time.Sleep(time.Duration(coolDown) * time.Second)

	return moveResp, nil
}

func getCoords(coord string) (x, y int) {

	parts := strings.Split(coord, ":")
	x, _ = strconv.Atoi(parts[0])
	y, _ = strconv.Atoi(parts[1])
	return x, y
}
