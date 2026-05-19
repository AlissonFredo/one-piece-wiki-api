package character

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidCharacterInput = errors.New("invalid character input")

type CharacterService struct {
	characterRepository CharacterRepository
}

func NewService(characterRepository CharacterRepository) *CharacterService {
	return &CharacterService{characterRepository: characterRepository}
}

func (service *CharacterService) CreateCharacter(context context.Context, name string) (Character, error) {
	if name == "" {
		return Character{}, ErrInvalidCharacterInput
	}

	newCharacter := Character{
		ID:   uuid.NewString(),
		Name: name,
	}

	return service.characterRepository.CreateCharacter(context, newCharacter)
}
