package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type getItemResponse struct {
	Data ExchangeItem `json:"data"`
}

type ExchangeItem struct {
	Code      string `json:"code"`
	Stock     int    `json:"stock"`
	SellPrice int    `json:"sell_price"`
	BuyPrice  int    `json:"buy_price"`
	Quantity  int    `json:"max_quantity"`
}

// GetItemExchangeData returns the Grand Exchange data for the supplied itemCode
func GetItemExchangeData(itemCode string) (*ExchangeItem, error) {

	url := fmt.Sprintf("https://api.artifactsmmo.com/ge/code&code=%s", itemCode)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response getItemResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
