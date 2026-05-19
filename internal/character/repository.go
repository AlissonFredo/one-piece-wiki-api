package character

import (
	"context"
	"sync"
)

type CharacterRepository interface {
	CreateCharacter(context context.Context, character Character) (Character, error)
}

type InMemoryCharacterRepository struct {
	mutex      sync.RWMutex
	characters map[string]Character
}

func NewInMemoryRepository() *InMemoryCharacterRepository {
	return &InMemoryCharacterRepository{
		characters: make(map[string]Character),
	}
}

func (repository *InMemoryCharacterRepository) CreateCharacter(context context.Context, character Character) (Character, error) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()

	repository.characters[character.ID] = character
	return character, nil
}
