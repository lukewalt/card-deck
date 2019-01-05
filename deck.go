package main

import "fmt"

// tells Go we want to create a new type
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

// reciever is like a method to an instance spec to type
// (input type) any variable of type deck gets access to this method
// uses one/two letters from type instead of 'self' or 'this'
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
