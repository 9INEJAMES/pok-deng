package models

import (
	"errors"
	"math/rand/v2"
)

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	deck := Deck{}

	for _, suit := range Suits {
		for _, rank := range Ranks {
			deck.Cards = append(deck.Cards, Card{
				Suit:  suit,
				Rank:  rank,
				Value: RankValue(rank),
			})
		}
	}

	return deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card, nil
}

func (d *Deck) Cut(n int) error {
	if n <= 0 || n >= len(d.Cards) {
		return errors.New("invalid cut amount")
	}

	d.Cards = append(d.Cards[n:], d.Cards[:n]...)

	return nil
}
