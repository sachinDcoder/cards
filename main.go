package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// err := cards.writeToFile("my-cards.txt")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	cards := newDeckFromFile("my-cards.txt")
	cards.print()
}
