package models

type Suit string

const (
	Clubs    Suit = "Clubs"
	Diamonds Suit = "Diamonds"
	Hearts   Suit = "Hearts"
	Spades   Suit = "Spades"
)

type Rank string

const (
	Ace   Rank = "A"
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "J"
	Queen Rank = "Q"
	King  Rank = "K"
)

var Suits = []Suit{Clubs, Diamonds, Hearts, Spades}
var Ranks = []Rank{
	Ace, Two, Three, Four, Five, Six, Seven,
	Eight, Nine, Ten, Jack, Queen, King,
}

type Action string

const (
	CUT        Action = "cut"
	BET        Action = "bet"
	DRAW       Action = "draw"
	STAY       Action = "stay"
	NEXT_ROUND Action = "next_round"
)

