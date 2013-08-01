// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package holdem

import (
	"fmt"
	"strings"

	"github.com/jteeuwen/deck"
)

// A set represents a subset of cards.
// Either in a players hands or on the table, or both.
//
// This can be used to determine if a player, at any given time
// during a game, has a winning hand or not and how he/she ranks
// compared to other players.
type Set []deck.Card

func (s Set) String() string {
	list := make([]string, len(s))

	for i, c := range s {
		list[i] = fmt.Sprintf("%-3s", c)
	}

	return strings.Join(list, " ")
}

// Hand determines the highest kind of hand this set represents.
// It returns the hand identifier, the total value of all cards
// in the set and the highcard value.
func (s Set) Hand() (Hand, int, int) {
	var high, total int

	for i := range s {
		if total = int(s[i].Value() + 1); total == 1 {
			total = 14 // Ace counts as 14
		}

		if total > high {
			high = total
		}
	}

	if total = s.isRoyalFlush(); total > -1 {
		return RoyalFlush, total, high
	}

	if total = s.isStraightFlush(); total > -1 {
		return StraightFlush, total, high
	}

	if total = s.isQuads(); total > -1 {
		return Quads, total, high
	}

	if total = s.isFullhouse(); total > -1 {
		return FullHouse, total, high
	}

	if total = s.isFlush(); total > -1 {
		return Flush, total, high
	}

	if total = s.isStraight(); total > -1 {
		return Straight, total, high
	}

	if total = s.isTrips(); total > -1 {
		return Trips, total, high
	}

	if total = s.isTwoPair(); total > -1 {
		return TwoPair, total, high
	}

	if total = s.isPair(); total > -1 {
		return Pair, total, high
	}

	return Highcard, high, high
}

func (s Set) isRoyalFlush() int {
	var st deck.Suit

	for st = 0; st < 4; st++ {
		if s.has(st, 0) && s.has(st, 9) && s.has(st, 10) &&
			s.has(st, 11) && s.has(st, 12) {
			return 60 // Ace counts as 14, not 1
		}
	}

	return -1
}

func (s Set) isStraightFlush() int {
	var st deck.Suit
	var v uint8

	for st = 0; st < 4; st++ {
		for v = 0; v < 9; v++ {
			if s.has(st, v) &&
				s.has(st, v+1) &&
				s.has(st, v+2) &&
				s.has(st, v+3) &&
				s.has(st, v+4) {
				return int(v+1)*5 + 15 // Ace counts as 1
			}
		}
	}

	return -1
}

func (s Set) isQuads() int {
	for v := uint8(0); v < 13; v++ {
		if s.hasValueCount(v, 4) {
			if v == 0 {
				return 56 // Aces count as 14, not 1
			}

			return int(v+1) * 4
		}
	}

	return -1
}

func (s Set) isFullhouse() int {
	var value int
	var a, b uint8

	for ; a < 13; a++ {
		if s.hasValueCount(a, 3) {
			if a == 0 {
				value = 42 // Ace counts as 14
			} else {
				value = int(a+1) * 3
			}

			break
		}
	}

	if value == 0 {
		return -1
	}

	for ; b < 13; b++ {
		if b != a && s.hasValueCount(b, 2) {
			if b == 0 {
				return value + 28 // Ace counts as 14
			}

			return value + int(b+1)*2
		}
	}

	return -1
}

func (s Set) isFlush() (value int) {
	for st := deck.Suit(0); st < 4; st++ {
		if value = s.hasSuitCount(st, 5); value > -1 {
			return
		}
	}

	return -1
}

func (s Set) isStraight() int {
	if s.hasValue(0) && s.hasValue(9) && s.hasValue(10) && s.hasValue(11) && s.hasValue(12) {
		return 60 // Ace counts as 14
	}

	for v := uint8(0); v < 9; v++ {
		if s.hasValue(v) &&
			s.hasValue(v+1) &&
			s.hasValue(v+2) &&
			s.hasValue(v+3) &&
			s.hasValue(v+4) {
			return int(v+1)*5 + 15 // Ace counts as 1
		}
	}

	return -1
}

func (s Set) isTrips() int {
	for v := uint8(0); v < 13; v++ {
		if s.hasValueCount(v, 3) {
			if v == 0 {
				return 42 // Ace counts as 14
			}

			return int(v+1) * 3
		}
	}

	return -1
}

func (s Set) isTwoPair() int {
	var value int
	var a, b uint8

	for ; a < 13; a++ {
		if s.hasValueCount(a, 2) {
			if a == 0 {
				value = 28 // Ace counts as 14
			} else {
				value = int(a+1) * 2
			}
			break
		}
	}

	if value == 0 {
		return -1
	}

	for ; b < 13; b++ {
		if b != a && s.hasValueCount(b, 2) {
			if b == 0 {
				return value + 28 // Ace counts as 14
			}

			return value + int(b+1)*2
		}
	}

	return -1
}

func (s Set) isPair() int {
	for v := uint8(0); v < 13; v++ {
		if s.hasValueCount(v, 2) {
			if v == 0 {
				return 28 // Ace counts as 14
			}

			return int(v+1) * 2
		}
	}

	return -1
}

// has tests if the set contains the given card.
func (s Set) has(st deck.Suit, v uint8) bool {
	for i := range s {
		if s[i].Suit() == st && s[i].Value() == v {
			return true
		}
	}

	return false
}

// hasValue tests if the set contains the given card value.
func (s Set) hasValue(v uint8) bool {
	for i := range s {
		if s[i].Value() == v {
			return true
		}
	}

	return false
}

// hasValueCount tests if the set contains N amount of cards with value V
func (s Set) hasValueCount(v uint8, n int) bool {
	var c int

	for i := range s {
		if s[i].Value() == v {
			if c++; c >= n {
				return true
			}
		}
	}

	return false
}

// hasSuitCount tests if the set contains N amount of cards with suit st.
func (s Set) hasSuitCount(st deck.Suit, n int) int {
	var c, val int

	for i := range s {
		if s[i].Suit() == st {
			val += int(s[i].Value() + 1)
			if c++; c >= n {
				return val
			}
		}
	}

	return -1
}
