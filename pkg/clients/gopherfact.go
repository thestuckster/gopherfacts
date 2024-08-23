package clients

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
	return &GopherFactClient{
		token:           token,
		CharacterClient: characterClient,
		AccountClient:   &MyAccountClient{token},
		EasyClient:      &EasyClient{token, characterClient},
	}
}
