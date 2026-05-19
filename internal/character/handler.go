package character

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	characterService *CharacterService
}

func NewHandler(characterService *CharacterService) *CharacterHandler {
	return &CharacterHandler{characterService: characterService}
}

type CreateCharacterRequest struct {
	Name string `json:"name"`
}

func (handler *CharacterHandler) CreateCharacter(context *gin.Context) {
	var request CreateCharacterRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	createdCharacter, err := handler.characterService.CreateCharacter(
		context.Request.Context(),
		request.Name,
	)

	if err != nil {
		handler.writeError(context, err)
		return
	}

	context.JSON(http.StatusCreated, createdCharacter)
}

func (hundler *CharacterHandler) writeError(context *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrInvalidCharacterInput):
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}
