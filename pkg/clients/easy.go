package clients

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/thestuckster/gopherfacts/internal"
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
const miningLocation = "1:5"

var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

type EasyClient struct {
	token      *string
	charClient *CharacterClient
	mapClient  *MapClient
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

func (c *EasyClient) MineCopper(characterName string) (*GatherData, Error) {

	miningLogger := logger.With().Str("character", characterName).Logger()

	_, err := c.moveToLocation(characterName, "2:0")
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

// TODO: this is broken
func (c *EasyClient) MoveToClosetLocation(characterName, resource string) (*MoveData, Error) {

	moveLogger := logger.With().Str("character", characterName).Str("resource", resource).Logger()

	characterData, err := c.charClient.GetCharacterInfo(characterName)
	if err != nil {
		return nil, err
	}

	mapTiles, err := c.mapClient.GetMapDataForResource(&resource, 0)
	if err != nil {
		return nil, err
	}

	startingPoint := internal.Point{
		X: float64(characterData.X),
		Y: float64(characterData.Y),
	}
	targets := buildTargetPoints(mapTiles)
	closestPoint := internal.ClosestPoint(&startingPoint, *targets)

	coordString := fmt.Sprintf("%d:%d", closestPoint.X, closestPoint.Y)
	moveLogger.Info().Msgf("Found cloest location to resource at %s", coordString)

	moveData, err := c.moveToLocation(characterName, coordString)
	if err != nil {
		return nil, err
	}

	return moveData, nil
}

func buildTargetPoints(mapTiles *[]MapTileData) *[]internal.Point {

	targets := make([]internal.Point, len(*mapTiles))
	for _, tile := range *mapTiles {
		targets = append(targets, internal.Point{
			X: float64(tile.X),
			Y: float64(tile.Y),
		})
	}

	return &targets
}

func (c *EasyClient) moveToLocation(characterName, coords string) (*MoveData, Error) {
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

func getCoords(coord string) (x, y int) {

	parts := strings.Split(coord, ":")
	x, _ = strconv.Atoi(parts[0])
	y, _ = strconv.Atoi(parts[1])
	return x, y
}
