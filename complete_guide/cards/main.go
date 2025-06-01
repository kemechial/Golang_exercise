package main

import (
	"fmt"
)

/*




func main() {
	var card string = "Ace of Spades"
	card2 := "King of Hearts"
	card2 = newCard()
	fmt.Println("Welcome to the card game!")
	fmt.Printf("My favorite card is: %s and also %s\n", card, card2)

	cards := []string{"Two of Clubs", "Three of Diamonds", "Four of Spades"}
	fmt.Println(cards)

	cards = append(cards, "Six of Hearts")
	for i, card := range cards {
		fmt.Println(i, card)
	}

}


func main() {

	//cards := deck{newCard(), "Two of Clubs", "Three of Diamonds", "Four of Spades"}
	//cards = append(cards, "Six of Hearts")
	cards := newDeck()
	cards.print()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	fmt.Println("Converted to string:")
	print(hand.toString()+"\n")
	cards.saveToFile("my_cards.txt")
	cardsFromFile, err := newDeckFromFile("my_cards.txt")
	if condition := err != nil; condition {
		fmt.Println("Error reading from file:", err)
		os.Exit(1)
	} else {
		fmt.Println("Cards read from file:")
		cardsFromFile.print()
		
	}

}   

func newCard() string {
	return "Five of Diamonds"
}
*/

func main()  {	
	cards := newDeck()
	cards.print()
	fmt.Println("######### Shuffled deck ##########")
	cards.shuffle()
	cards.print()
	
}