package models

type Card struct {
	Suit  Suit `json:"suit"`
	Rank  Rank `json:"rank"`
	Value int  `json:"value"`
}

var rankValues = map[Rank]int{
	Ace:   1,
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   0,
	Jack:  0,
	Queen: 0,
	King:  0,
}

func RankValue(rank Rank) int {
	return rankValues[rank]
}