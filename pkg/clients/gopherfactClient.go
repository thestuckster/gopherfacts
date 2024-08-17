package clients

type GopherFactClient struct {
	token             *string
	MyCharacterClient *MyCharacterClient
	MyAccountClient   *MyAccountClient
}

func NewClient(token *string) *GopherFactClient {
	return &GopherFactClient{
		token:             token,
		MyCharacterClient: &MyCharacterClient{token},
		MyAccountClient:   &MyAccountClient{token},
	}
}
