## Hold'em

This package is an extension to the `deck` package.

It offers means to determine the value and type of a Texas Hold'em poker
hand which is currently contained in a given set of cards.


### Usage

    go get github.com/jteeuwen/deck/holdem


### Example sets

Here are examples of all the possible hands this package can detect.
Starting with the input cards, followed by the total score of the hand,
the highcard value and finally the name of the detected hand.


	CARDS                      | SCORE | HIGHCARD VALUE | HAND
	===========================|=======|================|==================
	♥2  ♣4  ♠6  ♦8  ♥10 ♦Q  ♣K |    13 |             13 | High card
	♥4  ♣3  ♠7  ♦3  ♥A  ♦J  ♣6 |     6 |             14 | Pair
	♥4  ♣3  ♠7  ♦3  ♥A  ♦J  ♣J |    28 |             14 | Two pair
	♥4  ♣J  ♠7  ♦3  ♥A  ♦J  ♣J |    33 |             14 | Three of a kind
	♥4  ♣J  ♠7  ♦3  ♥2  ♦5  ♣6 |    25 |             11 | Straight
	♣2  ♣4  ♣6  ♦8  ♣10 ♦Q  ♣K |    35 |             13 | Flush
	♥2  ♣2  ♠Q  ♦8  ♥Q  ♦Q  ♣K |    40 |             13 | Full house
	♥3  ♣3  ♠3  ♦3  ♥A  ♦J  ♣6 |    12 |             14 | Four of a kind
	♥4  ♣J  ♠7  ♥3  ♥2  ♥5  ♥6 |    25 |             11 | Straight flush
	♦10 ♣J  ♦K  ♠7  ♦J  ♦A  ♦Q |    60 |             14 | Royal flush



### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

