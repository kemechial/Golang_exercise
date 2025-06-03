package main
import (
	"strings"
	"os"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}


func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	remainingDeck := d[handSize:]
	return hand, remainingDeck
}
/*
func (d deck) toString() string {
	s := ""
	for _, card := range d {
		s += card + ", "
	}
	return s[:len(s)-2] // Remove the trailing comma and space
}
*/

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {

	bs, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s := string(bs)
	s = strings.TrimSpace(s) // Remove any trailing newline characters
	cards := strings.Split(s, ", ")
	return deck(cards), nil
}
// This function shuffles the deck in place, it is a receiver function which takes a copy of the object, 
// however deck is a slice, so it is passed by reference, thus it can modify the original deck.
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}