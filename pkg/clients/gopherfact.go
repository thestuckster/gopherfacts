package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thestuckster/gopherfacts/internal"
	"net/http"
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
	resp, body := internal.MakeHttpRequest(req, false)
	if resp.StatusCode != 200 {
		return nil, NewCatchAllException(fmt.Sprintf("Error fetching server status http %d", resp.StatusCode))
	}

	var serverStatus serverStatusResponse
	err := json.Unmarshal(body, &serverStatus)
	if err != nil {
		return nil, err
	}

	return &serverStatus.Data, nil
}

// CreateAccount Returns &true if account creation was successful. Returns &false with either a UsernameAlreadyInUseException or
// EmailAlreadyInUseException otherwise.
func (c *GopherFactClient) CreateAccount(username, password, email string) (*bool, Error) {
	url := BASE_URL + "/accounts/create"

	body := make(map[string]string)
	body["username"] = username
	body["password"] = password
	body["email"] = email

	f := false
	jsonString, err := json.Marshal(body)
	if err != nil {
		return &f, err
	}

	req := internal.BuildPostRequest(url, "", bytes.NewReader(jsonString))
	resp, _ := internal.MakeHttpRequest(req, false)
	err = c.buildError(resp)
	if err != nil {
		return &f, err
	}

	t := true
	return &t, nil
}

func (c *GopherFactClient) buildError(resp *http.Response) Error {
	switch resp.StatusCode {
	case 200:
		return nil
	case 456:
		return NewUsernameAlreadyUsedException()
	case 457:
		return NewEmailAlreadyUsedException()
	default:
		return nil
	}
}
