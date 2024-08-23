package clients

type GopherFactClient struct {
	token           *string
	CharacterClient *CharacterClient
	AccountClient   *MyAccountClient
}

func NewClient(token *string) *GopherFactClient {
	return &GopherFactClient{
		token:           token,
		CharacterClient: &CharacterClient{token},
		AccountClient:   &MyAccountClient{token},
	}
}
