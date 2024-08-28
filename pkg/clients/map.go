package clients

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/thestuckster/gopherfacts/internal"
	"os"
	"strconv"
)

type MapClient struct {
}

type mapResponse struct {
	Data []MapTileData `json:"data"`
}

type MapTileData struct {
	Name    string         `json:"name"`
	Skin    string         `json:"skin"`
	X       int            `json:"x"`
	Y       int            `json:"y"`
	Content MapTileContent `json:"content"`
	Total   int            `json:"total"`
	Page    int            `json:"page"`
	Size    int            `json:"size"`
	Pages   int            `json:"pages"`
}

type MapTileContent struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

func GetMapDataForResource(resource *string, page int) (*[]MapTileData, Error) {

	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Str("resource", *resource).Int("page", page).Logger()

	//since Go doesn't have default param values
	if page == 0 {
		page = 1
	}

	url := ALL_MAPS + "?page=" + strconv.Itoa(page)
	if resource != nil {
		url += fmt.Sprintf("&resource=%s", *resource)
	}

	req := internal.BuildGetRequest(url, "")
	resp, respBody := internal.MakeHttpRequest(req, false)
	if resp.StatusCode != 200 {
		logger.Error().Msg("Error getting map data")
		return nil, NewCatchAllException("Error getting map data")
	}

	var data mapResponse
	err := json.Unmarshal(respBody, &data)
	if err != nil {
		logger.Error().Msg("Error unmarshalling map data")
		return nil, err
	}

	return &data.Data, nil
}
