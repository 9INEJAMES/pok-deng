package models

type GameResponse struct {
	GameID            string  `json:"game_id"`
	State             string  `json:"state"`
	Balance           int     `json:"balance"`
	PlayerHand        []Card  `json:"player_hand"`
	DealerHandVisible []Card  `json:"dealer_hand_visible"`
	PlayerScore       *int    `json:"player_score"`
	DealerScore       *int    `json:"dealer_score"`
	Winner            *string `json:"winner"`
}
