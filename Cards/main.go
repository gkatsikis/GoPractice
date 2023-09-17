package main

func main() {
	cards := newDeck()

	cards.print()

	drawnHand, cards := deal(cards, 5)

	drawnHand.print()
}
