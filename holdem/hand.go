// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package holdem

type Hand uint8

// Hand kinds
const (
	Highcard Hand = iota
	Pair
	TwoPair
	Trips
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
	RoyalFlush
)

func (h Hand) String() string {
	switch h {
	case Highcard:
		return "High card"
	case Pair:
		return "Pair"
	case TwoPair:
		return "Two pair"
	case Trips:
		return "Three of a kind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "Full house"
	case Quads:
		return "Four of a kind"
	case StraightFlush:
		return "Straight flush"
	case RoyalFlush:
		return "Royal flush"
	}

	panic("Unknown hand")
}
