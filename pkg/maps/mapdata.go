package maps

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MapData struct {
	Name    string  `json:"name"`
	Skin    string  `json:"skin,omitempty"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	Content Content `json:"content"`
}

type Content struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type mapDataResponse struct {
	Data  []MapData `json:"data"`
	Total int       `json:"total"`
	Page  int       `json:"page"`
	Size  int       `json:"size"`
	Pages int       `json:"pages"`
}

func GetAllMapData() ([]MapData, error) {
	maps := make([]MapData, 0)

	page := 1
	mapData, err := requestMapData(page)
	if err != nil {
		return nil, err
	}

	maps = append(maps, mapData.Data...)

	for page <= mapData.Pages {
		page++
		mapData, err = requestMapData(page)
		if err != nil {
			return nil, err
		}
		maps = append(maps, mapData.Data...)
	}

	return maps, nil
}

func requestMapData(page int) (*mapDataResponse, error) {

	url := fmt.Sprintf("https://api.artifactsmmo.com/maps?page=%d", page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response mapDataResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
