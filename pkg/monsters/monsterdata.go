package monsters

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Monster struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Level       int    `json:"level"`
	HP          int    `json:"hp"`
	AttackFire  int    `json:"attack_fire"`
	AttackEarth int    `json:"attack_earth"`
	AttackWater int    `json:"attack_water"`
	AttackAir   int    `json:"attack_air"`
	ResFire     int    `json:"res_fire"`
	ResEarth    int    `json:"res_earth"`
	ResWater    int    `json:"res_water"`
	ResAir      int    `json:"res_air"`
	MinGold     int    `json:"min_gold"`
	MaxGold     int    `json:"max_gold"`
	Drops       []Drop `json:"drops"`
}

type Drop struct {
	Code        string  `json:"code"`
	Rate        float64 `json:"rate"`
	MinQuantity int     `json:"min_quantity"`
	MaxQuantity int     `json:"max_quantity"`
}

type monsterResponse struct {
	Data  []Monster `json:"data"`
	Total int       `json:"total"`
	Page  int       `json:"page"`
	Size  int       `json:"size"`
	Pages int       `json:"pages"`
}

func GetAllMonsterData() ([]Monster, error) {
	monsters := make([]Monster, 0)
	page := 1
	monsterData, err := requestMonsterData(page)
	if err != nil {
		return nil, err
	}

	monsters = append(monsters, monsterData.Data...)

	for page <= monsterData.Pages {
		page++
		monsterData, err = requestMonsterData(page)
		if err != nil {
			return nil, err
		}
		monsters = append(monsters, monsterData.Data...)
	}

	return monsters, nil
}

func requestMonsterData(page int) (*monsterResponse, error) {

	url := fmt.Sprintf("https://api.artifactsmmo.com/monsters?page=%d", page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response monsterResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
