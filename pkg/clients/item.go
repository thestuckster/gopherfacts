package clients

import (
	"fmt"
	"github.com/thestuckster/gopherfacts/internal"
	"net/http"
)

type ItemClient struct {
	token *string
}

func (c *ItemClient) getItemData(itemName, skill string) Error {

	url := fmt.Sprintf(ITEM_INFO+"?name=%s&craft_skill=%s", itemName, skill)
	req := internal.BuildGetRequest(url, *c.token)
	resp, respBody := internal.MakeHttpRequest(req, false)
	err := c.buildError(resp)

	if err != nil {
		return err
	}

}

func (c *ItemClient) buildError(resp *http.Response) Error {
	switch resp.StatusCode {
	case 200:
		return nil
	case 403:
		return NewForbiddenException()
	case 404:
		return NewMapNotFoundException()
	case 422:
		return NewUnprocessableEntityException()
	case 478:
		return NewNotEnoughResourcesException()
	case 486:
		return NewActionAlreadyInProgressException()
	case 490:
		return NewCharacterAlreadyAtDestinationException()
	case 493:
		return NewSkillLevelToLow()
	case 497:
		return NewCharacterInventoryFullException()
	case 498:
		return NewCharacterNotFoundException()
	case 499:
		return NewInCoolDownException()
	case 598:
		return NewResourceNotFoundException()
	default:
		return nil
	}
}
