package clients

import (
	"encoding/json"
	"fmt"
	"github.com/thestuckster/gopherfacts/internal"
	"io"
	"net/http"
)

type MyAccountClient struct {
	token *string
}

type GetBankGoldResponse struct {
	Data BankGold `json:"data"`
}

func (c *MyAccountClient) GetBankGold() (*BankGold, Error) {
	req := internal.BuildGetRequest(BANK_GOLD, *c.token)
	resp, body := internal.MakeHttpRequest(req, false)

	if resp.StatusCode != 200 {
		//error thrown here
	}

	var data GetBankGoldResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err //TODO: custom error
	}

	return &data.Data, nil
}

type getBankItemResponse struct {
	Data  []Item `json:"data"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Pages int    `json:"pages"`
}

func (c *MyAccountClient) GetAllBankItems() ([]Item, Error) {

	items := make([]Item, 0)
	page := 1
	itemData, err := c.requestBankItems(page)
	if err != nil {
		return nil, err
	}

	items = append(items, itemData.Data...)
	for page <= itemData.Pages {
		page++
		itemData, err = c.requestBankItems(page)
		if err != nil {
			return nil, err
		}

		items = append(items, itemData.Data...)
	}

	return items, nil
}

func (c *MyAccountClient) requestBankItems(page int) (*getBankItemResponse, error) {
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/bank/items?page=%d", page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data getBankItemResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

type bankExpansionResponse struct {
	Data BankExpansionData `json:"data"`
}

type BankExpansionData struct {
	Cooldown Cooldown `json:"cooldown"`
}

func (c *MyAccountClient) BuyBankExpansion() (*BankExpansionData, Error) {
	req := internal.BuildPostRequestNoBody(BANK_EXPANSION, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)
	if resp.StatusCode != 200 {
		errorMessage := fmt.Sprintf("Error buying bank expansion, status code: %d", resp.StatusCode)
		return nil, NewCatchAllException(errorMessage)
	}

	var data bankExpansionResponse
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, err
	}
	return &data.Data, nil
}

func (c *MyAccountClient) buildError(resp *http.Response) Error {
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
	case 482:
		return NewNoItemAtThisPrice()
	case 483:
		return NewTransactionAlreadyInProgressException()
	case 486:
		return NewActionAlreadyInProgressException()
	case 487:
		return NewCharacterHasNoTask()
	case 488:
		return NewTaskNotCompletedException()
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
