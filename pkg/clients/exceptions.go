package clients

type Error interface {
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

type ResourceNotFoundException struct {
	Message string
}

func NewResourceNotFoundException() *ResourceNotFoundException {
	return &ResourceNotFoundException{
		Message: "Resource not found on this map space",
	}
}

func (e *ResourceNotFoundException) Error() string {
	return e.Message
}

type SkillLevelToLow struct {
	Message string
}

func NewSkillLevelToLow() *SkillLevelToLow {
	return &SkillLevelToLow{
		Message: "Skill level too low",
	}
}

func (e *SkillLevelToLow) Error() string {
	return e.Message
}

//GENERIC HTTP EXCEPTIONS

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
