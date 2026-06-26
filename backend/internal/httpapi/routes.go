package httpapi

import (
	"github.com/9inejames/pok-deng/backend/internal/handlers"
	"github.com/9inejames/pok-deng/backend/internal/repositories"
	"github.com/9inejames/pok-deng/backend/internal/services"
)

func RegisterRoutes(s *fiberServer) {
	initializeGameRoutes(s)
}

func initializeGameRoutes(s *fiberServer) {
	repo := repositories.NewGameRepository()
	svc := services.NewGameService(repo)
	h := handlers.NewGameHandler(svc)
	
	game := s.app.Group("/api/v1/game")
	game.Post("/start", h.Start)
	game.Post("/:game_id/action", h.Action)
}
