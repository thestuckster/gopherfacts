package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopherfacts/internal"
	"net/http"
)

type MyCharacterClient struct {
	token *string
}

type moveResponse struct {
	Data MoveData `json:"data"`
}

type MoveData struct {
	Cooldown  int             `json:"cooldown"`
	Content   MapContent      `json:"content"`
	Character CharacterSchema `json:"character"`
}

func (c *MyCharacterClient) Move(characterName string, x, y int) (*MoveData, Error) {

	body := make(map[string]int)
	body["x"] = x
	body["y"] = y

	url := fmt.Sprintf(MOVE, characterName)
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req := internal.BuildPostRequest(url, *c.token, bytes.NewReader(jsonData))
	resp, respBody := internal.MakeHttpRequest(req, false)
	err = c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var response moveResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

type fightResponse struct {
	Data FightData `json:"data"`
}

type FightData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Fight     Fight           `json:"fight"`
	Character CharacterSchema `json:"character"`
}

func (c *MyCharacterClient) Fight(characterName string) (*FightData, Error) {
	url := fmt.Sprintf(FIGHT, characterName)
	req := internal.BuildPostRequestNoBody(url, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)

	err := c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var data fightResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type gatherResponse struct {
	Data GatherData `json:"data"`
}

type GatherData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Details   CraftDetails    `json:"details"`
	Character CharacterSchema `json:"character"`
}

func (c *MyCharacterClient) Gather(characterName string) (*GatherData, Error) {
	url := fmt.Sprintf(GATHER, characterName)
	req := internal.BuildPostRequestNoBody(url, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)

	err := c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var data gatherResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type craftResponse struct {
	Data CraftData `json:"data"`
}

type craftRequest struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type CraftData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Details   CraftDetails    `json:"details"`
	Character CharacterSchema `json:"character"`
}

func (c *MyCharacterClient) Craft(characterName, itemCode string, amount int) (*CraftData, Error) {
	url := fmt.Sprintf(CRAFT, characterName)
	body := craftRequest{
		Quantity: amount,
		Code:     itemCode,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req := internal.BuildPostRequest(url, *c.token, bytes.NewReader(jsonData))
	resp, respBody := internal.MakeHttpRequest(req, false)
	err = c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var data craftResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

func (c *MyCharacterClient) buildError(resp *http.Response) Error {
	switch resp.StatusCode {
	case 200:
		return nil
	case 403:
		return NewForbiddenException()
	case 404:
		return NewMapNotFoundException()
	case 422:
		return NewUnprocessableEntityException()
	case 478:
		return NewNotEnoughResourcesException()
	case 486:
		return NewActionAlreadyInProgressException()
	case 490:
		return NewCharacterAlreadyAtDestinationException()
	case 493:
		return NewSkillLevelToLow()
	case 497:
		return NewCharacterInventoryFullException()
	case 498:
		return NewCharacterNotFoundException()
	case 499:
		return NewInCoolDownException()
	case 598:
		return NewResourceNotFoundException()
	default:
		return nil
	}
}
