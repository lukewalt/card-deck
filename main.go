package main

import "fmt"

func main() {
	// assigns a new deck to a cards var
	cards := newDeck()
	cards.shuffle()

	// two vars assigned to respective return values from deal()
	hand, remainingCards := deal(cards, 5)

	fmt.Println("-- Hand --")
	hand.print()

	fmt.Println("-- Remaining Cards --")
	remainingCards.print()
}
