package business

import (
	"fmt"
	"poker_decoder/internal/constants"
	"poker_decoder/internal/models"
)

// Init initializes the poker decoder
func Init() {
	fmt.Println("Starting poker decoder...")
	createDeck()

}

// CreateDeck creates a deck of 52 cards
func createDeck() {
	fmt.Println("Creating deck of 52 cards...")
	deck := []models.Card{}

	// Add cards to deck
	for _, suit := range constants.Suits {
		for _, value := range constants.Values {
			card := models.Card{Suit: suit, Value: value, Shortname: value + suit}
			deck = append(deck, card)
		}
	}
	fmt.Println("Deck created successfully!")
	fmt.Println("")

	ShuffleDeck(&deck)

	for i, card := range deck {
		fmt.Println(i, card.Shortname)
	}
}
