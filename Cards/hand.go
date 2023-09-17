package main

func deal(d deck, handSize int) (deck, deck) {
	drawnHand := d[:handSize]

	remainingDeck := d[handSize:]

	return drawnHand, remainingDeck
}