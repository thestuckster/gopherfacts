package clients

type error interface {
	Error() string
}

type MapNotFoundException struct {
	Message string
}

func NewMapNotFoundException() *MapNotFoundException {
	return &MapNotFoundException{
		Message: "Map not found",
	}
}

func (e *MapNotFoundException) Error() string {
	return e.Message
}

type ActionAlreadyInProgressException struct {
	Message string
}

func NewActionAlreadyInProgressException() *ActionAlreadyInProgressException {
	return &ActionAlreadyInProgressException{
		Message: "Action already in progress",
	}
}

func (e *ActionAlreadyInProgressException) Error() string {
	return e.Message
}

type CharacterAlreadyAtDestinationException struct {
	Message string
}

func NewCharacterAlreadyAtDestinationException() *CharacterAlreadyAtDestinationException {
	return &CharacterAlreadyAtDestinationException{
		Message: "Character already at destination",
	}
}

func (e *CharacterAlreadyAtDestinationException) Error() string {
	return e.Message
}

type CharacterNotFoundException struct {
	Message string
}

func NewCharacterNotFoundException() *CharacterNotFoundException {
	return &CharacterNotFoundException{
		Message: "Character not found",
	}
}

func (e *CharacterNotFoundException) Error() string {
	return e.Message
}

type InCoolDownException struct {
	Message string
}

func NewInCoolDownException() *InCoolDownException {
	return &InCoolDownException{
		Message: "In cool down",
	}
}

func (e *InCoolDownException) Error() string {
	return e.Message
}

type CharacterInventoryFullException struct {
	Message string
}

func NewCharacterInventoryFullException() *CharacterInventoryFullException {
	return &CharacterInventoryFullException{
		Message: "Character inventory full",
	}
}

func (e *CharacterInventoryFullException) Error() string {
	return e.Message
}

type MonsterNotFoundException struct {
	Message string
}

func NewMonsterNotFoundException() *MonsterNotFoundException {
	return &MonsterNotFoundException{
		Message: "Monster not found",
	}
}

func (e *MonsterNotFoundException) Error() string {
	return e.Message
}

type ForbiddenException struct {
	Message string
	Details string
}

func NewForbiddenException() *ForbiddenException {
	return &ForbiddenException{
		Message: "HTTP 403 Forbidden",
	}
}

func (e *ForbiddenException) Error() string {
	return e.Message
}

type UnprocessableEntityException struct {
	Message string
}

func NewUnprocessableEntityException() *UnprocessableEntityException {
	return &UnprocessableEntityException{
		Message: "HTTP 422 Unprocessable Entity",
	}
}

func (e *UnprocessableEntityException) Error() string {
	return e.Message
}
