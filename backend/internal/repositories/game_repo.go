package repositories

import (
	"sync"

	"github.com/9inejames/pok-deng/backend/internal/models"
)

type GameRepository struct {
	games map[string]*models.Game
	mu    sync.RWMutex
}

func NewGameRepository() *GameRepository {
	return &GameRepository{
		games: make(map[string]*models.Game),
	}
}

func (r *GameRepository) Create(game *models.Game) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.games[game.ID] = game
}

func (r *GameRepository) Get(id string) (*models.Game, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	game, ok := r.games[id]
	return game, ok
}
