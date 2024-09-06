package items

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ItemMetaData struct {
	Name        string   `json:"name"`
	Code        string   `json:"code"`
	Level       int      `json:"level"`
	Type        string   `json:"type"`
	Subtype     string   `json:"subtype"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
	Craft       Craft    `json:"craft"`
}

type Effect struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Craft struct {
	Skill    string      `json:"skill"`
	Level    int         `json:"level"`
	Items    []CraftItem `json:"items"`
	Quantity int         `json:"quantity"`
}

type CraftItem struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type itemDataResponse struct {
	Data  []ItemMetaData `json:"data"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Pages int            `json:"pages"`
}

func GetAllItemData() ([]ItemMetaData, error) {
	items := make([]ItemMetaData, 0)

	page := 1
	itemData, err := requestItems(page)
	if err != nil {
		return nil, err
	}

	items = append(items, itemData.Data...)

	for page < itemData.Pages {
		page++
		itemData, err := requestItems(page)
		if err != nil {
			return nil, err
		}

		items = append(items, itemData.Data...)
	}

	return items, nil
}

func requestItems(page int) (*itemDataResponse, error) {
	url := fmt.Sprintf("https://api.artifactsmmo.com/items?page=%d", page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response itemDataResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
