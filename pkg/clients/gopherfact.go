package clients

import (
	"encoding/json"
	"fmt"
	"github.com/thestuckster/gopherfacts/internal"
	"time"
)

type GopherFactClient struct {
	token           *string
	CharacterClient *CharacterClient
	AccountClient   *MyAccountClient
	EasyClient      *EasyClient
}

func NewClient(token *string) *GopherFactClient {

	characterClient := &CharacterClient{
		token: token,
	}

	mapClient := &MapClient{}

	return &GopherFactClient{
		token:           token,
		CharacterClient: characterClient,
		AccountClient:   &MyAccountClient{token},
		EasyClient: &EasyClient{
			token,
			characterClient,
			mapClient,
		},
	}
}

type serverStatusResponse struct {
	Data ServerStatus `json:"data"`
}

type ServerStatus struct {
	Status           string         `json:"status"`
	Version          string         `json:"version"`
	MaxLevel         int            `json:"max_level"`
	CharactersOnline int            `json:"characters_online"`
	ServerTime       time.Time      `json:"server_time"`
	Announcements    []Announcement `json:"announcements"`
	LastWipe         string         `json:"last_wipe"`
	NextWipe         string         `json:"next_wipe"`
}

type Announcement struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *GopherFactClient) CheckServerStatus() (*ServerStatus, Error) {
	req := internal.BuildGetRequest(BASE_URL, *c.token)
	resp, respbody := internal.MakeHttpRequest(req, false)
	if resp.StatusCode != 200 {
		return nil, NewCatchAllException(fmt.Sprintf("Error fetching server status http %d", resp.StatusCode))
	}

	var serverStatus serverStatusResponse
	err := json.Unmarshal(respbody, &serverStatus)
	if err != nil {
		return nil, err
	}

	return &serverStatus.Data, nil
}

func (c *GopherFactClient) GetAllMapData(contentType string) (any, Error) {
	err := validateContentType(&contentType)
	if err != nil {
		return nil, err
	}

	return nil, nil //TODO
}

func validateContentType(t *string) Error {
	if t == nil {
		return nil
	}

	switch *t {
	case "monster":
	case "workshop":
	case "resource":
	case "bank":
	case "grand_exchange":
	case "task_master":
		return nil
	default:
		return NewInvalidContentTypeException(*t)
	}

	return nil
}
