package models

type Game struct {
	ID    string    `json:"game_id"`
	State GameState `json:"state"`

	Balance int `json:"balance"`
	Bet     int `json:"bet"`

	PlayerHand []Card `json:"player_hand"`
	DealerHand []Card `json:"-"`

	PlayerScore int  `json:"player_score"`
	DealerScore int  `json:"-"`
	
	Deck        Deck `json:"-"`

	Winner Winner `json:"-"`
}
