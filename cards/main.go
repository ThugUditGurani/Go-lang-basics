package main

func main() {
	//Variable data type
	//var card string = "Ace of Spades"
	//card = "Overrite Ace of Shades New"

	//function Call
	//card := newCard()

	//Slices and Array
	// cards := deck{newCard(), newCard()}
	// cards = append(cards, "Six of Shades")
	// fmt.Println(cards)
	// cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//	fmt.Println(card)

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// fmt.Println(cards)
	// hand.print()
	// remainingCards.print()

	//Save Retrive and Create Files : string to byte typeCasting
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// Shuffle
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
