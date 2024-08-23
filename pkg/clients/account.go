package clients

import (
	"encoding/json"
	"github.com/thestuckster/gopherfacts/internal"
)

type MyAccountClient struct {
	token *string
}

type GetBankGoldResponse struct {
	Data BankGold `json:"data"`
}

type BankGold struct {
	Quantity int
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
