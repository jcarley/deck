// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package deck

import (
	"crypto/rand"
	"math/big"
)

const NoCard = 0xff // No card can have this value.

// A Deck represents a single deck of cards.
type Deck []Card

// New creates a new card deck.
func New() Deck {
	d := make(Deck, 0, 52)
	d.Reset()
	return d
}

// Reset resets the deck to a full set of cards.
func (d *Deck) Reset() {
	const (
		hearts   = Card(Hearts << 4)
		diamonds = Card(Diamonds << 4)
		clubs    = Card(Clubs << 4)
		spades   = Card(Spades << 4)
	)

	dd := (*d)[:0]

	for s := 0; s < 4; s++ {
		suit := s << 4
		for c := 0; c < 13; c++ {
			dd = append(dd, Card(suit|c))
		}
	}

	*d = dd
}

// Pop pops the top card from the deck.
// It returns NoCard when the deck is all used up.
func (d *Deck) Pop() Card {
	dd := *d
	if len(dd) == 0 {
		return NoCard
	}

	card := dd[0]
	*d = dd[1:]
	return card
}

// Shuffle shuffles the cards in the deck using a
// cryptographic random number generator.
//
// The count parameter denotes how often the deck should be shuffled.
func (d Deck) Shuffle(count int) {
	var bi *big.Int
	var idx int64
	var c, i int

	max := big.NewInt(int64(len(d)))

	if count <= 0 {
		count = 1
	}

	for c = 0; c < count; c++ {
		for i = 0; i < len(d); i++ {
			bi, _ = rand.Int(rand.Reader, max)
			idx = bi.Int64()
			d[i], d[idx] = d[idx], d[i]
		}
	}
}

// Len the amount of cards in the deck.
func (d Deck) Len() int { return len(d) }
