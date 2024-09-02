package clients

import (
	"encoding/json"
	"fmt"
	"github.com/thestuckster/gopherfacts/internal"
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

type GetBankItemsResponse struct {
	Data  []Item `json:"data"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Pages int    `json:"pages"`
}

func (c *MyAccountClient) GetBankItems(itemCode *string, page, size int) (*GetBankItemsResponse, Error) {
	//TODO:
	return nil, nil
}

type bankExpansionResponse struct {
	Data BankExpansionData `json:"data"`
}

type BankExpansionData struct {
	Cooldown Cooldown `json:"cooldown"`
}

func (c *CharacterClient) BuyBankExpansion() (*BankExpansionData, Error) {
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
