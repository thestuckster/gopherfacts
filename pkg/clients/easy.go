package clients

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
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
const forgeLocation = "1:5"
const gudgeonLocation = "4:2"

var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

type EasyClient struct {
	token      *string
	charClient *CharacterClient
	mapClient  *MapClient
}

func (c *EasyClient) BuyItem(characterName, itemCode string, amount, price int) (*ItemExchangeData, Error) {
	_, err := c.MoveToExchange(characterName)
	if err != nil {
		return nil, err
	}

	data, err := c.charClient.BuyItem(characterName, itemCode, amount, price)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(data.Cooldown.RemainingSeconds) * time.Second)
	return data, nil
}

func (c *EasyClient) SellItem(characterName, itemCode string, amount, price int) (*ItemExchangeData, Error) {
	_, err := c.MoveToExchange(characterName)
	if err != nil {
		return nil, err
	}

	data, err := c.charClient.BuyItem(characterName, itemCode, amount, price)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(data.Cooldown.RemainingSeconds) * time.Second)
	return data, nil
}

func (c *EasyClient) Cook(characterName, itemCode string, amount int) (*CraftData, Error) {
	_, err := c.MoveToCookingStation(characterName)
	if err != nil {
		return nil, err
	}

	return c.Craft(characterName, itemCode, amount)
}

func (c *EasyClient) CraftWeapon(characterName, itemCode string, amount int) (*CraftData, Error) {
	_, err := c.MoveToWeaponCraftingStation(characterName)
	if err != nil {
		return nil, err
	}

	return c.Craft(characterName, itemCode, amount)
}

func (c *EasyClient) CraftGear(characterName, itemCode string, amount int) (*CraftData, Error) {
	_, err := c.MoveToGearCraftingStation(characterName)
	if err != nil {
		return nil, err
	}

	return c.Craft(characterName, itemCode, amount)
}

func (c *EasyClient) CraftJewelry(characterName, itemCode string, amount int) (*CraftData, Error) {
	_, err := c.MoveToJewelryCraftingStation(characterName)
	if err != nil {
		return nil, err
	}

	return c.Craft(characterName, itemCode, amount)
}

func (c *EasyClient) Craft(characterName, itemCode string, amount int) (*CraftData, Error) {
	craftData, err := c.charClient.Craft(characterName, itemCode, amount)
	time.Sleep(time.Duration(craftData.Cooldown.RemainingSeconds) * time.Second)

	return craftData, err
}

func (c *EasyClient) DepositIntoBank(characterName, itemCode string, amount int) (*DepositData, Error) {
	_, err := c.MoveToBank(characterName)
	if err != nil {
		return nil, err
	}

	depositData, err := c.charClient.DepositItem(characterName, itemCode, amount)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(depositData.Cooldown.RemainingSeconds) * time.Second)
	return depositData, nil
}

func (c *EasyClient) WithdrawFromBank(characterName, itemCode string, amount int) (*WithdrawData, Error) {
	_, err := c.MoveToBank(characterName)
	if err != nil {
		return nil, err
	}

	withdrawData, err := c.charClient.WithdrawItem(characterName, itemCode, amount)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(withdrawData.Cooldown.RemainingSeconds) * time.Second)
	return withdrawData, nil
}

func (c *EasyClient) MoveToWeaponCraftingStation(characterName string) (*MoveData, Error) {
	return c.MoveToCoOrds(characterName, weaponCraftingLocation)
}

func (c *EasyClient) MoveToGearCraftingStation(characterName string) (*MoveData, Error) {
	return c.MoveToCoOrds(characterName, gearCraftingLocation)
}

func (c *EasyClient) MoveToJewelryCraftingStation(characterName string) (*MoveData, Error) {
	return c.MoveToCoOrds(characterName, jewelryLocation)
}

func (c *EasyClient) MoveToExchange(characterName string) (*MoveData, Error) {
	return c.MoveToCoOrds(characterName, geLocation)
}

func (c *EasyClient) MoveToChickens(characterName string) (*MoveData, Error) {
	return c.MoveToCoOrds(characterName, chickens)
}

func (c *EasyClient) MoveToCookingStation(characterName string) (*MoveData, Error) {
	return c.MoveToCoOrds(characterName, cookingLocation)
}

func (c *EasyClient) MoveToBank(characterName string) (*MoveData, Error) {
	return c.MoveToCoOrds(characterName, bankLocation)
}

func (c *EasyClient) MoveToForge(characterName string) (*MoveData, Error) {
	return c.MoveToCoOrds(characterName, forgeLocation)
}

func (c *EasyClient) MineCopper(characterName string) (*GatherData, Error) {

	miningLogger := logger.With().Str("character", characterName).Logger()

	_, err := c.MoveToCoOrds(characterName, "2:0")
	if err != nil {
		return nil, err
	}

	gatherData, err := c.charClient.Gather(characterName)
	if err != nil {
		return nil, err
	}

	coolDown := gatherData.Cooldown.RemainingSeconds
	miningLogger.Info().Msgf("Mining done. sleeping for %d seconds\n", coolDown)
	time.Sleep(time.Duration(coolDown) * time.Second)

	return gatherData, nil
}

func (c *EasyClient) MineIron(characterName string) (*GatherData, Error) {
	miningLogger := logger.With().Str("character", characterName).Logger()
	_, err := c.MoveToCoOrds(characterName, "1:7")
	if err != nil {
		return nil, err
	}

	gatherData, err := c.charClient.Gather(characterName)
	if err != nil {
		return nil, err
	}

	coolDown := gatherData.Cooldown.RemainingSeconds
	miningLogger.Info().Msgf("Mining done. sleeping for %d seconds\n", coolDown)
	time.Sleep(time.Duration(coolDown) * time.Second)

	return gatherData, nil
}

func (c *EasyClient) MineCoal(characterName string) (*GatherData, Error) {
	miningLogger := logger.With().Str("character", characterName).Logger()
	_, err := c.MoveToCoOrds(characterName, "1:6")
	if err != nil {
		return nil, err
	}

	gatherData, err := c.charClient.Gather(characterName)
	if err != nil {
		return nil, err
	}

	coolDown := gatherData.Cooldown.RemainingSeconds
	miningLogger.Info().Msgf("Mining done. sleeping for %d seconds\n", coolDown)
	time.Sleep(time.Duration(coolDown) * time.Second)
	return gatherData, nil
}

func (c *EasyClient) MineGold(characterName string) (*GatherData, Error) {
	miningLogger := logger.With().Str("character", characterName).Logger()
	_, err := c.MoveToCoOrds(characterName, "10:-4")
	if err != nil {
		return nil, err
	}

	gatherData, err := c.charClient.Gather(characterName)
	if err != nil {
		return nil, err
	}

	coolDown := gatherData.Cooldown.RemainingSeconds
	miningLogger.Info().Msgf("Mining done. sleeping for %d seconds\n", coolDown)
	time.Sleep(time.Duration(coolDown) * time.Second)
	return gatherData, nil
}

func (c *EasyClient) FishGudgeon(characterName string) (*GatherData, Error) {
	_, err := c.MoveToCoOrds(characterName, gudgeonLocation)
	if err != nil {
		return nil, err
	}

	gatherData, err := c.charClient.Gather(characterName)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(gatherData.Cooldown.RemainingSeconds) * time.Second)

	return gatherData, nil
}

// TODO: this is broken
//func (c *EasyClient) MoveToClosetLocation(characterName, resource string) (*MoveData, Error) {
//
//	moveLogger := logger.With().Str("character", characterName).Str("resource", resource).Logger()
//
//	characterData, err := c.charClient.GetCharacterInfo(characterName)
//	if err != nil {
//		return nil, err
//	}
//
//	mapTiles, err := c.mapClient.GetMapDataForResource(&resource, 0)
//	if err != nil {
//		return nil, err
//	}
//
//	startingPoint := internal.Point{
//		X: float64(characterData.X),
//		Y: float64(characterData.Y),
//	}
//	targets := buildTargetPoints(mapTiles)
//	closestPoint := internal.ClosestPoint(&startingPoint, *targets)
//
//	coordString := fmt.Sprintf("%d:%d", closestPoint.X, closestPoint.Y)
//	moveLogger.Info().Msgf("Found cloest location to resource at %s", coordString)
//
//	moveData, err := c.MoveToCoOrds(characterName, coordString)
//	if err != nil {
//		return nil, err
//	}
//
//	return moveData, nil
//}
//
//func buildTargetPoints(mapTiles *[]MapTileData) *[]internal.Point {
//
//	targets := make([]internal.Point, len(*mapTiles))
//	for _, tile := range *mapTiles {
//		targets = append(targets, internal.Point{
//			X: float64(tile.X),
//			Y: float64(tile.Y),
//		})
//	}
//
//	return &targets
//}

// MoveToCoOrds will move your character to the supplied coord string ("X:Y") and automatically handle the cooldown
// period. In the event that a character is already at its supplied location, it will return nothing.
func (c *EasyClient) MoveToCoOrds(characterName, coords string) (*MoveData, Error) {
	moveLogger := logger.With().Str("character", characterName).Str("coords", coords).Logger()

	x, y := getCoords(coords)
	moveResp, err := c.charClient.Move(characterName, x, y)
	if err != nil {
		if _, ok := err.(*CharacterAlreadyAtDestinationException); ok {
			// err is of type *CharacterAlreadyAtDestinationException
			moveLogger.Debug().Msg("Character is already at location")
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

// MoveTo will move your character to the supplied X Y spot on the map and automatically handle the cooldown period.
// In the event that a character is already at its supplied location, it will return nothing
func (c *EasyClient) MoveTo(characterName string, x, y int) (*MoveData, Error) {
	coords := fmt.Sprintf("%d:%d", x, y)
	return c.MoveToCoOrds(characterName, coords)
}

func getCoords(coord string) (x, y int) {

	parts := strings.Split(coord, ":")
	x, _ = strconv.Atoi(parts[0])
	y, _ = strconv.Atoi(parts[1])
	return x, y
}
