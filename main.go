package main

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades") // immutable function

	// pass var cards to print
	cards.print()
}
