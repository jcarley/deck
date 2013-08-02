// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package deck

import (
	"crypto/rand"
	"math/big"
)

const NoCard = 0xff // No card can have this value.

// A Deck represents a single deck of cards.
type Deck struct {
	list [52]Card
}

// New creates a new card deck.
func New() *Deck {
	const (
		hearts   = Card(Hearts << 4)
		diamonds = Card(Diamonds << 4)
		clubs    = Card(Clubs << 4)
		spades   = Card(Spades << 4)
	)

	d := new(Deck)
	d.list = [...]Card{
		hearts | 0,
		hearts | 1,
		hearts | 2,
		hearts | 3,
		hearts | 4,
		hearts | 5,
		hearts | 6,
		hearts | 7,
		hearts | 8,
		hearts | 9,
		hearts | 10,
		hearts | 11,
		hearts | 12,

		diamonds | 0,
		diamonds | 1,
		diamonds | 2,
		diamonds | 3,
		diamonds | 4,
		diamonds | 5,
		diamonds | 6,
		diamonds | 7,
		diamonds | 8,
		diamonds | 9,
		diamonds | 10,
		diamonds | 11,
		diamonds | 12,

		clubs | 0,
		clubs | 1,
		clubs | 2,
		clubs | 3,
		clubs | 4,
		clubs | 5,
		clubs | 6,
		clubs | 7,
		clubs | 8,
		clubs | 9,
		clubs | 10,
		clubs | 11,
		clubs | 12,

		spades | 0,
		spades | 1,
		spades | 2,
		spades | 3,
		spades | 4,
		spades | 5,
		spades | 6,
		spades | 7,
		spades | 8,
		spades | 9,
		spades | 10,
		spades | 11,
		spades | 12,
	}

	return d
}

// Pop pops the top available card from the deck and marks it as taken.
// It returns NoCard when the deck is all used up.
func (d *Deck) Pop() Card {
	for i := range d.list {
		if !d.list[i].Inuse() {
			d.list[i].use()
			return d.list[i]
		}
	}

	return NoCard
}

// Shuffle shuffles the cards in the deck using a
// cryptographic random number generator.
//
// The count parameter denotes how often you want
// to shuffle the entire deck.
func (d *Deck) Shuffle(count int) {
	var bi *big.Int
	var idx int64
	var c, i int

	max := big.NewInt(int64(len(d.list)))

	if count <= 0 {
		count = 1
	}

	for c = 0; c < count; c++ {
		for i = 0; i < len(d.list); i++ {
			bi, _ = rand.Int(rand.Reader, max)
			idx = bi.Int64()
			d.list[i], d.list[idx] = d.list[idx], d.list[i]
		}
	}
}

// Reset the deck; clears all card 'inuse' markers.
func (d *Deck) Reset() {
	for i := range d.list {
		d.list[i].unuse()
	}
}
