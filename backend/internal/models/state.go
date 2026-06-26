package models

type GameState string

const (
	WaitingForCut      GameState = "WAITING_FOR_CUT"
	WaitingForBet      GameState = "WAITING_FOR_BET"
	WaitingForDecision GameState = "WAITING_FOR_DECISION"
	RoundEnd           GameState = "ROUND_END"
)

type Winner string

const (
	Player Winner = "Player"
	Dealer Winner = "Dealer"
	Tie    Winner = "Tie"
)