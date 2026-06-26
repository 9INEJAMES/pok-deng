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
		ID:         uuid.NewString(),
		Balance:    input.InitialBalance,
		State:      models.WaitingForCut,
		Deck:       models.NewDeck(),
		PlayerHand: []models.Card{},
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

	case models.BET:
		return s.bet(game, input.Amount)

	case models.DRAW:
		return s.draw(game)

	case models.STAY:
		return s.stay(game)

	case models.NEXT_ROUND:
		return s.nextRound(game)

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

func (s *gameService) bet(game *models.Game, amount int) (*models.Game, error) {
	if game.State != models.WaitingForBet {
		return nil, appErr.ErrInvalidState
	}

	if amount <= 0 {
		return nil, appErr.ErrInvalidAmount
	}

	if amount > game.Balance {
		return nil, appErr.ErrInsufficientFund
	}

	game.Balance -= amount
	game.Bet = amount

	if err := drawCard(game, &game.PlayerHand); err != nil {
		return nil, err
	}
	if err := drawCard(game, &game.PlayerHand); err != nil {
		return nil, err
	}

	if err := drawCard(game, &game.DealerHand); err != nil {
		return nil, err
	}
	if err := drawCard(game, &game.DealerHand); err != nil {
		return nil, err
	}

	playerScore := calculateScore(game.PlayerHand)
	dealerScore := calculateScore(game.DealerHand)
	game.PlayerScore = playerScore

	if playerScore >= 8 || dealerScore >= 8 {
		return s.finishRound(game)
	}

	game.State = models.WaitingForDecision
	return game, nil
}

func calculateScore(cards []models.Card) int {
	sum := 0
	for _, c := range cards {
		sum += c.Value
	}

	return sum % 10
}

func drawCard(game *models.Game, hand *[]models.Card) error {
	card, err := game.Deck.Draw()
	if err != nil {
		return appErr.ErrInvalidAmount
	}
	*hand = append(*hand, card)

	return nil
}

func (s *gameService) draw(game *models.Game) (*models.Game, error) {
	if game.State != models.WaitingForDecision {
		return nil, appErr.ErrInvalidState
	}
	if err := drawCard(game, &game.PlayerHand); err != nil {
		return nil, err
	}
	game.PlayerScore = calculateScore(game.PlayerHand)

	return s.stay(game)
}

func (s *gameService) stay(game *models.Game) (*models.Game, error) {
	if game.State != models.WaitingForDecision {
		return nil, appErr.ErrInvalidState
	}

	return s.dealerTurn(game)
}

func (s *gameService) dealerTurn(game *models.Game) (*models.Game, error) {
	score := calculateScore(game.DealerHand)
	if score < 4 {
		card, err := game.Deck.Draw()
		if err == nil {
			game.DealerHand = append(game.DealerHand, card)
		}
	}

	return s.finishRound(game)
}

func (s *gameService) finishRound(game *models.Game) (*models.Game, error) {

	playerScore := calculateScore(game.PlayerHand)
	dealerScore := calculateScore(game.DealerHand)

	game.State = models.RoundEnd
	game.PlayerScore = playerScore
	game.DealerScore = dealerScore

	if playerScore > dealerScore {
		game.Balance += game.Bet * 2
		game.Winner = "Player"
		return game, nil
	}

	if dealerScore > playerScore {
		game.Winner = "Dealer"
		return game, nil
	}

	game.Balance += game.Bet
	game.Winner = "Tie"

	return game, nil
}

func (s *gameService) nextRound(game *models.Game) (*models.Game, error) {

	if game.State != models.RoundEnd {
		return nil, appErr.ErrInvalidState
	}

	game.PlayerHand = nil
	game.DealerHand = nil
	game.Bet = 0

	game.State = models.WaitingForCut

	game.Deck = models.NewDeck()
	game.Deck.Shuffle()

	return game, nil
}
