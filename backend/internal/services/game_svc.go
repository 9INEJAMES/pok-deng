package services

import (
	appErr "github.com/9inejames/pok-deng/backend/internal/errors"
	"github.com/9inejames/pok-deng/backend/internal/models"
	repo "github.com/9inejames/pok-deng/backend/internal/repositories"

	"github.com/google/uuid"
)

type GameService interface {
	Start(input *models.GameStart) (*models.Game, error)
	Action(id string, input *models.GameAction) (*models.Game, error)
}

type gameService struct {
	gamerepo *repo.GameRepository
}

func NewGameService(gamerepo *repo.GameRepository) *gameService {
	return &gameService{gamerepo: gamerepo}
}

func (s *gameService) Start(input *models.GameStart) (*models.Game, error) {

	game := &models.Game{
		ID:      uuid.NewString(),
		Balance: input.InitialBalance,
		State:   models.WaitingForCut,
		Deck:    models.NewDeck(),
	}

	game.Deck.Shuffle()

	s.gamerepo.Create(game)

	return game, nil
}

func (s *gameService) Action(id string, input *models.GameAction) (*models.Game, error) {
	game, ok := s.gamerepo.Get(id)
	if !ok {
		return nil, appErr.ErrSessionNotFound
	}

	switch input.Action {
	case models.CUT:
		return s.cut(game, input.Amount)

	// case models.BET:
	// 	return s.bet(game, input.Amount)

	// case models.DRAW:
	// 	return s.draw(game)

	// case models.STAY:
	// 	return s.stay(game)

	// case models.NEXTROUND:
	// 	return s.nextRound(game)

	default:
		return nil, appErr.ErrInvalidAction
	}

}

func (s *gameService) cut(game *models.Game, amount int) (*models.Game, error) {

	if game.State != models.WaitingForCut {
		return nil, appErr.ErrInvalidState
	}

	if err := game.Deck.Cut(amount); err != nil {
		return nil, appErr.ErrInvalidAmount
	}

	game.State = models.WaitingForBet

	return game, nil
}
