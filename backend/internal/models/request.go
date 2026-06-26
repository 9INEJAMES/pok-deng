package models

type GameStart struct {
	InitialBalance int `json:"initial_balance" validate:"required,gt=0"`
}

type GameAction struct {
	Action Action `json:"action" validate:"required,oneof=cut bet draw stay next_round"`
	Amount int    `json:"amount,omitempty"`
}
