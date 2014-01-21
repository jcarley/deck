// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package deck

import "strconv"

// A single card holds its value, suit and state encoded in 8 bits.
//
//     bits 0-3: The card value (0-12).
//     bits 4-5: The suit (0-3).
//     bit    6: <reserved for future use>
//     bit    7: <reserved for future use>
type Card uint8

// NewCard creates a new card from the given suit and value.
func NewCard(s Suit, num uint8) Card {
	return Card(uint8(s<<4) | num&15)
}

// Value returns the value for this card.
func (c Card) Value() uint8 { return uint8(c & 15) }

// Suit returns the suit for this card.
func (c Card) Suit() Suit { return Suit(c >> 4 & 3) }

func (c Card) String() string {
	num := uint8(c&15) + 1
	name := Suit(c >> 4 & 3).String()

	if num > 1 && num < 11 {
		return name + strconv.Itoa(int(num))
	}

	switch num {
	case 1:
		name += "A"
	case 11:
		name += "J"
	case 12:
		name += "Q"
	case 13:
		name += "K"
	}

	return name
}
