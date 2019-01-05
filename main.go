package main

func main() {
	// assigns a new deck to a cards var
	cards := newDeck()

	// two vars assigned to respective return values from deal()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
