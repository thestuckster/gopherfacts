package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thestuckster/gopherfacts/internal"
	"net/http"
)

type CharacterClient struct {
	token *string
}

type getCharactersResponse struct {
	Data []CharacterSchema `json:"data"`
}

func (c *CharacterClient) GetAllCharactersInfo() ([]CharacterSchema, Error) {
	url := fmt.Sprintf(CHARACTER, "characters")
	req := internal.BuildPostRequestNoBody(url, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)
	err := c.buildError(resp)
	if err != nil {
		return []CharacterSchema{}, err
	}

	var response getCharactersResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return []CharacterSchema{}, err
	}

	return response.Data, nil
}

type getSingleCharacterResponse struct {
	Data CharacterSchema `json:"data"`
}

func (c *CharacterClient) GetCharacterInfo(name string) (*CharacterSchema, Error) {
	url := fmt.Sprintf(CHARACTER_INFO, name)
	req := internal.BuildGetRequest(url, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)
	err := c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var response getSingleCharacterResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

type moveResponse struct {
	Data MoveData `json:"data"`
}

type MoveData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Content   MapContent      `json:"content"`
	Character CharacterSchema `json:"character"`
}

func (c *CharacterClient) Move(characterName string, x, y int) (*MoveData, Error) {

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

func (c *CharacterClient) Fight(characterName string) (*FightData, Error) {
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

func (c *CharacterClient) Gather(characterName string) (*GatherData, Error) {
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

func (c *CharacterClient) Craft(characterName, itemCode string, amount int) (*CraftData, Error) {
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

type bankRequest struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type depositResponse struct {
	Data DepositData `json:"data"`
}

type DepositData struct {
	Cooldown      Cooldown        `json:"cooldown"`
	Item          Item            `json:"item"`
	BankInventory []Item          `json:"bank"`
	Character     CharacterSchema `json:"character"`
}

func (c *CharacterClient) DepositItem(characterName, itemCode string, amount int) (*DepositData, Error) {
	url := fmt.Sprintf(DEPOSIT_CHARACTER_BANK, characterName)
	body := bankRequest{
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

	var data depositResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}
	return &data.Data, nil
}

type withdrawResponse struct {
	Data WithdrawData `json:"data"`
}

type WithdrawData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Item      Item            `json:"item"`
	Bank      []Item          `json:"bank"`
	Character CharacterSchema `json:"character"`
}

func (c *CharacterClient) WithdrawItem(characterName, itemCode string, amount int) (*WithdrawData, Error) {
	url := fmt.Sprintf(WITHDRAW_CHARACTER_BANK, characterName)
	body := bankRequest{
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

	var data withdrawResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type goldResponse struct {
	Data GoldData `json:"data"`
}

type GoldData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Bank      BankGold        `json:"bank"`
	Character CharacterSchema `json:"character"`
}

func (c *CharacterClient) DepositGold(characterName string, quantity int) (*GoldData, Error) {
	url := fmt.Sprintf(DEPOSIT_GOLD_BANK, characterName)
	body := make(map[string]int)
	body["quantity"] = quantity

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
	var data goldResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

func (c *CharacterClient) WithdrawGold(characterName string, quantity int) (*GoldData, Error) {
	url := fmt.Sprintf(WITHDRAW_GOLD_BANK, characterName)
	body := make(map[string]int)
	body["quantity"] = quantity
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

	var data goldResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type itemRequest struct {
	Quantity int    `json:"quantity"`
	Code     string `json:"code"`
	Price    int    `json:"price"`
}

type itemResponse struct {
	Data ItemExchangeData `json:"data"`
}

type ItemExchangeData struct {
	Cooldown    Cooldown        `json:"cooldown"`
	Transaction Transaction     `json:"transaction"`
	Character   CharacterSchema `json:"character"`
}

func (c *CharacterClient) SellItem(characterName, itemCode string, amount, price int) (*ItemExchangeData, Error) {
	return c.itemTransaction(characterName, itemCode, amount, price, true)
}

func (c *CharacterClient) BuyItem(characterName, itemCode string, amount, price int) (*ItemExchangeData, Error) {
	return c.itemTransaction(characterName, itemCode, amount, price, false)
}

func (c *CharacterClient) itemTransaction(characterName, itemCode string, amount, price int, sell bool) (*ItemExchangeData, Error) {
	url := ""
	if sell == true {
		url = fmt.Sprintf(GE_SELL, characterName)
	} else {
		url = fmt.Sprintf(GE_BUY, characterName)
	}

	body := itemRequest{
		Quantity: amount,
		Code:     itemCode,
		Price:    price,
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

	var data itemResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type equipRequest struct {
	Code     string `json:"code"`
	Slot     string `json:"slot"`
	Quantity int    `json:"quantity"`
}

type equipResponse struct {
	Data EquipData `json:"data"`
}

type EquipData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Slot      string          `json:"slot"`
	Item      Item            `json:"item"`
	Character CharacterSchema `json:"character"`
}

func (c *CharacterClient) EquipItem(characterName, itemCode, slot string, quantity int) (*EquipData, Error) {
	url := fmt.Sprintf(EQUIP, characterName)

	body := equipRequest{
		Quantity: quantity,
		Code:     itemCode,
		Slot:     slot,
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

	var data equipResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}
	return &data.Data, nil
}

func (c *CharacterClient) UnEquipItem(characterName, slot string, quantity int) (*EquipData, Error) {
	url := fmt.Sprintf(UNEQUIP, characterName)
	body := make(map[string]any)
	body["quantity"] = quantity
	body["slot"] = slot

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

	var data equipResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}
	return &data.Data, nil
}

type deleteItemRequest struct {
	Quantity int    `json:"quantity"`
	Code     string `json:"code"`
}
type deleteItemResponse struct {
	Data DeleteItemData `json:"data"`
}

type DeleteItemData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Item      Item            `json:"item"`
	Character CharacterSchema `json:"character"`
}

func (c *CharacterClient) DeleteItem(characterName, code string, quanity int) (*DeleteItemData, Error) {
	url := fmt.Sprintf(DELETE_ITEM, characterName)
	body := deleteItemRequest{
		Code:     code,
		Quantity: quanity,
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
	var data deleteItemResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type recycleRequest struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type recycleResponse struct {
	Data CraftData `json:"data"`
}

func (c *CharacterClient) RecycleItem(characterName, code string, quantity int) (*CraftData, Error) {
	url := fmt.Sprintf(RECYCLE, characterName)
	body := recycleRequest{
		Code:     code,
		Quantity: quantity,
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

	var data recycleResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type taskResponse struct {
	Data TaskData `json:"data"`
}

type TaskData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Task      Task            `json:"task"`
	Character CharacterSchema `json:"character"`
	Reward    Item            `json:"reward"`
}

func (c *CharacterClient) AcceptNewTask(characterName string) (*TaskData, Error) {
	url := fmt.Sprintf(NEW_TASK, characterName)
	req := internal.BuildPostRequestNoBody(url, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)

	err := c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var data taskResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type completeTaskResponse struct {
	Data TaskRewardData `json:"data"`
}

type TaskRewardData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Reward    Item            `json:"reward"`
	Character CharacterSchema `json:"character"`
}

func (c *CharacterClient) CompleteTask(characterName string) (*TaskRewardData, Error) {
	url := fmt.Sprintf(COMPLETE_TASK, characterName)
	req := internal.BuildPostRequestNoBody(url, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)
	err := c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var data completeTaskResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

type taskCancelResponse struct {
	Data TaskCancelData `json:"data"`
}

type TaskCancelData struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Character CharacterSchema `json:"character"`
}

func (c *CharacterClient) CancelTask(characterName string) (*TaskCancelData, Error) {
	url := fmt.Sprintf(CANCEL_TASK, characterName)
	req := internal.BuildPostRequestNoBody(url, *c.token)

	resp, respBody := internal.MakeHttpRequest(req, false)
	err := c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var data taskCancelResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

func (c *CharacterClient) ExchangeTaskCoins(characterName string) (*TaskRewardData, Error) {
	url := fmt.Sprintf(EXCHANGE_TASK, characterName)
	req := internal.BuildPostRequestNoBody(url, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)
	err := c.buildError(resp)
	if err != nil {
		return nil, err
	}

	var data completeTaskResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

func (c *CharacterClient) buildError(resp *http.Response) Error {
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
	//TODO: GE buy 482
	//TODO: GE buy 483
	case 486:
		return NewActionAlreadyInProgressException()
	//TODO: Task complete 487
	//TODO: task complete 488
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
