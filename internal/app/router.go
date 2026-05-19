package app

import (
	"github.com/AlissonFredo/one-piece-wiki-api/internal/character"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	characterRepoository := character.NewInMemoryRepository()
	characterService := character.NewService(characterRepoository)
	characterHandler := character.NewHandler(characterService)

	r.POST("/characters", characterHandler.CreateCharacter)
}
