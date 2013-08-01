// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package deck

type Suit uint8

// Available suits
const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

func (s Suit) String() string {
	switch s {
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		return "♠"
	}
}
