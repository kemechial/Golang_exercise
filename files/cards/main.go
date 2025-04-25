package main

import "fmt"
//decksize := 52 This would give an error because it is not declared in a function
var decksize int = 52 // global variable, declared outside of any function

func main()  {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" // shorthand declaration, := infer the type
	card = "Ace of Diamonds" // reassigning the value of card
	fmt.Println(card)
    fmt.Println(decksize)
	fmt.Println(newCard()) // calling the newCard function
}

func newCard() string {
	card := "Ace of Spades" // shorthand declaration, := infer the type
	return card
}