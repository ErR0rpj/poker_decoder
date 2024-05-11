package business

import (
	"fmt"
	"math/rand"
	"poker_decoder/internal/models"
)

func ShuffleDeck(deck *[]models.Card) {
	fmt.Println("Shuffling deck...")
	rand.Shuffle(len(*deck), func(i, j int) {
		(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
	})
	fmt.Println("Deck shuffled!")
	fmt.Println("")
}
