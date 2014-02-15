// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package deck

import (
	"math/rand"
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

// Shuffle shuffles the cards in the deck using the
// Knuth shuffle algorithm.
func (d Deck) Shuffle(rng *rand.Rand) {
	for a := 0; a < len(d); a++ {
		b := rng.Int31n(int32(a) + 1)
		d[a], d[b] = d[b], d[a]
	}
}

// Len the amount of cards in the deck.
func (d Deck) Len() int { return len(d) }
