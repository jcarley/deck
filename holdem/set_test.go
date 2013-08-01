// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package holdem

import (
	"testing"

	"github.com/jteeuwen/deck"
)

const (
	hearts   = deck.Card(deck.Hearts << 4)
	diamonds = deck.Card(deck.Diamonds << 4)
	clubs    = deck.Card(deck.Clubs << 4)
	spades   = deck.Card(deck.Spades << 4)
)

var tests = []struct {
	Cards []deck.Card
	Hand  Hand
}{
	{
		[]deck.Card{
			hearts | 1,
			clubs | 3,
			spades | 5,
			diamonds | 7,
			hearts | 9,
			diamonds | 11,
			clubs | 12,
		},
		Highcard,
	},
	{
		[]deck.Card{
			hearts | 3,
			clubs | 2,
			spades | 6,
			diamonds | 2,
			hearts | 0,
			diamonds | 10,
			clubs | 5,
		},
		Pair,
	},
	{
		[]deck.Card{
			hearts | 3,
			clubs | 2,
			spades | 6,
			diamonds | 2,
			hearts | 0,
			diamonds | 10,
			clubs | 10,
		},
		TwoPair,
	},
	{
		[]deck.Card{
			hearts | 3,
			clubs | 10,
			spades | 6,
			diamonds | 2,
			hearts | 0,
			diamonds | 10,
			clubs | 10,
		},
		Trips,
	},
	{
		[]deck.Card{
			hearts | 3,
			clubs | 10,
			spades | 6,
			diamonds | 2,
			hearts | 1,
			diamonds | 4,
			clubs | 5,
		},
		Straight,
	},
	{
		[]deck.Card{
			clubs | 1,
			clubs | 3,
			clubs | 5,
			diamonds | 7,
			clubs | 9,
			diamonds | 11,
			clubs | 12,
		},
		Flush,
	},
	{
		[]deck.Card{
			hearts | 1,
			clubs | 1,
			spades | 11,
			diamonds | 7,
			hearts | 11,
			diamonds | 11,
			clubs | 12,
		},
		FullHouse,
	},
	{
		[]deck.Card{
			hearts | 2,
			clubs | 2,
			spades | 2,
			diamonds | 2,
			hearts | 0,
			diamonds | 10,
			clubs | 5,
		},
		Quads,
	},
	{
		[]deck.Card{
			hearts | 3,
			clubs | 10,
			spades | 6,
			hearts | 2,
			hearts | 1,
			hearts | 4,
			hearts | 5,
		},
		StraightFlush,
	},
	{
		[]deck.Card{
			diamonds | 9,
			clubs | 10,
			diamonds | 12,
			spades | 6,
			diamonds | 10,
			diamonds | 0,
			diamonds | 11,
		},
		RoyalFlush,
	},
}

func TestSet(t *testing.T) {
	for _, tc := range tests {
		set := Set(tc.Cards)
		hand, _, _ := set.Hand()

		if hand != tc.Hand {
			t.Fatalf("Hand mismatch: Want %v, have %v", tc.Hand, hand)
		}

		println(set.String(), hand.String())
	}
}
