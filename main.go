package main

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades") // immutable function

	// pass var cards to print
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
