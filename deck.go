package main

import "fmt"

// tells Go we want to create array of strings as well as functions specific to it
// func w/ reciever is like a method to an instance
type deck []string

// (input type) any variable of type deck gets access to this method
// uses one/two letters from type instead of 'self' or 'this'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
